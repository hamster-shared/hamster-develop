package controller

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) templatesCategory(gin *gin.Context) {
	templateTypeStr := gin.Query("type")
	templateType, err := strconv.Atoi(templateTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.GetTemplateTypeList(templateType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) templates(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	languageTypeStr := gin.DefaultQuery("languageType", "1")
	languageType, err := strconv.Atoi(languageTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	deployTypeStr := gin.DefaultQuery("deployType", "0")
	deployType, err := strconv.Atoi(deployTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.GetTemplatesByTypeId(id, languageType, deployType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) templateDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.GetTemplateDetail(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) frontendTemplateDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.GetFrontendTemplateDetail(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) chainTemplateDetail(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.GetChainTemplateDetail(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
func (h *HandlerServer) templateShow(gin *gin.Context) {
	templateTypeStr := gin.Query("type")
	templateType, err := strconv.Atoi(templateTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	templateDeployType := 0
	languageTypeStr := gin.DefaultQuery("languageType", "1")
	if templateType == 2 {
		templateDeployTypeStr := gin.DefaultQuery("deployType", "1")
		templateDeployType, err = strconv.Atoi(templateDeployTypeStr)
		if err != nil {
			Fail(err.Error(), gin)
			return
		}
		languageTypeStr = gin.DefaultQuery("languageType", "0")
	}
	languageType, err := strconv.Atoi(languageTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.TemplateShow(templateType, languageType, templateDeployType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) templateDownload(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	repoName := gin.Query("repoName")
	if "repoName" == "" {
		Fail("download repo name is empty", gin)
		return
	}
	data := h.templateService.TemplateDownload(id, repoName)
	Success(data, gin)
}

func (h *HandlerServer) getDfxJsonData(gin *gin.Context) {
	projectId := gin.Param("projectId")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	dfxDataService := application.GetBean[*service.DfxDataService]("icpDfxDataService")
	data, err := dfxDataService.QueryDfxJsonDataByProjectId(projectId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) updateDfxJsonData(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	var updateData parameter.UpdateDfxData
	err = gin.BindJSON(&updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	dfxDataService := application.GetBean[*service.DfxDataService]("icpDfxDataService")
	err = dfxDataService.UpdateDfxJsonData(id, updateData.JsonData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}
