package vo

import (
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
	RecentCheck   RecentCheckVo `json:"recentCheck"`
	RecentBuild   RecentBuildVo `json:"recentBuild"`
	RecentDeploy  interface{}   `json:"recentDeploy"`
}

type ProjectDetailVo struct {
	Id            uuid.UUID     `json:"id"`
	Name          string        `json:"name"`
	Type          uint          `json:"type"`
	RepositoryUrl string        `json:"repositoryUrl"`
	FrameType     int           `json:"frameType"`
	Branch        string        `json:"branch"`
	RecentCheck   RecentCheckVo `json:"recentCheck"`
	RecentBuild   RecentBuildVo `json:"recentBuild"`
	RecentDeploy  interface{}   `json:"recentDeploy"`
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
}

type PackageDeployVo struct {
	Id         uint      `json:"id"`
	WorkflowId uint      `json:"workflowId"`
	Status     uint      `json:"status"`
	Version    string    `json:"version"`
	StartTime  time.Time `json:"startTime"`
}

type CreateProjectParam struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	Branch      string `json:"branch"`
	TemplateUrl string `json:"templateUrl"`
	FrameType   int    `json:"frameType"`
	UserId      int64  `json:"userId"`
}

type UpdateProjectParam struct {
	Name          string `json:"name"`
	UserId        int    `json:"userId"`
	RepositoryUrl string `json:"repositoryUrl"`
}

type UserAuth struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
