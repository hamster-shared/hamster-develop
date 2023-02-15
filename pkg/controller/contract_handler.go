package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hamster-shared/hamster-develop/pkg/application"
	"github.com/hamster-shared/hamster-develop/pkg/consts"
	"github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/copier"
	"time"
)

func (h *HandlerServer) contractDeployInfo(g *gin.Context) {

}

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

func (h *HandlerServer) saveContractDeployInfo(g *gin.Context) {

	var entity parameter.ContractDeployParam
	var contractDeploy db.ContractDeploy
	err := g.BindJSON(&entity)

	if err != nil {
		Fail(err.Error(), g)
		return
	}
	projectId, err := uuid.FromString(entity.ProjectId)
	if err != nil {
		Fail("projectId is empty or invalid", g)
		return
	}

	project, err := h.projectService.GetProject(projectId.String())

	_ = copier.Copy(&contractDeploy, &entity)
	contractService := application.GetBean[*service.ContractService]("contractService")
	contractDeploy.DeployTime = time.Now()
	contractDeploy.ProjectId = projectId

	if project.FrameType == consts.Evm {
		contractDeploy.Status = consts.STATUS_SUCCESS
	}

	err = contractService.SaveDeploy(contractDeploy)
	if err != nil {
		Fail(err.Error(), g)
		return
	}

	Success("", g)
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
