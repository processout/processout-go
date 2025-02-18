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

// ExportLayoutConfigurationAmount represents the ExportLayoutConfigurationAmount API object
type ExportLayoutConfigurationAmount struct {
	// Precision is the chosen precision for the amount fields in the export.
	Precision *string `json:"precision,omitempty"`
	// Separator is the chosen separator for the amount fields in the export.
	Separator *string `json:"separator,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationAmount object and its
// children
func (s *ExportLayoutConfigurationAmount) SetClient(c *ProcessOut) *ExportLayoutConfigurationAmount {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationAmount) Prefill(c *ExportLayoutConfigurationAmount) *ExportLayoutConfigurationAmount {
	if c == nil {
		return s
	}

	s.Precision = c.Precision
	s.Separator = c.Separator

	return s
}

// dummyExportLayoutConfigurationAmount is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationAmount() {
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
