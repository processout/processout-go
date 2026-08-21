package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// SubmerchantMapping represents the SubmerchantMapping API object
type SubmerchantMapping struct {
	// SubmerchantID is the iD of the ProcessOut submerchant this mapping belongs to
	SubmerchantID *string `json:"submerchant_id,omitempty"`
	// GatewayConfigurationID is the iD of the gateway configuration this mapping applies to
	GatewayConfigurationID *string `json:"gateway_configuration_id,omitempty"`
	// PspSubmerchantID is the submerchant ID at the PSP the gateway configuration connects to
	PspSubmerchantID *string `json:"psp_submerchant_id,omitempty"`
	// CreatedAt is the time at which the mapping was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the SubmerchantMapping object and its
// children
func (s *SubmerchantMapping) SetClient(c *ProcessOut) *SubmerchantMapping {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *SubmerchantMapping) Prefill(c *SubmerchantMapping) *SubmerchantMapping {
	if c == nil {
		return s
	}

	s.SubmerchantID = c.SubmerchantID
	s.GatewayConfigurationID = c.GatewayConfigurationID
	s.PspSubmerchantID = c.PspSubmerchantID
	s.CreatedAt = c.CreatedAt

	return s
}

// dummySubmerchantMapping is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummySubmerchantMapping() {
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
