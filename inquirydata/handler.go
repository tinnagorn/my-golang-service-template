package inquirydata

import (
	"encoding/json"
	"net/http"

	log "github.com/tOnkowzl/libs/logx"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) InquiryData(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		log.WithID(requestID).Errorf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res := h.svc.InquiryData(requestID, req)

	if res.Code == 0 {
		log.WithID(requestID).WithField("app_res_code", res.Code).Infof("response success")
	} else {
		log.WithID(requestID).WithField("app_res_code", res.Code).Errorf("response error : %+v", res.Message)
	}

	respJson, err := json.Marshal(res)
	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(`
						{
							"code" : 99,
							"message" : "can't not marshal req on response"
						}
				`),
		)
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, respJson)
}
