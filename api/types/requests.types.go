package types

import "github.com/jackc/pgx/v5/pgtype"

type CreateUserRequest struct {
	ProfileDetails ProfileInfo
	UserDetails    UserType
	AddressDetails AddressType
}

type LoginUserRequest struct {
	Appkey       string `json:"appkey"`
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserPassword string `json:"user_password" binding:"required"`
}

type CreateBusinessRequest struct {
	BusinessDetails BusinessType
	AddressDetails  AddressType
}

type FilterBusinessRequest struct {
	DomainCode string         `json:"domain_code"`
	StateCode  string         `json:"state_code"`
	MinAmount  pgtype.Numeric `json:"min_amount"`
	MaxAmount  pgtype.Numeric `json:"max_amount"`
}

// type Request struct {
// 	Appkey int               `json:"appkey"`
// 	Data   CreateUserRequest `json:"data"`
// }

// func NewRequest(data interface{}) *Request {
// 	return &Request{
// 		Appkey: 0,
// 		Data:   data,
// 	}
// }
