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

// GatewayConfiguration represents the GatewayConfiguration API object
type GatewayConfiguration struct {
	// ID is the iD of the gateway configuration
	ID *string `json:"id,omitempty"`
	// Project is the project to which the gateway configuration belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the gateway configuration belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Gateway is the gateway that the configuration configures
	Gateway *Gateway `json:"gateway,omitempty"`
	// GatewayID is the iD of the gateway to which the gateway configuration belongs
	GatewayID *int `json:"gateway_id,omitempty"`
	// Name is the name of the gateway configuration
	Name *string `json:"name,omitempty"`
	// DefaultCurrency is the default currency of the gateway configuration
	DefaultCurrency *string `json:"default_currency,omitempty"`
	// Enabled is the define whether or not the gateway configuration is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// PublicKeys is the public keys of the payment gateway configuration (key-value pair)
	PublicKeys *map[string]string `json:"public_keys,omitempty"`
	// CreatedAt is the date at which the gateway configuration was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// EnabledAt is the date at which the gateway configuration was enabled
	EnabledAt *time.Time `json:"enabled_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *GatewayConfiguration) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the GatewayConfiguration object and its
// children
func (s *GatewayConfiguration) SetClient(c *ProcessOut) *GatewayConfiguration {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Gateway != nil {
		s.Gateway.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *GatewayConfiguration) Prefill(c *GatewayConfiguration) *GatewayConfiguration {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Gateway = c.Gateway
	s.GatewayID = c.GatewayID
	s.Name = c.Name
	s.DefaultCurrency = c.DefaultCurrency
	s.Enabled = c.Enabled
	s.PublicKeys = c.PublicKeys
	s.CreatedAt = c.CreatedAt
	s.EnabledAt = c.EnabledAt

	return s
}

// GatewayConfigurationAllParameters is the structure representing the
// additional parameters used to call GatewayConfiguration.All
type GatewayConfigurationAllParameters struct {
	*Options
	*GatewayConfiguration
	ExpandMerchantAccounts interface{} `json:"expand_merchant_accounts"`
}

// All allows you to get all the gateway configurations.
func (s GatewayConfiguration) All(options ...GatewayConfigurationAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewGatewayConfiguration() method to create a new GatewayConfiguration object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayConfigurationAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.GatewayConfiguration)

	type Response struct {
		GatewayConfigurations []*GatewayConfiguration `json:"gateway_configurations"`

		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		ExpandMerchantAccounts interface{} `json:"expand_merchant_accounts"`
	}{
		Options:                opt.Options,
		ExpandMerchantAccounts: opt.ExpandMerchantAccounts,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/gateway-configurations"

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

// GatewayConfigurationFindParameters is the structure representing the
// additional parameters used to call GatewayConfiguration.Find
type GatewayConfigurationFindParameters struct {
	*Options
	*GatewayConfiguration
}

// Find allows you to find a gateway configuration by its ID.
func (s GatewayConfiguration) Find(configurationID string, options ...GatewayConfigurationFindParameters) (*GatewayConfiguration, error) {
	if s.client == nil {
		panic("Please use the client.NewGatewayConfiguration() method to create a new GatewayConfiguration object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayConfigurationFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.GatewayConfiguration)

	type Response struct {
		GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration"`
		HasMore              bool                  `json:"has_more"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
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

	path := "/gateway-configurations/" + url.QueryEscape(configurationID) + ""

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

	payload.GatewayConfiguration.SetClient(s.client)
	return payload.GatewayConfiguration, nil
}

// GatewayConfigurationSaveParameters is the structure representing the
// additional parameters used to call GatewayConfiguration.Save
type GatewayConfigurationSaveParameters struct {
	*Options
	*GatewayConfiguration
	Settings           interface{} `json:"settings"`
	SubAccountsEnabled interface{} `json:"sub_accounts_enabled"`
}

// Save allows you to save the updated gateway configuration attributes and settings.
func (s GatewayConfiguration) Save(options ...GatewayConfigurationSaveParameters) (*GatewayConfiguration, error) {
	if s.client == nil {
		panic("Please use the client.NewGatewayConfiguration() method to create a new GatewayConfiguration object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayConfigurationSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.GatewayConfiguration)

	type Response struct {
		GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration"`
		HasMore              bool                  `json:"has_more"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
	}

	data := struct {
		*Options
		ID                 interface{} `json:"id"`
		Name               interface{} `json:"name"`
		Enabled            interface{} `json:"enabled"`
		DefaultCurrency    interface{} `json:"default_currency"`
		Settings           interface{} `json:"settings"`
		SubAccountsEnabled interface{} `json:"sub_accounts_enabled"`
	}{
		Options:            opt.Options,
		ID:                 s.ID,
		Name:               s.Name,
		Enabled:            s.Enabled,
		DefaultCurrency:    s.DefaultCurrency,
		Settings:           opt.Settings,
		SubAccountsEnabled: opt.SubAccountsEnabled,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/gateway-configurations/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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

	payload.GatewayConfiguration.SetClient(s.client)
	return payload.GatewayConfiguration, nil
}

// GatewayConfigurationDeleteParameters is the structure representing the
// additional parameters used to call GatewayConfiguration.Delete
type GatewayConfigurationDeleteParameters struct {
	*Options
	*GatewayConfiguration
}

// Delete allows you to delete the gateway configuration.
func (s GatewayConfiguration) Delete(options ...GatewayConfigurationDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewGatewayConfiguration() method to create a new GatewayConfiguration object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayConfigurationDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.GatewayConfiguration)

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/gateway-configurations/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
	}

	return nil
}

// GatewayConfigurationCreateParameters is the structure representing the
// additional parameters used to call GatewayConfiguration.Create
type GatewayConfigurationCreateParameters struct {
	*Options
	*GatewayConfiguration
	Settings           interface{} `json:"settings"`
	SubAccountsEnabled interface{} `json:"sub_accounts_enabled"`
}

// Create allows you to create a new gateway configuration.
func (s GatewayConfiguration) Create(gatewayName string, options ...GatewayConfigurationCreateParameters) (*GatewayConfiguration, error) {
	if s.client == nil {
		panic("Please use the client.NewGatewayConfiguration() method to create a new GatewayConfiguration object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := GatewayConfigurationCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.GatewayConfiguration)

	type Response struct {
		GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration"`
		HasMore              bool                  `json:"has_more"`
		Success              bool                  `json:"success"`
		Message              string                `json:"message"`
		Code                 string                `json:"error_type"`
	}

	data := struct {
		*Options
		ID                 interface{} `json:"id"`
		Name               interface{} `json:"name"`
		Enabled            interface{} `json:"enabled"`
		DefaultCurrency    interface{} `json:"default_currency"`
		Settings           interface{} `json:"settings"`
		SubAccountsEnabled interface{} `json:"sub_accounts_enabled"`
	}{
		Options:            opt.Options,
		ID:                 s.ID,
		Name:               s.Name,
		Enabled:            s.Enabled,
		DefaultCurrency:    s.DefaultCurrency,
		Settings:           opt.Settings,
		SubAccountsEnabled: opt.SubAccountsEnabled,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/gateways/" + url.QueryEscape(gatewayName) + "/gateway-configurations"

	req, err := http.NewRequest(
		"POST",
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

	payload.GatewayConfiguration.SetClient(s.client)
	return payload.GatewayConfiguration, nil
}

// dummyGatewayConfiguration is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyGatewayConfiguration() {
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
