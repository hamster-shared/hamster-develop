package db

import (
	"database/sql"
	"time"
)

type ChainTemplateDetail struct {
	Id             uint         `gorm:"primaryKey" json:"id"`
	TemplateId     string       `json:"template_id"`
	Name           string       `json:"name"`
	Audited        bool         `json:"audited"`
	Description    string       `json:"description"`
	Author         string       `json:"author"`
	RepositoryUrl  string       `json:"repositoryUrl"`
	RepositoryName string       `json:"repositoryName"`
	Version        string       `json:"version"`
	Branch         string       `json:"branch"`
	ShowUrl        string       `json:"showUrl"`
	CreateTime     time.Time    `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime     time.Time    `json:"update_time"`
	DeleteTime     sql.NullTime `json:"delete_time"`
}
