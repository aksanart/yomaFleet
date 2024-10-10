package model

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

type DeviceLang string

type Message struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

