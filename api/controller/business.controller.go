package controller

import (
	"investify/api/services"
	"investify/api/types"
	"investify/api/types/errors"
	"investify/db/aws"
	db "investify/db/sqlc"
	"net/http"

	DefaultError "errors"

	"github.com/gin-gonic/gin"
)

type BusinessController struct {
	store       db.Store
	businessSrv services.BusinessService
	s3          *aws.S3Service
}

func NewBusinessController(store db.Store, BusinessSrv services.BusinessService, s3 *aws.S3Service) *BusinessController {
	return &BusinessController{store: store, businessSrv: BusinessSrv, s3: s3}
}

func (b *BusinessController) CreateBusiness(ctx *gin.Context) {
	var req types.CreateBusinessRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(errors.ErrParsingRequest, http.StatusBadRequest, "position 1"))
		return
	}

	_, err = b.businessSrv.CreateBusinessService(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Business Created Successfully"))
}

func (b *BusinessController) GetBusinessByIdController(ctx *gin.Context) {
	respObject, err := b.businessSrv.GetBusinessService(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business fetched"))
}

func (b *BusinessController) GetBusinessByOwnerController(ctx *gin.Context) {
	respObject, err := b.businessSrv.GetBusinessServiceByOwner(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Business fetched"))
}

func (b *BusinessController) GetInvestorFeedController(ctx *gin.Context) {
	respObject, err := b.businessSrv.GetInvestorFeedService(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(respObject, "Investor feed fetched"))
}

func (b *BusinessController) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	// Ensure the directory exists
	url, err := b.businessSrv.UploadImageService(ctx, file, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(url, "Image Uploaded"))
}

func (b *BusinessController) GetImageByBusinessId(ctx *gin.Context) {
	images, err := b.businessSrv.GetImageService(ctx)

	if err != nil {
		// Convert []error to error by joining error messages
		errorMessage := ""
		for _, e := range err {
			errorMessage += e.Error() + " "
		}
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(DefaultError.New(errorMessage), http.StatusInternalServerError, ""))

		return
	}
	ctx.JSON(http.StatusOK, types.GenerateResponse(images, "Images Fetched"))
}

func (b *BusinessController) DeleteImageById(ctx *gin.Context) {
	err := b.businessSrv.DeleteImageService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(DefaultError.New(err.Error()), http.StatusInternalServerError, ""))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Deleted sucessfully"))
}
func (b *BusinessController) FilterBusinesses(ctx *gin.Context) {
	var req types.FilterBusinessRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(errors.ErrParsingRequest, http.StatusBadRequest, "position 1"))
		return
	}

	businesses, err := b.businessSrv.FilterBusinesses(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "Failed to fetch businesses"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(businesses, "Businesses fetched successfully"))
}
