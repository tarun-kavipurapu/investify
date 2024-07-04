package services

import (
	"database/sql"
	defaultError "errors"
	"investify/api/types"
	"investify/api/types/errors"
	"investify/db/aws"
	db "investify/db/sqlc"
	"investify/util"
	"log"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type BusinessService interface {
	CreateBusinessService(ctx *gin.Context, req types.CreateBusinessRequest) (types.CreateBusinessResponse, error)
	GetBusinessService(ctx *gin.Context) (types.GetBusinessResponse, error)
	GetBusinessServiceByOwner(ctx *gin.Context) (types.GetBusinessFeedResponse, error)
	GetInvestorFeedService(ctx *gin.Context) (types.GetInvestorFeedResponse, error)
	UploadImageService(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error)
	GetImageService(ctx *gin.Context) ([]types.ImageDetails, []error)

	DeleteImageService(ctx *gin.Context) error
}

type BusinessServiceImpl struct {
	store db.Store
	s3    *aws.S3Service
}

func NewBusinessService(store db.Store, s3 *aws.S3Service) BusinessService {
	return &BusinessServiceImpl{store: store, s3: s3}
}

func (b *BusinessServiceImpl) CreateBusinessService(ctx *gin.Context, req types.CreateBusinessRequest) (types.CreateBusinessResponse, error) {
	//initiate the databse Trasaction
	//cnnect the authentication middleware  with this where role to create business should be the owner
	//Extract the user_id form the acess token
	//check the existance of that userId
	//get owner object with the user id
	//inser the adress the into the dtabse
	//extarct the adress id and insert it in the Business table databse
	//transaction done
	var respObject types.CreateBusinessResponse

	err := b.store.ExecTx(ctx, func(tx *db.Queries) error {

		user, err := util.CurrentUser(ctx, b.store)
		if err != nil {
			return errors.ErrUserNotFound
		}

		owner, err := tx.GetOwnerByUserId(ctx, user.UserID)

		if err != nil {
			return errors.ErrCreateOwner
		}

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
		respObject.AddressInfo = address
		business, err := tx.CreateBusiness(ctx, db.CreateBusinessParams{
			BusinessOwnerID:        owner.OwnerID,
			BusinessOwnerFirstname: req.BusinessDetails.BusinessOwnerFirstname,
			BusinessOwnerLastname:  req.BusinessDetails.BusinessOwnerLastname,
			BusinessEmail:          req.BusinessDetails.BusinessEmail,
			BusinessName:           req.BusinessDetails.BusinessName,
			BusinessContact:        req.BusinessDetails.BusinessContact,
			BusinessAddressID:      address.AddressID,
			BusinessRatings:        req.BusinessDetails.BusinessRatings,
			BusinessMinamount:      req.BusinessDetails.BusinessMinamount,
		})

		if err != nil {
			return errors.ErrCreateBusiness
		}
		respObject.BusinessInfo = business

		return nil //commit transaction
	})

	if err != nil {
		return types.CreateBusinessResponse{}, err
	}

	return respObject, nil
}
func (b *BusinessServiceImpl) GetBusinessService(ctx *gin.Context) (types.GetBusinessResponse, error) {
	idstr := ctx.Param("id")
	// Convert the string ID to int64
	log.Println(idstr)
	var respObject types.GetBusinessResponse
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return types.GetBusinessResponse{}, errors.ErrInvalidID
	}
	business, err := b.store.GetBusinessById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetBusinessResponse{}, errors.ErrGetBusiness
		}
		return types.GetBusinessResponse{}, errors.ErrGetBusiness
	}
	address, err := b.store.GetAddressById(ctx, business.BusinessAddressID)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetBusinessResponse{}, errors.ErrGetBusiness
		}
		return types.GetBusinessResponse{}, errors.ErrGetBusiness
	}
	// b.store.get
	respObject.BusinessInfo = business
	respObject.AddressInfo = address
	return respObject, nil
}

// asingle owner can have mulltiple business
func (b *BusinessServiceImpl) GetBusinessServiceByOwner(ctx *gin.Context) (types.GetBusinessFeedResponse, error) {
	user, err := util.CurrentUser(ctx, b.store)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrUserNotFound
	}
	owner, err := b.store.GetOwnerByUserId(ctx, user.UserID)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrGetBusinessByOwner
	}

	var respObject types.GetBusinessFeedResponse
	business, err := b.store.GetBusinessByOwnerId(ctx, owner.OwnerID)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrGetBusinessByOwner
	}
	for _, elem := range business {
		respObject.BusinessInfo = append(respObject.BusinessInfo, elem)
	}
	return respObject, nil
}
func (b *BusinessServiceImpl) GetInvestorFeedService(ctx *gin.Context) (types.GetInvestorFeedResponse, error) {

	var respObject types.GetInvestorFeedResponse
	investors, err := b.store.GetInvestorFeed(ctx)
	if err != nil {
		return types.GetInvestorFeedResponse{}, errors.ErrGetInvestorFeed
	}
	//filter what to send in the feed
	for _, elem := range investors {
		respObject.InvestorInfo = append(respObject.InvestorInfo, elem)
	}
	return respObject, nil
}
func (b *BusinessServiceImpl) UploadImageService(ctx *gin.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	idstr := ctx.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return "", err
	}
	//perform owner verification
	user, err := util.CurrentUser(ctx, b.store)
	if err != nil {
		return "", err
	}
	owner, err := b.store.GetOwnerByUserId(ctx, user.UserID)
	if err != nil {
		return "", err
	}
	//perform the check here
	business, err := b.store.GetBusinessById(ctx, id)
	if err != nil {
		return "", err
	}
	if owner.OwnerID != business.BusinessOwnerID {
		return "", defaultError.New("You are not the owner of this Business")
	}

	var photoURL string

	err = b.store.ExecTx(ctx, func(tx *db.Queries) error {

		url, err := b.s3.UploadFileAndDeleteTemp(ctx, file, header)
		if err != nil {
			return err
		}
		photoURL = url

		_, err = tx.UploadBusinessImage(ctx, db.UploadBusinessImageParams{
			BusinessID: business.BusinessID,
			ImageUrl:   photoURL,
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
func (b *BusinessServiceImpl) GetImageService(ctx *gin.Context) ([]types.ImageDetails, []error) {
	idstr := ctx.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return nil, []error{err}
	}
	user, err := util.CurrentUser(ctx, b.store)
	if err != nil {
		return nil, []error{err}
	}
	owner, err := b.store.GetOwnerByUserId(ctx, user.UserID)
	if err != nil {
		return nil, []error{err}
	}
	business, err := b.store.GetBusinessById(ctx, id)
	if err != nil {
		return nil, []error{err}
	}
	if owner.OwnerID != business.BusinessOwnerID {
		return nil, []error{defaultError.New("You are not the owner of this Business")}
	}

	imagesInfo, err := b.store.GetImageByBusinessId(ctx, business.BusinessID)
	if err != nil {
		return nil, []error{err}
	}
	var imageDetails []types.ImageDetails
	for _, image := range imagesInfo {
		imageDetails = append(imageDetails, types.ImageDetails{
			ImageID:  image.ImageID,
			ImageUrl: image.ImageUrl,
		})
	}

	return imageDetails, nil
}
func (b *BusinessServiceImpl) DeleteImageService(ctx *gin.Context) error {
	idstr := ctx.Param("id") //this is the imageid
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return err
	}
	//get with image id
	image, err := b.store.GetImageByImageId(ctx, id)
	if err != nil {
		return defaultError.New("Image not found id may be wrong")
	}
	photoLink := image.ImageUrl
	if photoLink == "" {
		return defaultError.New("Photo Link not found May be already deleted")
	}
	arr := strings.Split(photoLink, "/")
	objectKey := arr[3]

	//delete in s3
	err = b.s3.DeleteImage(objectKey)
	if err != nil {
		return err
	}
	//delete in database
	err = b.store.DeleteByImageId(ctx, image.ImageID)
	if err != nil {
		return err
	}
	return nil
}
