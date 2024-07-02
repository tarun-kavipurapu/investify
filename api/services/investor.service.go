package services

import (
	"database/sql"
	"investify/api/types"
	"investify/api/types/errors"
	db "investify/db/sqlc"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvestorService interface {
	GetBusinessFeedService(ctx *gin.Context) (types.GetBusinessFeedResponse, error)
	GetInvestorByIdService(ctx *gin.Context) (types.GetInvestorResponse, error)
}

type InvestorServiceImpl struct {
	store db.Store
}

func NewInvestorService(store db.Store) InvestorService {
	return &InvestorServiceImpl{store: store}
}

func (i *InvestorServiceImpl) GetInvestorByIdService(ctx *gin.Context) (types.GetInvestorResponse, error) {
	idstr := ctx.Param("id")
	// Convert the string ID to int64
	log.Println(idstr)
	var respObject types.GetInvestorResponse
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return types.GetInvestorResponse{}, errors.ErrInvalidID
	}
	investor, err := i.store.GetInvestorById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.GetInvestorResponse{}, errors.ErrInvestorNotFound
		}
		return types.GetInvestorResponse{}, errors.ErrInvestorNotFound
	}
	respObject.InvestorInfo = investor

	return respObject, nil
}

func (i *InvestorServiceImpl) GetBusinessFeedService(ctx *gin.Context) (types.GetBusinessFeedResponse, error) {
	var respObject types.GetBusinessFeedResponse
	business, err := i.store.GetBusinessFeed(ctx)
	if err != nil {
		return types.GetBusinessFeedResponse{}, errors.ErrBusinessFeed
	}
	// filter what to send in the feed
	for _, elem := range business {
		respObject.BusinessInfo = append(respObject.BusinessInfo, elem)
	}

	return respObject, nil
}
