package health

type Request struct {
}
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
