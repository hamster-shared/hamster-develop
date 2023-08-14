package vo

import (
	"github.com/google/go-github/v48/github"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	uuid "github.com/iris-contrib/go.uuid"
	"time"
)

type ProjectPage struct {
	Data     []ProjectListVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type ProjectListVo struct {
	Id            uuid.UUID     `gorm:"primaryKey" json:"id"`
	Name          string        `json:"name"`
	UserId        int64         `json:"userId"`
	Type          uint          `json:"type"`
	RepositoryUrl string        `json:"repositoryUrl"`
	FrameType     int           `json:"frameType"`
	Branch        string        `json:"branch"`
	DeployType    int           `json:"deployType"`
	LabelDisplay  string        `json:"labelDisplay"`
	GistId        string        `json:"gistId"`
	DefaultFile   string        `json:"defaultFile"`
	RecentCheck   RecentCheckVo `json:"recentCheck"`
	RecentBuild   RecentBuildVo `json:"recentBuild"`
	RecentDeploy  interface{}   `json:"recentDeploy"`
}

type ProjectDetailVo struct {
	Id              uuid.UUID               `json:"id"`
	Name            string                  `json:"name"`
	Type            uint                    `json:"type"`
	RepositoryUrl   string                  `json:"repositoryUrl"`
	FrameType       consts.ProjectFrameType `json:"frameType"` // see# consts.Evm
	Branch          string                  `json:"branch"`
	DeployType      int                     `json:"deployType"`
	GistId          string                  `json:"gistId"`
	DefaultFile     string                  `json:"defaultFile"`
	LabelDisplay    string                  `json:"labelDisplay"`
	RecentCheck     RecentCheckVo           `json:"recentCheck"`
	RecentBuild     RecentBuildVo           `json:"recentBuild"`
	RecentDeploy    interface{}             `json:"recentDeploy"`
	EvmTemplateType uint                    `json:"evmTemplateType"`
}

type RecentCheckVo struct {
	Id         uint      `json:"id"`
	WorkflowId uint      `json:"workflowId"`
	Status     uint      `json:"status"`
	StartTime  time.Time `json:"startTime"`
}

type RecentBuildVo struct {
	Id         uint      `json:"id"`
	WorkflowId uint      `json:"workflowId"`
	Status     uint      `json:"status"`
	StartTime  time.Time `json:"startTime"`
	Version    string    `json:"version"`
}

type RecentDeployVo struct {
	Id         uint      `json:"id"`
	Version    string    `json:"version"`
	DeployTime time.Time `json:"deployTime"`
	Status     uint      `json:"status"` // 1: deploying, 2: success , 3: fail
}

type PackageDeployVo struct {
	Id         uint      `json:"id"`
	WorkflowId uint      `json:"workflowId"`
	PackageId  uint      `json:"packageId"`
	Status     uint      `json:"status"`
	Version    string    `json:"version"`
	StartTime  time.Time `json:"startTime"`
}

type CreateProjectParam struct {
	Name         string                  `json:"name"`
	Type         int                     `json:"type"`
	Branch       string                  `json:"branch"`
	TemplateUrl  string                  `json:"templateUrl"`
	FrameType    consts.ProjectFrameType `json:"frameType"`
	DeployType   int                     `json:"deployType"`
	UserId       int64                   `json:"userId"`
	GistId       string                  `json:"gistId"`
	DefaultFile  string                  `json:"defaultFile"`
	LabelDisplay string                  `json:"labelDisplay"`
}

type UpdateProjectParam struct {
	Name          string `json:"name"`
	UserId        int    `json:"userId"`
	RepositoryUrl string `json:"repositoryUrl"`
}

type UpdateProjectParams struct {
	Params string `json:"params"`
}

type UserAuth struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Config struct {
	Package struct {
		Name    string `toml:"name"`
		Version string `toml: "version"`
	} `toml:"package"`

	Addresses map[string]string `toml:"addresses"`
}

type RepoListPage struct {
	Data     []RepoVo `json:"data"`
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}

type RepoVo struct {
	Name       string           `json:"name"`
	UpdatedAt  github.Timestamp `json:"updatedAt"`
	Language   string           `json:"language"`
	GithubUrl  string           `json:"githubUrl"`
	Visibility string           `json:"Visibility"`
	RepoOwner  string           `json:"repoOwner"`
	Branch     string           `json:"branch"`
}

type RepoFrameType struct {
	Type     consts.ProjectFrameType `json:"type"`
	EvmFrame uint                    `json:"evmFrame"`
}
