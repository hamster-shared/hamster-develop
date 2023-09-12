package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
)

func (h *HandlerServer) saveContractArrange(gin *gin.Context) {
	var contractArrangeParam parameter.ContractArrangeParam
	err := gin.BindJSON(&contractArrangeParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	arrangeService := application.GetBean[*service.ArrangeService]("arrangeService")
	arrangeId, err := arrangeService.SaveContractArrange(contractArrangeParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(arrangeId, gin)
}

func (h *HandlerServer) saveContractArrangeExecute(gin *gin.Context) {
	var contractArrangeExecuteParam parameter.ContractArrangeExecuteParam
	err := gin.BindJSON(&contractArrangeExecuteParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	arrangeService := application.GetBean[*service.ArrangeService]("arrangeService")
	arrangeId, err := arrangeService.SaveContractArrangeExecute(contractArrangeExecuteParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(arrangeId, gin)
}

func (h *HandlerServer) updateContractArrangeExecute(gin *gin.Context) {
	var update parameter.ContractArrangeExecuteUpdateParam
	err := gin.BindJSON(&update)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	arrangeService := application.GetBean[*service.ArrangeService]("arrangeService")
	arrangeId, err := arrangeService.UpdateContractArrangeExecute(update)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(arrangeId, gin)
}

func (h *HandlerServer) getContractArrangeExecuteInfo(gin *gin.Context) {
	executeId := gin.Param("executeId")
	if executeId == "" {
		Fail("execute id is empty", gin)
		return
	}
	arrangeService := application.GetBean[*service.ArrangeService]("arrangeService")
	arrangeId, err := arrangeService.GetContractArrangeExecuteInfo(executeId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(arrangeId, gin)
}

func (h *HandlerServer) getDeployArrangeContractList(gin *gin.Context) {
	id := gin.Param("id")
	if id == "" {
		Fail("project id is empty", gin)
		return
	}
	version := gin.Param("version")
	if version == "" {
		Fail("version is empty", gin)
		return
	}
	arrangeService := application.GetBean[*service.ArrangeService]("arrangeService")
	arrangeId, err := arrangeService.GetDeployArrangeContractList(id, version)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(arrangeId, gin)
}
