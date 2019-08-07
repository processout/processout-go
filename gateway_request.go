package processout

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// GatewayRequest is the struture representing an abstracted payment
// gateway request
type GatewayRequest struct {
	GatewayConfigurationUID string            `json:"gateway_configuration_id"`
	URL                     string            `json:"url"`
	Method                  string            `json:"method"`
	Headers                 map[string]string `json:"headers"`
	Body                    string            `json:"body"`
}

const headerWhitelist = map[string]interface{}{
	"accept":          struct{}{},
	"accept-language": struct{}{},
	"content-type":    struct{}{},
	"referer":         struct{}{},
	"user-agent":      struct{}{},
}

// NewGatewayRequest creates a new GatewayRequest from the given gateway
// configuration ID and request
func NewGatewayRequest(gatewayConfigurationID string,
	req *http.Request) *GatewayRequest {

	h := map[string]string{}
	for n, v := range req.Header {
		if _, ok := headerWhitelist[strings.ToLower(n)]; !ok {
			continue
		}
		h[n] = v[0]
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		body = []byte("")
	}

	// Keep request unaltered
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	req.ContentLength = int64(len(body))

	return &GatewayRequest{
		GatewayConfigurationUID: gatewayConfigurationID,
		URL:                     req.URL.String(),
		Method:                  req.Method,
		Headers:                 h,
		Body:                    string(body),
	}
}

// String encodes the GatewayRequest to a source readable by ProcessOut
func (gr *GatewayRequest) String() string {
	j, _ := json.Marshal(gr)

	return "gway_req_" + base64.StdEncoding.EncodeToString(j)
}
