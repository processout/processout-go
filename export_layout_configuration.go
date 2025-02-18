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

// ExportLayoutConfiguration represents the ExportLayoutConfiguration API object
type ExportLayoutConfiguration struct {
	// Columns is the columns that will be exported.
	Columns *[]*ExportLayoutConfigurationColumn `json:"columns,omitempty"`
	// Time is the time related configurations.
	Time *ExportLayoutConfigurationTime `json:"time,omitempty"`
	// Amount is the amount related configurations.
	Amount *ExportLayoutConfigurationAmount `json:"amount,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfiguration object and its
// children
func (s *ExportLayoutConfiguration) SetClient(c *ProcessOut) *ExportLayoutConfiguration {
	if s == nil {
		return s
	}
	s.client = c
	if s.Time != nil {
		s.Time.SetClient(c)
	}
	if s.Amount != nil {
		s.Amount.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfiguration) Prefill(c *ExportLayoutConfiguration) *ExportLayoutConfiguration {
	if c == nil {
		return s
	}

	s.Columns = c.Columns
	s.Time = c.Time
	s.Amount = c.Amount

	return s
}

// dummyExportLayoutConfiguration is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfiguration() {
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
