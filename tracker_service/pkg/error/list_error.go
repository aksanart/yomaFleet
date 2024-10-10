package error

import "net/http"

var InternalServerError = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "Svr-5000",
	ErrorMessage: "something is wrong",
	LocalizedMessage: Message{
		English:   "something is wrong",
		Indonesia: "ada masalah",
	},
}

var ErrorValidationRequest = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "Svr-4000",
	ErrorMessage: "not valid request, check your request",
	LocalizedMessage: Message{
		English:   "not valid request, check your request",
		Indonesia: "parameter tidak benar, periksa lagi",
	},
}

var ErrorUnauthorized = &Error{
	StatusCode:   http.StatusUnauthorized,
	ErrorCode:    "Svr-4001",
	ErrorMessage: "can't access",
	LocalizedMessage: Message{
		English:   "can't access",
		Indonesia: "tidak bisa akses",
	},
}
