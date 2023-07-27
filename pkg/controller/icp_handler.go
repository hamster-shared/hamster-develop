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
	dfxDataService := application.GetBean[*service.DfxDataService]("icpDfxDataService")
	data, err := dfxDataService.QueryDfxJsonDataByProjectId(projectId)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
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
	dfxDataService := application.GetBean[*service.DfxDataService]("icpDfxDataService")
	err = dfxDataService.UpdateDfxJsonData(id, updateData.JsonData)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) getCanisterList(gin *gin.Context) {
	Success("", gin)
}

func (h *HandlerServer) rechargeCanister(gin *gin.Context) {
	var rechargeCanisterParam parameter.RechargeCanisterParam
	err := gin.BindJSON(&rechargeCanisterParam)
	if err != nil {
		Fail(err.Error(), gin)
		return
	}
	Success("", gin)
}

func (h *HandlerServer) rechargeWallet(gin *gin.Context) {
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpService := application.GetBean[*service.IcpService]("icpService")
	walletInfoVo, err := icpService.RechargeWallet(user.Id)
	if err != err {
		Fail(err.Error(), gin)
		return
	}
	Success(walletInfoVo, gin)
}

func (h *HandlerServer) getWalletInfo(gin *gin.Context) {
	userAny, _ := gin.Get("user")
	user, _ := userAny.(db2.User)
	icpService := application.GetBean[*service.IcpService]("icpService")
	walletInfoVo, err := icpService.GetWalletInfo(user.Id)
	if err != err {
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
	if err != err {
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
	if err != err {
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
	if err != err {
		Fail(err.Error(), gin)
		return
	}
	Success(walletInfoVo, gin)
}
