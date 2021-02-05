package inquirydata

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gavv/httpexpect"
	"github.com/tinnagorn/my-golang-service-template/router"
)

func TestHTTPInquiryData(t *testing.T) {

	var serviceName = "/inquiry-data"

	handler := router.NewEcho()

	// run server using httptest
	server := httptest.NewServer(handler)

	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, "http://localhost:1323")

	payload := map[string]interface{}{
		"customerID": "100011",
	}

	obj := e.POST(serviceName).
		WithJSON(payload).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("code").ValueEqual("code", 0)
	obj.ContainsKey("message").ValueEqual("message", "Inquiry Success !")

	payload = map[string]interface{}{
		"customerIDs": "100011",
	}

	obj = e.POST(serviceName).
		WithJSON(payload).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("code").ValueNotEqual("code", 0)
}
