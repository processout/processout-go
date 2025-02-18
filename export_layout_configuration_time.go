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

// ExportLayoutConfigurationTime represents the ExportLayoutConfigurationTime API object
type ExportLayoutConfigurationTime struct {
	// Format is the format of the time fields in the export.
	Format *string `json:"format,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationTime object and its
// children
func (s *ExportLayoutConfigurationTime) SetClient(c *ProcessOut) *ExportLayoutConfigurationTime {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationTime) Prefill(c *ExportLayoutConfigurationTime) *ExportLayoutConfigurationTime {
	if c == nil {
		return s
	}

	s.Format = c.Format

	return s
}

// dummyExportLayoutConfigurationTime is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationTime() {
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
