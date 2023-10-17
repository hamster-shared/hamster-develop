package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dontpanicdao/caigo/gateway"
	"github.com/dontpanicdao/caigo/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hamster-shared/hamster-develop/pkg/vo"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"testing"
	"time"
)

func NewTestContractService() *ContractService {
	DSN := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/aline?charset=utf8&parseTime=True&loc=Local", "123456")
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DSN,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	contractService := &ContractService{
		db: db,
	}
	return contractService
}

func TestQueryContract(t *testing.T) {

	contractService := NewTestContractService()
	_, err := contractService.QueryContractByWorkflow("id", 1, 1)
	assert.NoError(t, err)
}

func TestSync(t *testing.T) {
	gw := gateway.NewClient(gateway.WithChain(gateway.GOERLI_ID))

	ctx := context.Background()
	var contractClass types.ContractClass

	compiledContract, err := os.ReadFile("/Users/mohaijiang/tmp/my-starkware-erc20-1/cairo-project/ERC20.cairo.starknet.output.json")

	err = json.Unmarshal(compiledContract, &contractClass)

	if err != nil {
		assert.NoError(t, err)
		return
	}

	declare, err := gw.Declare(ctx, contractClass, gateway.DeclareRequest{})
	if err != nil {
		assert.NoError(t, err)
		return
	}

	fmt.Println("declare.TransactionHash: ", declare.TransactionHash)
	fmt.Println("declare.ClassHash: ", declare.ClassHash)

	_, receipt, err := gw.WaitForTransaction(ctx, declare.TransactionHash, 3, 10)
	if err != nil {
		assert.NoError(t, err)
		return
	}

	fmt.Println(receipt)
}

//func TestDeployContract(t *testing.T) {
//	contractService := NewTestContractService()
//	projectId, _ := uuid.FromString("e3a02994-8c27-4539-a9d8-641a823cfaa1")
//	deploy := db.ContractDeploy{
//		ContractId: 55,
//		ProjectId:  projectId,
//	}
//	err := contractService.SaveDeploy(deploy)
//
//	if err != nil {
//		t.Fatalf("deploy contract fail :%v\n", err)
//	}
//
//}

func TestReadToml(t *testing.T) {
	path := "/Users/mohaijiang/workdir/aptos/Move.toml"

	var config vo.Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		panic(err)
	}

	for k, v := range config.Addresses {
		if v == "_" {
			fmt.Println(k)
		}
	}
}

func TestSyncEvmContract(t *testing.T) {
	// 连接以太坊节点
	client, err := ethclient.Dial("https://rpc-moonbeam.hamster.newtouch.com")
	if err != nil {
		log.Fatal(err)
	}

	// 你要查询的合约部署交易的哈希
	transactionHash := common.HexToHash("0x4a20c425a3a11b1a1d0eaceb2169b10b6ec04b71ce810b6670b1e73f83eb8c54")

	// 获取交易的详细信息
	transaction, _, err := client.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		log.Fatal(err)
	}

	// 确认交易是合约创建交易
	if transaction.To() != nil {
		log.Fatalf("Transaction is not a contract deployment")
	}

	// 等待合约部署成功
	receipt, err := waitForContractDeployment(client, transaction.Hash(), 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contract deployed at address: %s\n", receipt.ContractAddress.Hex())
}
