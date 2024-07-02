package controller

import (
	"investify/api/services"
	db "investify/db/sqlc"
)

type OwnerController struct {
	store    db.Store
	ownerSrv services.OwnerService
}

func NewOwnerController(store db.Store, OwnerSrv services.OwnerService) *OwnerController {
	return &OwnerController{store: store, ownerSrv: OwnerSrv}
}
