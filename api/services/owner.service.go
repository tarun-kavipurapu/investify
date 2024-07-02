package services

import db "investify/db/sqlc"

type OwnerService interface {
}

type OwnerServiceImpl struct {
	store db.Store
}

func NewOwnerService(store db.Store) OwnerService {
	return &OwnerServiceImpl{store: store}
}
