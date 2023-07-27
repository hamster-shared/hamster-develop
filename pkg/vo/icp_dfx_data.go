package vo

import (
	"database/sql"
	"github.com/hamster-shared/hamster-develop/pkg/db"
)

type IcpDfxDataVo struct {
	Id        int    `json:"id"`
	ProjectId string `json:"projectId"`
	DfxData   string `json:"dfxData"`
}

type IcpCanisterVo struct {
	Id           int               `json:"id"`
	ProjectId    string            `json:"projectId"`
	CanisterId   string            `json:"canisterId"`
	CanisterName string            `json:"canisterName"`
	Cycles       string            ` json:"cycles"`
	Status       db.CanisterStatus `json:"status"`
	CreateTime   sql.NullTime      `json:"createTime"`
	UpdateTime   sql.NullTime      `json:"updateTime"`
}
