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
	ErrorMessage: "not valid payload, check your payload",
	LocalizedMessage: Message{
		English:   "not valid payload, check your payload",
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
var ErrorNotFound = &Error{
	StatusCode:   http.StatusNotFound,
	ErrorCode:    "Svr-4004",
	ErrorMessage: "not found data",
	LocalizedMessage: Message{
		English:   "not found data",
		Indonesia: "data tidak ditemukan",
	},
}

var ErrorSessionExpired = &Error{
	StatusCode:   http.StatusUnauthorized,
	ErrorCode:    "Svr-4004",
	ErrorMessage: "session expired, please login",
	LocalizedMessage: Message{
		English:   "session expired, please login",
		Indonesia: "harap login lagi",
	},
}

var ErrorRequiredSession = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "Svr-4002",
	ErrorMessage: "session is required, check your session",
	LocalizedMessage: Message{
		English:   "session is required, check your session",
		Indonesia: "session id perlu diisi",
	},
}

var ErrorUserRegistered = &Error{
	StatusCode:   http.StatusUnauthorized,
	ErrorCode:    "Svr-4001",
	ErrorMessage: "user already registered",
	LocalizedMessage: Message{
		English:   "user already registered",
		Indonesia: "user sudah terdaftar",
	},
}

var ErrorStream = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "STRM-500",
	ErrorMessage: "Stream Error",
	LocalizedMessage: Message{
		English:   "Stream Error",
		Indonesia: "Stream bermasalah",
	},
}

var ErrorNoData = &Error{
	StatusCode:   http.StatusNotFound,
	ErrorCode:    "STRM-404",
	ErrorMessage: "no data found",
	LocalizedMessage: Message{
		English:   "no data found",
		Indonesia: "tidak ada data",
	},
}
