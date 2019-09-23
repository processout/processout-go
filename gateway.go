package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v4/errors"
)

// Gateway represents the Gateway API object
type Gateway struct {
	// ID is the iD of the gateway
	ID *string `json:"id,omitempty"`
	// Name is the name of the payment gateway
	Name *string `json:"name,omitempty"`
	// DisplayName is the name of the payment gateway that can be displayed
	DisplayName *string `json:"display_name,omitempty"`
	// LogoURL is the logo URL of the payment gateway
	LogoURL *string `json:"logo_url,omitempty"`
	// URL is the uRL of the payment gateway
	URL *string `json:"url,omitempty"`
	// Flows is the supported flow by the gateway (one-off, subscription or tokenization)
	Flows *[]string `json:"flows,omitempty"`
	// Tags is the gateway tags. Mainly used to filter gateways depending on their attributes (e-wallets and such)
	Tags *[]string `json:"tags,omitempty"`
	// CanPullTransactions is the true if the gateway can pull old transactions into ProcessOut, false otherwise
	CanPullTransactions *bool `json:"can_pull_transactions,omitempty"`
	// CanRefund is the true if the gateway supports refunds, false otherwise
	CanRefund *bool `json:"can_refund,omitempty"`
	// IsOauthAuthentication is the true if the gateway supports oauth authentication, false otherwise
	IsOauthAuthentication *bool `json:"is_oauth_authentication,omitempty"`
	// Description is the description of the payment gateway
	Description *string `json:"description,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Gateway) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
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

// GatewayFetchGatewayConfigurationsParameters is the structure representing the
// additional parameters used to call Gateway.FetchGatewayConfigurations
type GatewayFetchGatewayConfigurationsParameters struct {
	*Options
	*Gateway
}

// FetchGatewayConfigurations allows you to get all the gateway configurations of the gateway
func (s Gateway) FetchGatewayConfigurations(options ...GatewayFetchGatewayConfigurationsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewGateway() method to create a new Gateway object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayFetchGatewayConfigurationsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Gateway)

	type Response struct {
		GatewayConfigurations []*GatewayConfiguration `json:"gateway_configurations"`

		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/gateways/" + url.QueryEscape(*s.Name) + "/gateway-configurations"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	gatewayConfigurationsList := []Identifiable{}
	for _, o := range payload.GatewayConfigurations {
		gatewayConfigurationsList = append(gatewayConfigurationsList, o.SetClient(s.client))
	}
	gatewayConfigurationsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    gatewayConfigurationsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"gateway_configurations"`
				HasMore bool            `json:"has_more"`
			}{}
			if err := json.NewDecoder(b).Decode(&r); err != nil {
				return false, err
			}
			if err := json.Unmarshal(r.Data, i); err != nil {
				return false, err
			}
			return r.HasMore, nil
		},
		client:      s.client,
		hasMoreNext: payload.HasMore,
		hasMorePrev: false,
	}
	return gatewayConfigurationsIterator, nil
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
