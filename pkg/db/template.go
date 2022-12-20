package db

import (
	"database/sql"
	"time"
)

type TemplateType struct {
	Id          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        uint         `json:"type"`
	CreateTime  time.Time    `json:"create_time"`
	UpdateTime  time.Time    `json:"update_time"`
	DeleteTime  sql.NullTime `json:"delete_time"`
}

type Template struct {
	Id             uint         `gorm:"primaryKey" json:"id"`
	TemplateTypeId uint         `json:"template_type_id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Audited        bool         `json:"audited"`
	LastVersion    string       `json:"last_version"`
	Logo           string       `json:"logo"`
	CreateTime     time.Time    `json:"create_time"`
	UpdateTime     time.Time    `json:"update_time"`
	DeleteTime     sql.NullTime `json:"delete_time"`
}

type TemplateDetail struct {
	Id            uint   `gorm:"primaryKey" json:"id"`
	TemplateId    string `json:"template_id"`
	MarkdownInfo  string `json:"markdown_info"`
	Name          string `json:"name"`
	Audited       bool   `json:"audited"`
	Extensions    string
	Description   string
	Examples      string
	Resources     string
	AbiInfo       string
	Author        string
	RepositoryUrl string
	Version       string
	Branch        string
	CodeSources   string
	CreateTime    time.Time    `json:"create_time"`
	UpdateTime    time.Time    `json:"update_time"`
	DeleteTime    sql.NullTime `json:"delete_time"`
}
