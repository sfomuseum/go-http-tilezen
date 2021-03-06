package algnhsa

import (
	"encoding/base64"
	"net/http/httptest"

	"github.com/aws/aws-lambda-go/events"
)

const acceptAllContentType = "*/*"

func newAPIGatewayResponse(w *httptest.ResponseRecorder, binaryContentTypes map[string]bool) (events.APIGatewayProxyResponse, error) {
	event := events.APIGatewayProxyResponse{}

	// Set status code.
	event.StatusCode = w.Code

	// Set headers.
	event.MultiValueHeaders = w.Result().Header

	// Set body.
	contentType := w.Header().Get("Content-Type")
	if binaryContentTypes[acceptAllContentType] || binaryContentTypes[contentType] {
		event.Body = base64.StdEncoding.EncodeToString(w.Body.Bytes())
		event.IsBase64Encoded = true
	} else {
		event.Body = w.Body.String()
	}

	return event, nil
}
