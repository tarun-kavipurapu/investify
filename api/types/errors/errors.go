package errors

import (
	"errors"
	"net/http"
)

var (
	ErrUserNotFound           = errors.New("user not found")
	ErrIncorrectPassword      = errors.New("incorrect password")
	ErrFailedTokenCreation    = errors.New("failed to create token")
	ErrFailedProfileRetrieval = errors.New("failed to retrieve user profile")
	ErrParsingRequest         = errors.New("error in parsing the user request")
	ErrUserService            = errors.New("error in the user service")
	ErrUnauthorized           = errors.New("unauthorized access")
	ErrHashPassword           = errors.New("failed to hash password")
	ErrCreateAddress          = errors.New("failed to create address")
	ErrCreateUser             = errors.New("Duplicate User found")
	ErrCreateOwner            = errors.New("failed to create owner")
	ErrCreateInvestor         = errors.New("failed to create investor")
	ErrInvalidID              = errors.New("invalid ID format")
	ErrInvestorNotFound       = errors.New("investor not found")
	ErrBusinessFeed           = errors.New("failed to retrieve business feed")
	ErrCreateBusiness         = errors.New("failed to create business")
	ErrGetBusiness            = errors.New("failed to get business")
	ErrGetBusinessByOwner     = errors.New("failed to get business by owner")
	ErrGetInvestorFeed        = errors.New("failed to get investor feed")
	ErrGetAddress             = errors.New("failed to get Address feed")
	// ErrProfileImage           = errors.New("Error UploadingImages")
)

type BaseErrorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
	Extra   string `json:"extra,omitempty"`
}

// add error codes
func GenerateErrorResponse(err error, code int, extra string) *BaseErrorResponse {
	if extra == "" {
		extra = "" // Set extra to an empty string if it's not provided
	}

	message := "An error occurred"
	switch err {
	case ErrUserNotFound:
		message = "The email address provided does not match any account"
		code = http.StatusNotFound
	case ErrIncorrectPassword:
		message = "The password you entered is incorrect"
		code = http.StatusUnauthorized
	case ErrFailedTokenCreation:
		message = "There was an error generating the token"
		code = http.StatusInternalServerError
	case ErrFailedProfileRetrieval:
		message = "Failed to retrieve user profile"
		code = http.StatusInternalServerError
	case ErrParsingRequest:
		message = "Invalid request payload"
		code = http.StatusBadRequest
	case ErrUserService:
		message = "User service error"
		code = http.StatusInternalServerError
	case ErrUnauthorized:
		message = "Unauthorized access"
		code = http.StatusUnauthorized
	case ErrHashPassword:
		message = "Failed to hash password"
		code = http.StatusInternalServerError
	case ErrCreateAddress:
		message = "Failed to create address"
		code = http.StatusInternalServerError
	case ErrCreateUser:
		message = "Failed to create user"
		code = http.StatusInternalServerError
	case ErrCreateOwner:
		message = "Failed to create owner"
		code = http.StatusInternalServerError
	case ErrCreateInvestor:
		message = "Failed to create investor"
		code = http.StatusInternalServerError

	case ErrGetAddress:
		message = "Failed to get Address"
		code = http.StatusInternalServerError
	default:
		message = err.Error()
		code = http.StatusInternalServerError
	}

	return &BaseErrorResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Detail:  err.Error(),
		Extra:   extra,
	}
}
