package controller

import (
	"investify/api/services"
	"investify/api/types"
	"investify/api/types/errors"
	"investify/db/aws"
	db "investify/db/sqlc"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	store   db.Store
	userSrv services.UserService
	s3      *aws.S3Service
}

func NewUserController(store db.Store, userSrv services.UserService, s3 *aws.S3Service) *UserController {
	return &UserController{store: store, userSrv: userSrv, s3: s3}
}

func (u *UserController) Test(c *gin.Context) {
	log.Print("i am in controllers")
	c.JSON(http.StatusAccepted, gin.H{"message": "Everything ok"})
}

//logically
// 1. Parse the request
//initiate the transaction
//2. Extract the adress,user,Profile
//first create Adress
//create the user
//according to the role id in the user create the type of profile investor or owner
//commit the transaction

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req types.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(errors.ErrParsingRequest, http.StatusBadRequest, "position 1"))
		return
	}

	_, err = u.userSrv.CreateUserService(ctx, req) // Delegate creation logic to user service
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.GenerateErrorResponse(err, http.StatusInternalServerError, "position 3"))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Successful Signup"))
}

func (u *UserController) LoginUser(ctx *gin.Context) {
	var req types.LoginUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(errors.ErrParsingRequest, http.StatusBadRequest, ""))
		return
	}

	reqObj, err := u.userSrv.LoginUserService(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errors.GenerateErrorResponse(err, http.StatusUnauthorized, ""))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(reqObj, "Login Successful"))
}
func (u *UserController) LogOut(ctx *gin.Context) {

	// err = u.userSrv.LogOutUserService(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, errors.GenerateErrorResponse(err, http.StatusUnauthorized, ""))
	// 	return
	// }

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Logout Successful"))
}

func (u *UserController) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file is received"})
		return
	}
	defer file.Close()

	// Ensure the directory exists
	url, err := u.userSrv.UploadImageService(ctx, file, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(url, "Image Uploaded"))
}
func (u *UserController) GetImage(ctx *gin.Context) {
	imageUrl, err := u.userSrv.GetImageService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No imageUrl is retreived"})
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(imageUrl, "Url fetched Sucessfully"))

}
func (u *UserController) DeleteImage(ctx *gin.Context) {
	err := u.userSrv.DeleteImageService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(err, http.StatusInternalServerError, ""))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(nil, "Sucessfully Deleted"))

}

func (u *UserController) GetUserProfileDetails(ctx *gin.Context) {
	userDetails, err := u.userSrv.GetUserProfileService(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.GenerateErrorResponse(err, http.StatusInternalServerError, ""))
		return
	}

	ctx.JSON(http.StatusOK, types.GenerateResponse(userDetails, "Sucessfully fetched User Details"))

}
