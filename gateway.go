package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Gateway represents the Gateway API object
type Gateway struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the gateway
	ID string `json:"id,omitempty"`
	// Name is the name of the payment gateway
	Name string `json:"name,omitempty"`
	// DisplayName is the name of the payment gateway that can be displayed
	DisplayName string `json:"display_name,omitempty"`
	// LogoURL is the logo URL of the payment gateway
	LogoURL string `json:"logo_url,omitempty"`
	// URL is the uRL of the payment gateway
	URL string `json:"url,omitempty"`
	// Flows is the supported flow by the gateway (one-off, subscription or tokenization)
	Flows []string `json:"flows,omitempty"`
	// Tags is the gateway tags. Mainly used to filter gateways depending on their attributes (e-wallets and such)
	Tags []string `json:"tags,omitempty"`
	// Description is the description of the payment gateway
	Description string `json:"description,omitempty"`
}

// SetClient sets the client for the Gateway object and its
// children
func (s *Gateway) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
}

// dummyGateway is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyGateway() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New(nil, "", "")
}
