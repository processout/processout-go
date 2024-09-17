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

// ExportLayoutConfigurationConfigurationOptionsTime represents the ExportLayoutConfigurationConfigurationOptionsTime API object
type ExportLayoutConfigurationConfigurationOptionsTime struct {
	// Format is the format options for configuration.
	Format *[]string `json:"format,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationConfigurationOptionsTime object and its
// children
func (s *ExportLayoutConfigurationConfigurationOptionsTime) SetClient(c *ProcessOut) *ExportLayoutConfigurationConfigurationOptionsTime {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationConfigurationOptionsTime) Prefill(c *ExportLayoutConfigurationConfigurationOptionsTime) *ExportLayoutConfigurationConfigurationOptionsTime {
	if c == nil {
		return s
	}

	s.Format = c.Format

	return s
}

// dummyExportLayoutConfigurationConfigurationOptionsTime is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationConfigurationOptionsTime() {
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
