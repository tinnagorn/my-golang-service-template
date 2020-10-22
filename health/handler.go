package health

import (
	"net/http"

	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) HealthCheck(c echo.Context) error {
	var (
		req       = new(Request)
		requestID = c.Response().Header().Get(echo.HeaderXRequestID)
	)

	if err := c.Bind(req); err != nil {
		// logx.WithID(requestID).Error(err.Error())
		log.Printf("Error : Request ID : " + requestID + " , " + err.Error())
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte("can't not bind request"))
	}

	res, err := h.svc.HealthCheck(requestID, req)

	if err != nil {
		return c.Blob(http.StatusBadRequest, echo.MIMEApplicationJSON, []byte(err.Error()))
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, res)
}
