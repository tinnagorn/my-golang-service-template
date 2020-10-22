package inquirydata

type Request struct {
	CustomerID string `json:"customerID" validate:"required"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
