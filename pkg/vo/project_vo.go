package vo

import (
	"time"
)

type ProjectPage struct {
	Data     []ProjectListVo `json:"data"`
	Total    int             `json:"total"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

type ProjectListVo struct {
	Id            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	UserId        int64          `json:"UserId"`
	Type          uint           `json:"type"`
	RepositoryUrl string         `json:"RepositoryUrl"`
	FrameType     int            `json:"frameType"`
	RecentCheck   RecentCheckVo  `json:"recentCheck"`
	RecentBuild   RecentBuildVo  `json:"recentBuild"`
	RecentDeploy  RecentDeployVo `json:"recentDeploy"`
}

type ProjectDetailVo struct {
	Id            uint           `json:"id"`
	Name          string         `json:"name"`
	Type          uint           `json:"type"`
	RepositoryUrl string         `json:"RepositoryUrl"`
	FrameType     int            `json:"frameType"`
	RecentCheck   RecentCheckVo  `json:"recentCheck"`
	RecentBuild   RecentBuildVo  `json:"recentBuild"`
	RecentDeploy  RecentDeployVo `json:"recentDeploy"`
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
}

type RecentDeployVo struct {
	Id         uint      `json:"id"`
	Version    string    `json:"version"`
	DeployTime time.Time `json:"deployTime"`
}

type CreateProjectParam struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	TemplateUrl string `json:"templateUrl"`
	FrameType   int    `json:"frameType"`
	UserId      int64  `json:"userId"`
}

type UpdateProjectParam struct {
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}

type UserAuth struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
