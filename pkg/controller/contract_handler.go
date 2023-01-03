package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/a-line/pkg/application"
	"github.com/hamster-shared/a-line/pkg/db"
	"github.com/hamster-shared/a-line/pkg/service"
	"strconv"
)

func (h *HandlerServer) contractDeployInfo(g *gin.Context) {

}

func (h *HandlerServer) contractDeployDetailByVersion(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	version := gin.Query("version")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractDeployByVersion(id, version)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) contractInfo(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	version := gin.Query("version")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractByVersion(id, version)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) saveContractDeployInfo(g *gin.Context) {

	var contractDeploy db.ContractDeploy

	err := g.BindJSON(contractDeploy)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	contractService := application.GetBean[*service.ContractService]("contractService")

	err = contractService.SaveDeploy(contractDeploy)
	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success("", g)
}

func (h *HandlerServer) versionList(gin *gin.Context) {
	idStr := gin.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryVersionList(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
