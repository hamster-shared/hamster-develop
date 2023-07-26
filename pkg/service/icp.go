package service

import (
	"database/sql"
	"errors"
	"github.com/hamster-shared/aline-engine/logger"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"gorm.io/gorm"
	"os/exec"
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
	newIdentityCmd := "dfx identity new " + identityName
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
