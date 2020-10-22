package health

import (
	"encoding/json"

	status_code "github.com/tinnagorn/my-golang-service-template/statuscode"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HealthCheck(requestID string, req *Request) ([]byte, error) {
	result, err := json.Marshal(Response{
		Code:    status_code.Success,
		Message: "Health is ok",
	})
	return result, err
}
