package db

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type TemplateType struct {
	Id          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Type        uint           `json:"type"`
	CreateTime  time.Time      `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	DeleteTime  gorm.DeletedAt `json:"delete_time"`
}

type Template struct {
	Id             uint         `gorm:"primaryKey" json:"id"`
	TemplateTypeId uint         `json:"template_type_id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Audited        bool         `json:"audited"`
	LastVersion    string       `json:"last_version"`
	Logo           string       `json:"logo"`
	Image          string       `json:"image"`
	LabelDisplay   string       `json:"labelDisplay"`
	LanguageType   int          `json:"languageType"`
	CreateTime     time.Time    `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime     time.Time    `json:"update_time"`
	DeleteTime     sql.NullTime `json:"delete_time"`
}

type TemplateDetail struct {
	Id                uint         `gorm:"primaryKey" json:"id"`
	TemplateId        string       `json:"template_id"`
	Name              string       `json:"name"`
	Audited           bool         `json:"audited"`
	Extensions        string       `json:"extensions"`
	Description       string       `json:"description"`
	Examples          string       `json:"examples"`
	Resources         string       `json:"resources"`
	AbiInfo           string       `json:"abiInfo"`
	ByteCode          string       `json:"byteCode"`
	Author            string       `json:"author"`
	RepositoryUrl     string       `json:"repositoryUrl"`
	RepositoryName    string       `json:"repositoryName"`
	Version           string       `json:"version"`
	Branch            string       `json:"branch"`
	CodeSources       string       `json:"codeSources"`
	CreateTime        time.Time    `gorm:"column:create_time;default:current_timestamp" json:"create_time"`
	UpdateTime        time.Time    `json:"update_time"`
	DeleteTime        sql.NullTime `json:"delete_time"`
	Title             string       `json:"title"`
	TemplateType      string       `json:"templateType"`
	ShowUrl           string       `json:"showUrl"`
	TitleDescription  string       `json:"titleDescription"`
	HowUseDescription string       `json:"howUseDescription"`
	LabelDisplay      string       `json:"labelDisplay"`
}
