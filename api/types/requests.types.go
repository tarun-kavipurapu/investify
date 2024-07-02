package types

type CreateUserRequest struct {
	ProfileDetails ProfileInfo
	UserDetails    UserType
	AddressDetails AddressType
}

type LoginUserRequest struct {
	Appkey       string `json:"appkey"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
type CreateBusinessRequest struct {
	BusinessDetails BusinessType
	AdressDetails   AddressType
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
