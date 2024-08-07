package types

import (
	db "investify/db/sqlc"
)

type BaseHttpResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GenerateResponse(data interface{}, message string) *BaseHttpResponse {
	return &BaseHttpResponse{
		Status:  "success",
		Data:    data,
		Message: message,
	}
}

type CreateUserResponse struct {
	UserInfo    db.BkUser    `json:"user_info"`
	AddressInfo db.BkAddress `json:"address_info"`
	ProfileInfo interface{}  `json:"profile_info"`
}
type LoginUserResponse struct {
	UserProfileName string `json:"user_profile_name"`
	Role            int    `json:"user_role_id"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
}

type CreateBusinessResponse struct {
	BusinessInfo db.BkBusiness `json:"business_info"`
	AddressInfo  db.BkAddress  `json:"address_info"`
}

type GetBusinessResponse struct {
	BusinessInfo db.BkBusiness `json:"business_info"`
	AddressInfo  db.BkAddress  `json:"address_info"`
}
type GetBusinessFeedResponse struct {
	BusinessInfo []db.BkBusiness `json:"business_info"`
}
type GetInvestorResponse struct {
	InvestorInfo db.BkInvestor `json:"investor_info"`
}
type GetInvestorFeedResponse struct {
	InvestorInfo []db.BkInvestor `json:"investor_info"`
}

type UserProfileResponse struct {
	UserEmail       string      `json:"user_email"`
	UserPhoneNumber string      `json:"user_phone_number"`
	UserPhotoLink   string      `json:"user_photo_link"`
	UserRoleId      int         `json:"user_role_id"`
	AddressInfo     AddressType `json:"address_info"`
	ProfileInfo     ProfileInfo `json:"profile_info"`
}
