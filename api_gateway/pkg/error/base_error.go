package error

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	errDetails "google.golang.org/genproto/googleapis/rpc/errdetails"
)

type Message struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

type Error struct {
	DeviceLang       DeviceLang             `json:"device_language"`
	StatusCode       int                    `json:"-"`
	ErrorCode        string                 `json:"error_code"`
	ErrorMessage     string                 `json:"error_message"`
	ErrorField       string                 `json:"error_field,omitempty"`
	LocalizedMessage Message                `json:"localized_message"`
	Data             map[string]interface{} `json:"data,omitempty"`
	ErrorData        interface{}            `json:"error_data,omitempty"`
}

func (e *Error) Error() string {
	if e.DeviceLang == DEVICE_LANG_ID {
		return e.LocalizedMessage.Indonesia
	}
	return e.LocalizedMessage.English
}

func (e *Error) WithData(data map[string]interface{}) *Error {
	return &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: e.ErrorMessage,
		LocalizedMessage: Message{
			English:   e.LocalizedMessage.English,
			Indonesia: e.LocalizedMessage.Indonesia,
		},
		Data: data,
	}
}

func (e *Error) WithErrorData(errCode string, errdata interface{}) *Error {
	return &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    errCode,
		ErrorMessage: e.ErrorMessage,
		LocalizedMessage: Message{
			English:   e.LocalizedMessage.English,
			Indonesia: e.LocalizedMessage.Indonesia,
		},
		ErrorData: errdata,
	}
}

func (e *Error) WithParameter(parameter ...interface{}) *Error {
	return &Error{
		StatusCode:   e.StatusCode,
		ErrorCode:    e.ErrorCode,
		ErrorMessage: fmt.Sprintf(e.ErrorMessage, parameter...),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(e.LocalizedMessage.English, parameter...),
			Indonesia: fmt.Sprintf(e.LocalizedMessage.Indonesia, parameter...),
		},
	}
}

func (e Error) WithField(field string) *Error {
	e.ErrorField = field
	return &e
}

func NewError(code, message, english, indonesia string) *Error {
	return &Error{
		ErrorCode:    code,
		ErrorMessage: message,
		LocalizedMessage: Message{
			English:   english,
			Indonesia: indonesia,
		},
	}
}

func NewErrorWithStatus(status int, code, message, english, indonesia string) *Error {
	return &Error{
		StatusCode:   status,
		ErrorCode:    code,
		ErrorMessage: message,
		LocalizedMessage: Message{
			English:   english,
			Indonesia: indonesia,
		},
	}
}

func GetInvalidParameterMessage(err *Error, parameter ...interface{}) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, parameter),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, parameter),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, parameter),
		},
	}
}

func GetCustomMessageWithParameters(err *Error, paramID, paramEN string) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, paramEN),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, paramEN),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, paramID),
		},
	}
}

func GetCustomMessageWithArgs(err *Error, args ...interface{}) *Error {
	return &Error{
		StatusCode:   err.StatusCode,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: fmt.Sprintf(err.ErrorMessage, args...),
		LocalizedMessage: Message{
			English:   fmt.Sprintf(err.LocalizedMessage.English, args...),
			Indonesia: fmt.Sprintf(err.LocalizedMessage.Indonesia, args...),
		},
	}
}

func GetUnauthorizedAccess(errCode string) *Error {
	return &Error{
		StatusCode: http.StatusUnauthorized,
		ErrorCode:  errCode,
		LocalizedMessage: Message{
			English:   "user-info not exist",
			Indonesia: "user-info tidak ditemukan",
		},
	}
}

func (e Error) GrpcCode() codes.Code {
	switch e.StatusCode {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusNotImplemented:
		return codes.Unimplemented
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusInternalServerError:
		return codes.Internal
	default:
		return codes.Unknown
	}
}

func getDeviceLanguageFromCtx(ctx context.Context) DeviceLang {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return DEVICE_LANG_EN
	}
	langArr := md.Get(DEVICE_LANG)
	if langArr != nil {
		lang := strings.ToUpper(langArr[0])
		if lang == "ID" {
			return DEVICE_LANG_ID
		} else {
			return DEVICE_LANG_EN
		}
	} else {
		return DEVICE_LANG_EN // Default language
	}
}

func (e *Error) BuildError(ctx context.Context) error {
	e.DeviceLang = getDeviceLanguageFromCtx(ctx)

	// Message error will dynamic base on DeviceLang
	st := status.New(e.GrpcCode(), e.Error())

	// Set Reason as ErrorCode and Domain is StatusCode
	errorCode := &errDetails.ErrorInfo{Reason: e.ErrorCode, Domain: strconv.Itoa(e.StatusCode)}

	// set localization message for error
	en := &errDetails.LocalizedMessage{Locale: DEVICE_LANG_EN, Message: e.LocalizedMessage.English}
	id := &errDetails.LocalizedMessage{Locale: DEVICE_LANG_ID, Message: e.LocalizedMessage.Indonesia}

	st, _ = st.WithDetails(en, id, errorCode)
	return st.Err()
}
