package vo

import (
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"time"
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
	CreateTime   time.Time         `json:"createTime"`
	UpdateTime   time.Time         `json:"updateTime"`
	Contract     string            `json:"contract"`
}
