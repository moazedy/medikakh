package errorHandler

import "net/http"

type ErrorText map[string]string

type ErrorCode map[string]int

type Errors struct {
	FA ErrorText
	EN ErrorText
}

type ErrorModel struct {
	ErrorString Errors
	ErrorCode   ErrorCode
}

const (
	InternalServerError = "internal_server_error"
	UserDoesNotExists   = "user_does_not_exists"
)

var EngErrors = ErrorText{
	InternalServerError: "internal server error",
	UserDoesNotExists:   "user does not exists",
}

var FaErrors = ErrorText{
	InternalServerError: "خطای داخلی سرور",
	UserDoesNotExists:   "کاربر مورد نظر وجود ندارد",
}

var ErrorCodes = ErrorCode{
	InternalServerError: http.StatusInternalServerError,
	UserDoesNotExists:   http.StatusNotFound,
}

var TheErrors = ErrorModel{
	ErrorString: Errors{
		EN: EngErrors,
		FA: FaErrors,
	},
	ErrorCode: ErrorCodes,
}
