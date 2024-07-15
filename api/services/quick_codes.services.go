package services

import (
	db "investify/db/sqlc"

	"github.com/gin-gonic/gin"
)

type QuickCodesService interface {
	GetAllStatesService(ctx *gin.Context) ([]db.GetAllStatesRow, error)
	GetAllDomainsService(ctx *gin.Context) ([]db.GetAllDomainsRow, error)
}

type QuickCodesServiceImpl struct {
	store db.Store
}

func NewQuickCodesService(store db.Store) QuickCodesService {
	return &QuickCodesServiceImpl{store: store}
}

func (s *QuickCodesServiceImpl) GetAllStatesService(ctx *gin.Context) ([]db.GetAllStatesRow, error) {
	states, err := s.store.GetAllStates(ctx)
	if err != nil {
		return nil, err
	}
	return states, nil
}

func (s *QuickCodesServiceImpl) GetAllDomainsService(ctx *gin.Context) ([]db.GetAllDomainsRow, error) {
	domains, err := s.store.GetAllDomains(ctx)
	if err != nil {
		return nil, err
	}
	return domains, nil
}
