package controller

import (
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
	data, err := h.templateService.GetTemplatesByTypeId(id)
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
func (h *HandlerServer) templateShow(gin *gin.Context) {
	templateTypeStr := gin.Query("type")
	templateType, err := strconv.Atoi(templateTypeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	data, err := h.templateService.TemplateShow(templateType)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
