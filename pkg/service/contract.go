package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dontpanicdao/caigo/gateway"
	"github.com/dontpanicdao/caigo/types"
	"github.com/goperate/convert/core/array"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/utils"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ContractService struct {
	db *gorm.DB
	gw *gateway.Gateway
}

func NewContractService() *ContractService {
	gw := gateway.NewClient(gateway.WithChain(gateway.GOERLI_ID))
	return &ContractService{
		db: application.GetBean[*gorm.DB]("db"),
		gw: gw,
	}
}

func (c *ContractService) DoStarknetDeclare(compiledContract []byte) (txHash string, classHash string, err error) {
	gw := gateway.NewClient(gateway.WithChain(gateway.GOERLI_ID))

	ctx := context.Background()
	var contractClass types.ContractClass

	err = json.Unmarshal(compiledContract, &contractClass)

	if err != nil {
		return "", "", err
	}

	declare, err := gw.Declare(ctx, contractClass, gateway.DeclareRequest{})
	if err != nil {
		return "", "", err
	}

	fmt.Println("declare.TransactionHash: ", declare.TransactionHash)
	fmt.Println("declare.ClassHash: ", declare.ClassHash)

	_, receipt, err := gw.WaitForTransaction(ctx, declare.TransactionHash, 3, 10)
	if err != nil {
		fmt.Printf("could not declare contract: %v\n", err)
		return "", "", err
	}
	if receipt.Status != types.TransactionAcceptedOnL1 && receipt.Status != types.TransactionAcceptedOnL2 {
		fmt.Printf("unexpected status: %s\n", receipt.Status)
		return "", "", err
	}

	return declare.TransactionHash, declare.ClassHash, nil
}

func (c *ContractService) SaveDeploy(entity db2.ContractDeploy) (uint, error) {
	var contract db2.Contract
	err := c.db.Model(db2.Contract{}).Where("id = ?", entity.ContractId).First(&contract).Error
	if err != nil {
		return 0, err
	}
	version, err := strconv.Atoi(entity.Version)
	if err != nil {
		return 0, err
	}
	entity.Type = contract.Type
	if entity.AbiInfo == "" && version > 1 {
		for {
			if version > 1 {
				var contractDeploy db2.ContractDeploy
				c.db.Model(db2.ContractDeploy{}).Where("contract_id = ? and version = ? ", entity.ContractId, entity.Version).First(&contractDeploy)
				if contractDeploy.AbiInfo != "" {
					entity.AbiInfo = contractDeploy.AbiInfo
					break
				}
				version = version - 1
			} else {
				break
			}
		}
	}
	if contract.AbiInfo == "" {
		contract.AbiInfo = entity.AbiInfo
	}
	err = c.db.Create(&entity).Error
	if err != nil {
		return 0, err
	}

	contract.Network = sql.NullString{
		String: entity.Network,
		Valid:  true,
	}
	contract.Status = entity.Status
	c.db.Save(&contract)
	return entity.Id, err
}

func (c *ContractService) QueryContracts(projectId string, query, version, network string, page int, size int) (vo.Page[db2.Contract], error) {
	var contracts []db2.Contract
	var afterData []db2.Contract
	sql := fmt.Sprintf("select id, project_id,workflow_id,workflow_detail_id,name,version,group_concat( DISTINCT `network` SEPARATOR ',' ) as network,build_time,abi_info,byte_code,create_time from t_contract where project_id = ? ")
	if query != "" && version != "" && network != "" {
		sql = sql + "and name like CONCAT('%',?,'%') and version = ? and network = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, query, version, network).Scan(&contracts)
	} else if query != "" && version != "" {
		sql = sql + "and name like CONCAT('%',?,'%') and version = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, query, version).Scan(&contracts)
	} else if query != "" && network != "" {
		sql = sql + "and name like CONCAT('%',?,'%') and network = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, query, network).Scan(&contracts)
	} else if version != "" && network != "" {
		sql = sql + "and version = ? and network = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, version, network).Scan(&contracts)
	} else if query != "" {
		sql = sql + "and name like CONCAT('%',?,'%') group by id order by create_time desc"
		c.db.Raw(sql, projectId, query).Scan(&contracts)
	} else if network != "" {
		sql = sql + "and network = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, network).Scan(&contracts)
	} else if version != "" {
		sql = sql + "and version = ? group by id order by create_time desc"
		c.db.Raw(sql, projectId, version).Scan(&contracts)
	} else {
		sql = sql + "group by id order by create_time desc"
		c.db.Raw(sql, projectId).Scan(&contracts)
	}
	if len(contracts) > 0 {
		start, end := utils.SlicePage(int64(page), int64(size), int64(len(contracts)))
		afterData = contracts[start:end]
	}
	return vo.NewPage[db2.Contract](afterData, len(contracts), page, size), nil
}

func (c *ContractService) QueryContractByWorkflow(workflowId, workflowDetailId int) ([]db2.Contract, error) {
	var contracts []db2.Contract
	res := c.db.Model(db2.Contract{}).Where("workflow_id = ? and workflow_detail_id = ?", workflowId, workflowDetailId).Order("version DESC").Find(&contracts)
	if res != nil {
		return contracts, res.Error
	}
	return contracts, nil
}

func (c *ContractService) QueryContractByVersion(projectId string, version string) ([]vo.ContractVo, error) {
	var contracts []db2.Contract
	var data []vo.ContractVo
	res := c.db.Model(db2.Contract{}).Where("project_id = ? and version = ?", projectId, version).Find(&contracts)
	if res.Error != nil {
		return data, res.Error
	}
	if len(contracts) > 0 {
		_ = copier.Copy(&data, &contracts)
	}
	return data, nil
}

func (c *ContractService) QueryContractDeployByVersion(projectId string, version string) (vo.ContractDeployInfoVo, error) {
	var data vo.ContractDeployInfoVo
	var contractDeployData []db2.ContractDeploy
	res := c.db.Model(db2.ContractDeploy{}).Where("project_id = ? and version = ?", projectId, version).Find(&contractDeployData)
	if res.Error != nil {
		return data, res.Error
	}
	contractInfo := make(map[string]vo.ContractInfoVo)
	if len(contractDeployData) > 0 {
		arr := array.NewObjArray(contractDeployData, "ContractId")
		res2 := arr.ToIdMapArray().(map[uint][]db2.ContractDeploy)
		for u, deploys := range res2 {
			var contractData db2.Contract
			res := c.db.Model(db2.Contract{}).Where("id = ?", u).First(&contractData)
			var contractInfoVo vo.ContractInfoVo
			copier.Copy(&contractInfoVo, &contractData)
			if res.Error == nil {
				var deployInfo []vo.DeployInfVo
				if len(deploys) > 0 {
					for _, deploy := range deploys {
						var deployData vo.DeployInfVo
						copier.Copy(&deployData, &deploy)
						deployInfo = append(deployInfo, deployData)
						if deploy.AbiInfo != "" && contractInfoVo.AbiInfo == "" {
							contractInfoVo.AbiInfo = deploy.AbiInfo
						}
					}
				}
				contractInfoVo.DeployInfo = deployInfo
				contractInfo[contractData.Name] = contractInfoVo
			}
		}
	}
	data.Version = version
	data.ContractInfo = contractInfo
	return data, nil
}

func (c *ContractService) QueryVersionList(projectId string) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Contract{}).Distinct("version").Select("version").Where("project_id = ?", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}

func (c *ContractService) QueryContractNameList(projectId string) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Contract{}).Distinct("name").Select("name").Where("project_id = ?", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}

func (c *ContractService) QueryNetworkList(projectId string) ([]string, error) {
	var data []string
	res := c.db.Model(db2.Contract{}).Distinct("network").Select("network").Where("project_id = ? and network != '' ", projectId).Find(&data)
	if res.Error != nil {
		return data, res.Error
	}
	return data, nil
}

func (c *ContractService) SyncStarkWareContract() {

	var contractDeploys []db2.ContractDeploy

	const running = 1
	err := c.db.Model(db2.ContractDeploy{}).Where("type=? and status = ?", consts.StarkWare, running).Find(&contractDeploys).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, deploy := range contractDeploys {

		if time.Since(deploy.DeployTime) > time.Minute*15 {
			// fail
			deploy.Status = 3
		} else if deploy.DeployTxHash != "" {
			receipt, err := c.gw.TransactionReceipt(context.Background(), deploy.DeployTxHash)
			if err != nil {
				continue
			}
			if receipt.Status == types.TransactionAcceptedOnL2 {
				// success

				// query contract address
				if len(receipt.Events) > 0 {
					event1 := receipt.Events[0].(map[string]interface{})
					data := event1["data"].([]interface{})
					if len(data) > 0 {
						contractAddress := data[0].(string)
						deploy.Address = contractAddress
						deploy.Status = 1

					}
				}
			}
		}
		err := c.db.Save(&deploy).Error
		if err != nil {
			fmt.Println("save contractDeploy error")
			continue
		}
		var contract db2.Contract
		err = c.db.Model(db2.Contract{}).Where("id = ?", deploy.ContractId).First(&contract).Error
		if err != nil {
			fmt.Println("save Contract error")
			continue
		}
		contract.Status = deploy.Status
		c.db.Save(contract.Status)
	}
}

func (c *ContractService) GetContractDeployInfo(id int) (db2.ContractDeploy, error) {
	var result db2.ContractDeploy
	err := c.db.Model(db2.ContractDeploy{}).First(&result, id).Error
	return result, err
}
