package controller

import (
	"investify/api/types"
	"investify/db/aws"
	db "investify/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExtraController struct {
	store db.Store
	s3    *aws.S3Service
}

func NewExtraController(store db.Store, s3 *aws.S3Service) *ExtraController {
	return &ExtraController{store: store, s3: s3}
}

func (e *ExtraController) UploadImage(ctx *gin.Context) {
	log.Println("i am before form upload")

	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()
	//ensure the directory exists

	url, err := e.s3.UploadFileAndDeleteTemp(ctx, file, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(url, "Image Uploaded"))
}
