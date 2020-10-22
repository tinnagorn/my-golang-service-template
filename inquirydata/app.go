package inquirydata

import (
	status_code "github.com/tinnagorn/my-golang-service-template/statuscode"

	"gopkg.in/go-playground/validator.v9"
)

type Service struct {
	validator *validator.Validate
}

func NewService() *Service {
	validator := validator.New()
	return &Service{validator: validator}
}

func (s *Service) InquiryData(requestID string, req *Request) *Response {
	if err := s.validator.Struct(req); err != nil {
		result := setResponse(status_code.ValidationError, err.Error())
		return result
	}

	result := Response{
		Code:    status_code.Success,
		Message: "Inquiry Success !",
		// Data:    ResponseData{},
	}
	return &result

}

func setResponse(code int, Message string) *Response {
	return &Response{
		Code:    code,
		Message: Message,
	}
}
