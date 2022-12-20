package db

import (
	"database/sql"
	"fmt"
	"github.com/hamster-shared/a-line/cmd"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cmd.DSN, // data source name
		DefaultStringSize:         256,     // default size for string fields
		DisableDatetimePrecision:  true,    // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,    // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,    // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,   // auto configure based on currently MySQL version

	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                             // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		return
	}

	//db.AutoMigrate(&User{})

	templateType := &TemplateType{
		Name:        "Name",
		Description: "description",
		Type:        1,
		CreateTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	db.Create(templateType)

	fmt.Println(templateType.Id)

}
