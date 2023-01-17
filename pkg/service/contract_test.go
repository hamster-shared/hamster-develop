package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestQueryContract(t *testing.T) {
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

	_, err := contractService.QueryContractByWorkflow(1, 1)
	assert.NoError(t, err)
}
