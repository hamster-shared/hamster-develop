package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"strconv"
)

func (h *HandlerServer) contractDeployDetailByVersion(gin *gin.Context) {
	id := gin.Param("id")
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
	id := gin.Param("id")
	version := gin.Param("version")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractByVersion(id, version)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) saveContractDeployIngInfo(g *gin.Context) {

	var entity parameter.ContractDeployIngParam
	err := g.BindJSON(&entity)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	contractService := application.GetBean[*service.ContractService]("contractService")
	err = contractService.SaveDeployIng(entity)
	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success("", g)
}

func (h *HandlerServer) saveContractDeployInfo(g *gin.Context) {

	var entity parameter.ContractDeployParam

	err := g.BindJSON(&entity)

	if err != nil {
		Fail(err.Error(), g)
		return
	}

	contractService := application.GetBean[*service.ContractService]("contractService")

	contractDeployId, err := contractService.SaveDeploy(entity)
	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success(contractDeployId, g)
}

func (h *HandlerServer) contractDeployInfo(gin *gin.Context) {
	id := gin.Param("contractDeployId")
	contractDeployId, err := strconv.Atoi(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	contractService := application.GetBean[*service.ContractService]("contractService")
	contractDeployInfo, err := contractService.GetContractDeployInfo(contractDeployId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(contractDeployInfo, gin)
}
func (h *HandlerServer) versionList(gin *gin.Context) {
	id := gin.Param("id")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryVersionList(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) getVersionAndCodeInfoList(gin *gin.Context) {
	id := gin.Param("id")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryVersionAndCodeInfoList(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) queryContractNameList(gin *gin.Context) {
	id := gin.Param("id")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryContractNameList(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) queryNetworkList(gin *gin.Context) {
	id := gin.Param("id")
	contractService := application.GetBean[*service.ContractService]("contractService")
	data, err := contractService.QueryNetworkList(id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}
