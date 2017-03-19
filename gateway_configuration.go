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
	GatewayID *string `json:"gateway_id,omitempty"`
	// Name is the name of the gateway configuration
	Name *string `json:"name,omitempty"`
	// FeeFixed is the fixed fee of the gateway configuration, if specified
	FeeFixed *float64 `json:"fee_fixed,omitempty"`
	// FeePercentage is the percentage fee of the gateway configuration, if specified
	FeePercentage *float64 `json:"fee_percentage,omitempty"`
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
	s.FeeFixed = c.FeeFixed
	s.FeePercentage = c.FeePercentage
	s.DefaultCurrency = c.DefaultCurrency
	s.Enabled = c.Enabled
	s.PublicKeys = c.PublicKeys
	s.CreatedAt = c.CreatedAt
	s.EnabledAt = c.EnabledAt

	return s
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
