package db

import "database/sql"

type IcpDfxData struct {
	Id         int          `json:"id"`
	ProjectId  string       `json:"projectId"`
	DfxData    string       `json:"dfxData"`
	CreateTime sql.NullTime `json:"createTime"`
}
