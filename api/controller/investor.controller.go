package controller

import (
	"investify/api/services"
	"investify/api/types"
	"investify/api/types/errors"
	db "investify/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvestorController struct {
	store       db.Store
	investorSrv services.InvestorService
}

func NewInvestorController(store db.Store, InvestorSrv services.InvestorService) *InvestorController {
	return &InvestorController{store: store, investorSrv: InvestorSrv}
}

func (i *InvestorController) GetBusinessFeedController(ctx *gin.Context) {
	respObject, err := i.investorSrv.GetBusinessFeedService(ctx)
	// Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business feed"))
}

func (i *InvestorController) GetInvestorByIdController(ctx *gin.Context) {
	respObject, err := i.investorSrv.GetInvestorByIdService(ctx)
	// Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Investor fetched"))
}
