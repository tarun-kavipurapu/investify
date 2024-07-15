package controller

import (
	"net/http"

	"investify/api/services"
	"investify/api/types"
	"investify/api/types/errors"

	"github.com/gin-gonic/gin"
)

type QuickCodesController struct {
	quickCodesSrv services.QuickCodesService
}

func NewQuickCodesController(quickCodesSrv services.QuickCodesService) *QuickCodesController {
	return &QuickCodesController{quickCodesSrv: quickCodesSrv}
}

func (c *QuickCodesController) GetAllStates(ctx *gin.Context) {
	states, err := c.quickCodesSrv.GetAllStatesService(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "Failed to fetch states"))
		return
	}
	ctx.JSON(http.StatusOK, types.GenerateResponse(states, "Success"))
}

func (c *QuickCodesController) GetAllDomains(ctx *gin.Context) {
	domains, err := c.quickCodesSrv.GetAllDomainsService(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "Failed to fetch domains"))
		return
	}
	ctx.JSON(http.StatusOK, types.GenerateResponse(domains, "Success"))
}
