package vo

import db2 "github.com/hamster-shared/a-line/pkg/db"

type ProjectPage struct {
	Data     []db2.Project `json:"data"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
}

type CreateProjectParam struct {
	Name        string `json:"name"`
	Type        int    `json:"type"`
	TemplateUrl string `json:"templateUrl"`
	FrameType   string `json:"frameType"`
	UserId      int    `json:"userId"`
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
