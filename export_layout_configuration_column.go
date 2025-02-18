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

// ExportLayoutConfigurationColumn represents the ExportLayoutConfigurationColumn API object
type ExportLayoutConfigurationColumn struct {
	// Name is the name of the column. Must match with supported ones for chosen export type.
	Name *string `json:"name,omitempty"`
	// Rename is the rename of the chosen column if needed.
	Rename *string `json:"rename,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationColumn object and its
// children
func (s *ExportLayoutConfigurationColumn) SetClient(c *ProcessOut) *ExportLayoutConfigurationColumn {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationColumn) Prefill(c *ExportLayoutConfigurationColumn) *ExportLayoutConfigurationColumn {
	if c == nil {
		return s
	}

	s.Name = c.Name
	s.Rename = c.Rename

	return s
}

// dummyExportLayoutConfigurationColumn is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationColumn() {
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
