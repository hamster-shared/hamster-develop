package controller

import (
	"github.com/hamster-shared/hamster-develop/pkg/application"
	db2 "github.com/hamster-shared/hamster-develop/pkg/db"
	"github.com/hamster-shared/hamster-develop/pkg/parameter"
	"github.com/hamster-shared/hamster-develop/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerServer) getDfxJsonData(gin *gin.Context) {
	projectId := gin.Param("id")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	data, err := icpService.QueryDfxJsonDataByProjectId(projectId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) isConfigJsonData(gin *gin.Context) {
	projectId := gin.Param("id")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	data := icpService.IsConfigJsonData(projectId)
	Success(data, gin)
}

func (h *HandlerServer) updateDfxJsonData(gin *gin.Context) {
	idStr := gin.Param("dfxId")
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
	icpService := application.GetBean[*service.IcpService]("icpService")
	err = icpService.UpdateDfxJsonData(id, updateData.JsonData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) saveDfxJsonData(gin *gin.Context) {
	projectId := gin.Param("id")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	var updateData parameter.UpdateDfxData
	err := gin.BindJSON(&updateData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	err = icpService.SaveDfxJsonData(projectId, updateData.JsonData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) getCanisterList(gin *gin.Context) {
	projectId := gin.Param("id")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	pageStr := gin.DefaultQuery("page", "1")
	sizeStr := gin.DefaultQuery("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	data, err := icpService.QueryIcpCanisterList(projectId, page, size)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) getCanisterInfo(gin *gin.Context) {
	projectId := gin.Param("id")
	if projectId == "" {
		Fail("project id is empty", gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	data, err := icpService.QueryIcpCanister(projectId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(data, gin)
}

func (h *HandlerServer) getIcpAccount(gin *gin.Context) {
	icpService := application.GetBean[*service.IcpService]("icpService")
	userId, exist := gin.Get("userId")
	if !exist {
		Fail("no login", gin)
		return
	}
	icpAccount, err := icpService.GetIcpAccount(userId.(uint))
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(icpAccount, gin)
}

func (h *HandlerServer) rechargeCanister(gin *gin.Context) {
	var rechargeCanisterParam parameter.RechargeCanisterParam
	err := gin.BindJSON(&rechargeCanisterParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpService := application.GetBean[*service.IcpService]("icpService")
	canisterBalanceVo, err := icpService.RechargeCanister(user.Id, rechargeCanisterParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(canisterBalanceVo, gin)
}

func (h *HandlerServer) rechargeWallet(gin *gin.Context) {
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpService := application.GetBean[*service.IcpService]("icpService")
	walletInfoVo, err := icpService.RechargeWallet(user.Id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(walletInfoVo, gin)
}

func (h *HandlerServer) getWalletInfo(gin *gin.Context) {
	userId, exists := gin.Get("userId")
	if !exists {
		Fail("no login", gin)
		return
	}
	icpService := application.GetBean[*service.IcpService]("icpService")
	walletInfoVo, err := icpService.GetWalletInfo(userId.(uint))
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(walletInfoVo, gin)
}

func (h *HandlerServer) getAccountInfo(gin *gin.Context) {
	icpService := application.GetBean[*service.IcpService]("icpService")
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpInfoVo, err := icpService.GetAccountInfo(user.Id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(icpInfoVo, gin)
}

func (h *HandlerServer) createIdentity(gin *gin.Context) {
	icpService := application.GetBean[*service.IcpService]("icpService")
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpInfoVo, err := icpService.CreateIdentity(user.Id)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(icpInfoVo, gin)
}

func (h *HandlerServer) redeemFaucetCoupon(gin *gin.Context) {
	var redeemFaucetCouponParam parameter.RedeemFaucetCouponParam
	err := gin.BindJSON(&redeemFaucetCouponParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpService := application.GetBean[*service.IcpService]("icpService")
	walletInfoVo, err := icpService.RedeemFaucetCoupon(user.Id, redeemFaucetCouponParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success(walletInfoVo, gin)
}
