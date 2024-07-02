package services

import (
	defaultError "errors"
	"investify/db/aws"
	db "investify/db/sqlc"
	"investify/util"
	"log"
	"mime/multipart"
	"strings"
	"time"

	"investify/api/types"
	"investify/api/types/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService interface {
	CreateUserService(ctx *gin.Context, req types.CreateUserRequest) (types.CreateUserResponse, error)
	LoginUserService(ctx *gin.Context, req types.LoginUserRequest) (types.LoginUserResponse, error)
	UploadImageService(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error)
	GetImageService(ctx *gin.Context) (string, error)
	DeleteImageService(ctx *gin.Context) error
}

type UserServiceImpl struct {
	store db.Store
	s3    *aws.S3Service
}

func NewUserService(store db.Store, s3 *aws.S3Service) UserService {
	return &UserServiceImpl{store: store, s3: s3}
}

func (u *UserServiceImpl) CreateUserService(ctx *gin.Context, req types.CreateUserRequest) (types.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.UserDetails.UserPassword)
	if err != nil {
		return types.CreateUserResponse{}, errors.ErrHashPassword
	}

	var respObject types.CreateUserResponse

	err = u.store.ExecTx(ctx, func(tx *db.Queries) error {
		address, err := tx.CreateAddress(ctx, db.CreateAddressParams{
			AddressStreet:  req.AddressDetails.AddressStreet,
			AddressCity:    req.AddressDetails.AddressCity,
			AddressState:   req.AddressDetails.AddressState,
			AddressCountry: req.AddressDetails.AddressCountry,
			AddressZipcode: req.AddressDetails.AddressZipcode,
		})
		if err != nil {
			return errors.ErrCreateAddress
		}

		user, err := tx.CreateUser(ctx, db.CreateUserParams{
			UserEmail:    req.UserDetails.UserEmail,
			UserPassword: hashedPassword,
			UsersRoleID:  req.UserDetails.UserRoleID,
		})
		if err != nil {
			return errors.ErrCreateUser
		}

		if user.UsersRoleID == 1 {
			owner, err := tx.CreateOwner(ctx, db.CreateOwnerParams{
				OwnerName:      req.ProfileDetails.ProfileName,
				OwnerAddressID: address.AddressID,
				OwnerUserID:    user.UserID,
			})
			if err != nil {
				return errors.ErrCreateOwner
			}
			respObject.ProfileInfo = owner

		} else if user.UsersRoleID == 2 {
			investor, err := tx.CreateInvestor(ctx, db.CreateInvestorParams{
				InvestorName:      req.ProfileDetails.ProfileName,
				InvestorAddressID: address.AddressID,
				InvestorUserID:    user.UserID,
			})
			if err != nil {
				return errors.ErrCreateInvestor
			}
			respObject.ProfileInfo = investor
		}

		respObject.UserInfo = user
		respObject.AddressInfo = address

		return nil // Commit transaction
	})

	if err != nil {
		return types.CreateUserResponse{}, err
	}

	// Return success response
	return respObject, nil
}

func (u *UserServiceImpl) LoginUserService(ctx *gin.Context, req types.LoginUserRequest) (types.LoginUserResponse, error) {
	var loginUserResponse types.LoginUserResponse

	user, err := u.store.GetUserByEmail(ctx, req.UserEmail)
	if err != nil {
		return types.LoginUserResponse{}, errors.ErrUserNotFound
	}
	loginUserResponse.Role = int(user.UsersRoleID)

	if user.UsersRoleID == 2 {
		profile, err := u.store.GetInvestorByUserId(ctx, user.UserID)
		if err != nil {
			return types.LoginUserResponse{}, errors.ErrFailedProfileRetrieval
		}
		loginUserResponse.UserProfileName = profile.InvestorName.String
	} else {
		profile, err := u.store.GetOwnerByUserId(ctx, user.UserID)
		if err != nil {
			return types.LoginUserResponse{}, errors.ErrFailedProfileRetrieval
		}
		loginUserResponse.UserProfileName = profile.OwnerName.String
	}

	err = util.CheckPassword(req.UserPassword, user.UserPassword)
	if err != nil {
		return types.LoginUserResponse{}, errors.ErrIncorrectPassword
	}

	token, err := util.GenerateJWT(user)
	if err != nil {
		return types.LoginUserResponse{}, errors.ErrFailedTokenCreation
	}
	loginUserResponse.AccessToken = token
	uuidToken := uuid.New()

	refreshToken, err := u.store.CreateToken(ctx, db.CreateTokenParams{
		TokenValue:      uuidToken,
		TokenUserID:     user.UserID,
		TokenExpiryDate: pgtype.Timestamptz{Time: time.Now().Add(7 * 24 * time.Hour), Valid: true},
	})

	if err != nil {
		return types.LoginUserResponse{}, errors.ErrFailedTokenCreation
	}
	loginUserResponse.RefreshToken = refreshToken.TokenValue.String()
	return loginUserResponse, nil
}

func (u *UserServiceImpl) LogOutUserService(ctx *gin.Context) error {
	// Maintain a map to store the invalidated token
	// Protect the map with a mutex
	// Invalidate the access token
	return nil
}

func (u *UserServiceImpl) UploadImageService(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	var photoURL string

	err := u.store.ExecTx(ctx, func(tx *db.Queries) error {
		url, err := u.s3.UploadFileAndDeleteTemp(ctx, file, header)
		if err != nil {
			return err
		}
		photoURL = url

		user, err := util.CurrentUser(ctx, u.store)
		if err != nil {
			return err
		}

		err = tx.UpdateUserPhotoURL(ctx, db.UpdateUserPhotoURLParams{
			UserID:         user.UserID,
			UsersPhotoLink: pgtype.Text{String: url, Valid: true},
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return photoURL, nil
}
func (u *UserServiceImpl) GetImageService(ctx *gin.Context) (string, error) {

	user, err := util.CurrentUser(ctx, u.store)
	if err != nil {
		log.Panicln(err)
		return "", err
	}

	log.Println(user.UsersPhotoLink.String)
	return string(user.UsersPhotoLink.String), nil
}
func (u *UserServiceImpl) DeleteImageService(ctx *gin.Context) error {

	user, err := util.CurrentUser(ctx, u.store)
	if err != nil {
		return err
	}
	photoLink := user.UsersPhotoLink.String
	if photoLink == "" {
		return defaultError.New("Photo Link not found May be already deleted")
	}
	arr := strings.Split(photoLink, "/")
	objectKey := arr[3]
	// log.Println(objectKey)

	err = u.s3.DeleteImage(objectKey)
	if err != nil {
		log.Println(err)
		return err
	}
	err = u.store.RemoveUserPhotoLink(ctx, user.UserID)
	if err != nil {
		return err
	}

	return nil
}
