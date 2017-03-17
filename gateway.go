package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Gateway represents the Gateway API object
type Gateway struct {
	Identifier

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
	// CanPullTransactions is the true if the gateway can pull old transactions into ProcessOut, false otherwise
	CanPullTransactions bool `json:"can_pull_transactions,omitempty"`
	// CanRefund is the true if the gateway supports refunds, false otherwise
	CanRefund bool `json:"can_refund,omitempty"`
	// IsOauthAuthentication is the true if the gateway supports oauth authentication, false otherwise
	IsOauthAuthentication bool `json:"is_oauth_authentication,omitempty"`
	// Description is the description of the payment gateway
	Description string `json:"description,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the Gateway object and its
// children
func (s *Gateway) SetClient(c *ProcessOut) *Gateway {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Gateway) Prefill(c *Gateway) *Gateway {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Name = c.Name
	s.DisplayName = c.DisplayName
	s.LogoURL = c.LogoURL
	s.URL = c.URL
	s.Flows = c.Flows
	s.Tags = c.Tags
	s.CanPullTransactions = c.CanPullTransactions
	s.CanRefund = c.CanRefund
	s.IsOauthAuthentication = c.IsOauthAuthentication
	s.Description = c.Description

	return s
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
