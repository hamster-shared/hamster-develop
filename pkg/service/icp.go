package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"os"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type IcpService struct {
	db *gorm.DB
}

func NewIcpService() *IcpService {
	return &IcpService{
		db: application.GetBean[*gorm.DB]("db"),
	}
}

func (i *IcpService) CreateIdentity(userId uint) (vo vo.UserIcpInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err == nil {
		return vo, errors.New("you have created an identity")
	}
	if err != gorm.ErrRecordNotFound {
		return vo, err
	}
	identityName := strconv.Itoa(int(userId))
	newIdentityCmd := "dfx identity new " + identityName + " --storage-mode plaintext"
	_, error = i.execDfxCommand(newIdentityCmd)
	if err != nil {
		return vo, err
	}
	aId, pId, err := i.getLedgerInfo(identityName)
	if err == nil {
		return vo, err
	}
	userIcp.FkUserId = userId
	userIcp.IdentityName = identityName
	userIcp.AccountId = strings.TrimSpace(aId)
	userIcp.PrincipalId = strings.TrimSpace(pId)
	userIcp.CreateTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	userIcp.UpdateTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = i.db.Model(db.UserIcp{}).Create(&userIcp).Error
	if err != nil {
		return vo, err
	}
	vo.UserId = int(userId)
	vo.AccountId = aId
	vo.IcpBalance = "0.0000000 ICP"
	return vo, nil
}

func (i *IcpService) GetAccountInfo(userId uint) (vo vo.UserIcpInfoVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	ledgerBalanceCmd := "dfx ledger balance --network ic"
	balance, err := i.execDfxCommand(ledgerBalanceCmd)
	if err != nil {
		return vo, err
	}
	vo.UserId = int(userIcp.FkUserId)
	vo.AccountId = userIcp.AccountId
	vo.IcpBalance = strings.TrimSpace(balance)
	return vo, nil
}

func (i *IcpService) RedeemFaucetCoupon(userId uint, redeemFaucetCouponParam parameter.RedeemFaucetCouponParam) (vo vo.IcpCanisterBalanceVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	redeemCouponCmd := "dfx wallet --network ic redeem-faucet-coupon " + redeemFaucetCouponParam.Coupon
	output, err := i.execDfxCommand(redeemCouponCmd)
	if err != nil {
		return vo, err
	}
	walletId := ""
	split := strings.Split(output, "\n")
	for _, str := range split {
		if strings.Contains(str, "new wallet:") {
			lastIndex := strings.LastIndex(str, ":")
			if lastIndex != -1 {
				walletId = strings.TrimSpace(str[lastIndex+1:])
			}
		}
	}
	if walletId == "" {
		return vo, errors.New("failed to generate wallet")
	}
	userIcp.WalletId = walletId
	error = i.db.Model(db.UserIcp{}).Updates(&userIcp).Error
	if error != nil {
		return vo, errors.New("failed to save wallet ID")
	}
	return i.GetWalletInfo(userId)
}

func (i *IcpService) GetWalletInfo(userId uint) (vo vo.IcpCanisterBalanceVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	walletBalanceCmd := "dfx wallet balance --network ic"
	balance, err := i.execDfxCommand(walletBalanceCmd)
	if err != nil {
		return vo, err
	}
	vo.UserId = int(userIcp.FkUserId)
	vo.CanisterId = userIcp.WalletId
	vo.CyclesBalance = balance
	return vo, nil
}

func (i *IcpService) RechargeWallet(userId uint) (vo vo.IcpCanisterBalanceVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	if userIcp.WalletId == "" {
		walletId, err := i.InitWallet(userIcp)
		if err != nil {
			return vo, err
		}
		userIcp.WalletId = walletId
		err = i.db.Model(db.UserIcp{}).Updates(&userIcp).Error
		if err != nil {
			return vo, err
		}
	} else {
		err := i.WalletTopUp(userIcp.IdentityName, userIcp.WalletId)
		if err != nil {
			return vo, err
		}
	}
	return i.GetWalletInfo(userId)
}

func (i *IcpService) RechargeCanister(userId uint, rechargeCanisterParam parameter.RechargeCanisterParam) (vo vo.IcpCanisterBalanceVo, error error) {
	var userIcp db.UserIcp
	err := i.db.Model(db.UserIcp{}).Where("fk_user_id = ?", userId).First(&userIcp).Error
	if err != nil {
		return vo, err
	}
	// 判断当前目录是否存在 dfx.json 文件
	if _, err := os.Stat("dfx.json"); os.IsNotExist(err) {
		// 不存在，则新建并写入数据 {}
		data := map[string]interface{}{}
		dataJSON, err := json.Marshal(data)
		if err != nil {
			return vo, err
		}
		err = os.WriteFile("dfx.json", dataJSON, 0644)
		if err != nil {
			return vo, err
		}
	}
	depositCycles := rechargeCanisterParam.Amount * 1e12
	i.canisterRechargeCycles(userIcp.IdentityName, strconv.FormatFloat(float64(depositCycles), 'f', -1, 64), rechargeCanisterParam.CanisterId)
	err = os.Remove("dfx.json")
	if err != nil {
		return vo, err
	}
	return vo, nil
}

func (i *IcpService) canisterRechargeCycles(identityName string, cycles string, canisterId string) (error error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	useIdentityCmd := "dfx identity use " + identityName
	_, err := i.execDfxCommand(useIdentityCmd)
	if err != nil {
		return err
	}

	depositCyclesCmd := "dfx canister deposit-cycles" + cycles + " " + canisterId + " --network ic "
	output, err := i.execDfxCommand(depositCyclesCmd)
	logger.Infof("userid-> %s canisterId-> %s deposit-cycles result is: %s \n", identityName, canisterId, output)
	if err != nil {
		return err
	}
	return nil
}

func (i *IcpService) InitWallet(userIcp db.UserIcp) (walletId string, error error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	useIdentityCmd := "dfx identity use " + userIcp.IdentityName
	_, err := i.execDfxCommand(useIdentityCmd)
	if err != nil {
		return "", err
	}
	balance, err := i.getLedgerIcpBalance()
	if err != nil {
		return "", err
	}
	createCanisterCmd := "dfx ledger --network ic create-canister " + userIcp.PrincipalId + " --amount " + balance
	output, err := i.execDfxCommand(createCanisterCmd)
	logger.Infof("userid-> %s create-canister result is: %s \n", userIcp.IdentityName, output)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`Canister created with id: "(.*?)"`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		walletId = matches[1]
	} else {
		return "", errors.New("failure to create-canister")
	}
	deployWalletCmd := "dfx identity --network ic deploy-wallet " + walletId
	output, err = i.execDfxCommand(deployWalletCmd)
	logger.Infof("userid-> %s walletId-> %s deploy-wallet result is: %s \n", userIcp.IdentityName, walletId, output)
	if err != nil {
		return "", err
	}
	return walletId, nil
}

func (i *IcpService) WalletTopUp(identityName string, walletId string) (error error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	useIdentityCmd := "dfx identity use " + identityName
	_, err := i.execDfxCommand(useIdentityCmd)
	if err != nil {
		return err
	}
	balance, err := i.getLedgerIcpBalance()
	if err != nil {
		return err
	}

	walletTopUpCmd := "  dfx ledger --network ic top-up " + walletId + " --amount " + balance
	output, err := i.execDfxCommand(walletTopUpCmd)
	if err != nil {
		return err
	}
	logger.Infof("identityName-> %s walletId-> %s top-up result is: %s \n", identityName, walletId, output)
	return nil
}

func (i *IcpService) getLedgerIcpBalance() (string, error) {
	ledgerBalanceCmd := "dfx ledger balance --network ic"
	balance, err := i.execDfxCommand(ledgerBalanceCmd)
	if err != nil {
		return "", err
	}
	balanceSplit := strings.Split(balance, " ")
	if len(balanceSplit) > 0 {
		return balanceSplit[0], nil
	} else {
		return "", errors.New("failure to obtain ICP balances")
	}
}

func (i *IcpService) getLedgerInfo(identityName string) (string, string, error) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()
	useIdentityCmd := "dfx identity use " + identityName
	_, err := i.execDfxCommand(useIdentityCmd)
	if err != nil {
		return "", "", err
	}
	accountIdCmd := "dfx ledger account-id"
	accountId, err := i.execDfxCommand(accountIdCmd)
	if err != nil {
		return "", "", err
	}
	pIdCmd := "dfx identity get-principal"
	pId, err := i.execDfxCommand(pIdCmd)
	if err != nil {
		return "", "", err
	}

	return accountId, pId, nil
}

func (i *IcpService) execDfxCommand(cmd string) (string, error) {
	output, err := exec.Command("bash", "-c", cmd).Output()
	if exitError, ok := err.(*exec.ExitError); ok {
		logger.Errorf("%s Exit status: %d, Exit str: %s", cmd, exitError.ExitCode(), string(exitError.Stderr))
		return "", err
	} else if err != nil {
		// 输出其他类型的错误
		logger.Errorf("%s Failed to execute command: %s", cmd, err)
		return "", err
	}
	return string(output), nil
}

func (i *IcpService) QueryIcpCanisterList(projectId string, page, size int) (*vo.IcpCanisterPage, error) {
	var total int64
	var pageData vo.IcpCanisterPage
	var canisters []db.IcpCanister
	var vo []vo.IcpCanisterVo
	err := i.db.Model(db.IcpCanister{}).Where("project_id = ?", projectId).Order("create_time DESC").Offset((page - 1) * size).Limit(size).Find(&canisters).Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		return &pageData, err
	}
	for _, canister := range canisters {
		if !canister.Cycles.Valid {
			data, err := i.queryCanisterStatus(canister.CanisterId)
			if err == nil {
				canister.Cycles = sql.NullString{
					String: data.Balance,
					Valid:  true,
				}
				canister.UpdateTime = sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				}
				i.db.Save(&canister)
			}
		} else {
			isThreeHoursAgo := isTimeThreeHoursAgo(canister.UpdateTime.Time, time.Now())
			if isThreeHoursAgo {
				data, err := i.queryCanisterStatus(canister.CanisterId)
				if err == nil {
					canister.Cycles = sql.NullString{
						String: data.Balance,
						Valid:  true,
					}
					canister.UpdateTime = sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					}
					i.db.Save(&canister)
				}
			}
		}
	}
	copier.Copy(&vo, &canisters)
	pageData.Data = vo
	pageData.Page = page
	pageData.PageSize = size
	pageData.Total = int(total)
	return &pageData, nil
}

func (i *IcpService) queryCanisterStatus(canisterId string) (vo.CanisterStatusRes, error) {
	var res vo.CanisterStatusRes
	canisterCmd := fmt.Sprintf("dfx canister status %s", canisterId)
	cmd := exec.Command("bash", "-c", canisterCmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("cmd exec failed: %s", err)
		return res, err
	}
	re := regexp.MustCompile(`Balance: ([0-9_]+) Cycles`)
	matches := re.FindStringSubmatch(string(out))
	if len(matches) > 1 {
		value := matches[1]
		value = strings.ReplaceAll(value, "_", "")
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logger.Errorf("balance parse int failed:%s", err)
			return res, err
		}
		data := float64(number) / math.Pow(10, 12)
		balance := fmt.Sprintf("%.2f\n", data)
		res.Balance = balance
	} else {
		logger.Info("balance not found!")
	}
	statusRegex := regexp.MustCompile(`Status: (.+)`)
	statusMatch := statusRegex.FindStringSubmatch(string(out))
	if len(statusMatch) > 1 {
		res.Status = statusMatch[1]
	} else {
		logger.Info("status not found!")
	}
	return res, nil
}

func isTimeThreeHoursAgo(t time.Time, now time.Time) bool {
	duration := now.Sub(t)
	return duration >= 3*time.Hour
}

func (i *IcpService) SaveDfxJsonData(projectId string, jsonData string) error {
	var dfxData db.IcpDfxData
	err := i.db.Where("project_id = ?", projectId).First(&dfxData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			dfxData.ProjectId = projectId
			dfxData.DfxData = jsonData
			dfxData.CreateTime = sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			err = i.db.Create(&dfxData).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}
func (i *IcpService) QueryDfxJsonDataByProjectId(projectId string) (vo.IcpDfxDataVo, error) {
	var data db.IcpDfxData
	var vo vo.IcpDfxDataVo
	err := i.db.Model(db.IcpDfxData{}).Where("project_id = ?", projectId).First(&data).Error
	if err != nil {
		return vo, err
	}
	copier.Copy(&vo, &data)
	return vo, nil
}
func (i *IcpService) UpdateDfxJsonData(id int, jsonData string) error {
	var data db.IcpDfxData
	err := i.db.Model(db.IcpDfxData{}).Where("id = ?", id).First(&data).Error
	if err != nil {
		return err
	}
	data.DfxData = jsonData
	err = i.db.Save(&data).Error
	return err
}
