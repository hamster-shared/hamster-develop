package db

import "time"

type HandlerFailedData struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	InstallationId int64     `json:"installationId"`
	Action         string    `json:"action"`
	CreateTime     time.Time `json:"createTime"`
}

type GitRepo struct {
	Id             int       `gorm:"primaryKey" json:"id"`
	UserId         int64     `json:"userId"`
	RepoId         int64     `json:"repoId"`
	InstallationId int64     `json:"installationId"`
	Name           string    `json:"name"`
	CloneUrl       string    `json:"cloneUrl"`
	SshUrl         string    `json:"sshUrl"`
	DefaultBranch  string    `json:"defaultBranch"`
	Private        bool      `json:"private"`
	Language       string    `json:"language"`
	UpdateAt       time.Time `json:"updateAt"`
	CreateTime     time.Time `json:"createTime"`
}

type RepoPage struct {
	Data     []GitRepo `json:"data"`
	Total    int64     `json:"total"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
}

type GitAppInstall struct {
	Id                  int64     `gorm:"primaryKey" json:"id"`
	InstallUserId       int64     `json:"installUserId"`
	Name                string    `json:"name"`
	RepositorySelection string    `json:"repositorySelection"`
	InstallId           int64     `json:"installId"`
	UserId              int64     `json:"userId"`
	AvatarUrl           string    `json:"avatarUrl"`
	CreateTime          time.Time `json:"createTime"`
}
