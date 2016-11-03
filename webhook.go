package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Webhooks manages the Webhook struct
type Webhooks struct {
	p *ProcessOut
}

type Webhook struct {
	// ID : ID of the recurring invoice
	ID string `json:"id"`
	// Project : Project to which the webhook belongs
	Project *Project `json:"project"`
	// Event : Event the webhook is linked to
	Event *Event `json:"event"`
	// RequestURL : URL to which the webhook will be posted
	RequestURL string `json:"request_url"`
	// RequestMethod : Method used to send the webhook (GET or POST)
	RequestMethod string `json:"request_method"`
	// ResponseBody : The response body the webhook received when sending its payload
	ResponseBody string `json:"response_body"`
	// ResponseCode : The response code the webhook received when sending its payload
	ResponseCode string `json:"response_code"`
	// ResponseHeaders : The response headers the webhook received when sending its payload
	ResponseHeaders string `json:"response_headers"`
	// ResponseTimeMs : The time it took for the webhook to send its payload
	ResponseTimeMs int `json:"response_time_ms"`
	// Status : The status of the webhook. 0: pending, 1: success, 2: error
	Status int `json:"status"`
	// CreatedAt : Date at which the webhook was created
	CreatedAt time.Time `json:"created_at"`
	// ReleaseAt : Date at webhook will be/is released
	ReleaseAt time.Time `json:"release_at"`
}

// dummyWebhook is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyWebhook() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New("")
}
