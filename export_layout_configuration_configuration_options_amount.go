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

// ExportLayoutConfigurationConfigurationOptionsAmount represents the ExportLayoutConfigurationConfigurationOptionsAmount API object
type ExportLayoutConfigurationConfigurationOptionsAmount struct {
	// Precision is the precision options for configuration.
	Precision *[]string `json:"precision,omitempty"`
	// Separator is the separator options for configuration.
	Separator *[]string `json:"separator,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationConfigurationOptionsAmount object and its
// children
func (s *ExportLayoutConfigurationConfigurationOptionsAmount) SetClient(c *ProcessOut) *ExportLayoutConfigurationConfigurationOptionsAmount {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationConfigurationOptionsAmount) Prefill(c *ExportLayoutConfigurationConfigurationOptionsAmount) *ExportLayoutConfigurationConfigurationOptionsAmount {
	if c == nil {
		return s
	}

	s.Precision = c.Precision
	s.Separator = c.Separator

	return s
}

// dummyExportLayoutConfigurationConfigurationOptionsAmount is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationConfigurationOptionsAmount() {
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
