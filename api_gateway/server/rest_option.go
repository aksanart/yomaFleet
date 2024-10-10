package server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	commonError "github.com/aksan/weplus/apigw/pkg/error"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ErrorCustomFormat(ctx context.Context, sm *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s, _ := status.FromError(err)
	customErr := commonError.InternalServerError
	customErr.ErrorMessage = s.Message()
	for _, detail := range s.Details() {
		switch t := detail.(type) {
		case *errdetails.LocalizedMessage:
			if t.Locale == "ID" {
				customErr.LocalizedMessage.Indonesia = t.Message
			} else if t.Locale == "EN" {
				customErr.LocalizedMessage.English = t.Message
			}
		case *errdetails.ErrorInfo:
			if statsuHeader, err := strconv.Atoi(t.Domain); err == nil {
				customErr.StatusCode = statsuHeader
			}
			customErr.ErrorCode = t.Reason
		}
	}
	body, _ := json.Marshal(customErr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(customErr.StatusCode)
	w.Write(body)
}

func CustomMatcherMrg(key string) (string, bool) {
	// in here we can set allow or not header
	return key, true
}

type CustomMarshaler struct{}

func (j *CustomMarshaler) ContentType(_ interface{}) string {
	return "application/json"
}
func (j *CustomMarshaler) Marshal(v interface{}) ([]byte, error) {
	byteData, errMarshal := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}.Marshal(v.(proto.Message))
	if errMarshal != nil {
		return nil, errMarshal
	}
	return byteData, nil
}
func (j *CustomMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
func (j *CustomMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return json.NewDecoder(r)
}
func (j *CustomMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	return json.NewEncoder(w)
}
func (j *CustomMarshaler) Delimiter() []byte {
	return []byte("\n")
}
