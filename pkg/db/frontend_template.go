package db

import (
	"database/sql"
	"time"
)

type FrontendTemplateDetail struct {
	Id             uint         `gorm:"primaryKey" json:"id"`
	TemplateId     string       `json:"template_id"`
	Name           string       `json:"name"`
	Audited        bool         `json:"audited"`
	Description    string       `json:"description"`
	Examples       string       `json:"examples"`
	Author         string       `json:"author"`
	RepositoryUrl  string       `json:"repositoryUrl"`
	RepositoryName string       `json:"repositoryName"`
	TemplateType   uint         `json:"templateType"`
	Version        string       `json:"version"`
	Branch         string       `json:"branch"`
	CreateTime     time.Time    `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime     time.Time    `json:"update_time"`
	DeleteTime     sql.NullTime `json:"delete_time"`
}
