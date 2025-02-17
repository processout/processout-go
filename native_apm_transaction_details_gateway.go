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

// NativeAPMTransactionDetailsGateway represents the NativeAPMTransactionDetailsGateway API object
type NativeAPMTransactionDetailsGateway struct {
	// DisplayName is the native APM Gateway display name
	DisplayName *string `json:"display_name,omitempty"`
	// LogoURL is the native APM Gateway logo url
	LogoURL *string `json:"logo_url,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the NativeAPMTransactionDetailsGateway object and its
// children
func (s *NativeAPMTransactionDetailsGateway) SetClient(c *ProcessOut) *NativeAPMTransactionDetailsGateway {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *NativeAPMTransactionDetailsGateway) Prefill(c *NativeAPMTransactionDetailsGateway) *NativeAPMTransactionDetailsGateway {
	if c == nil {
		return s
	}

	s.DisplayName = c.DisplayName
	s.LogoURL = c.LogoURL

	return s
}

// dummyNativeAPMTransactionDetailsGateway is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyNativeAPMTransactionDetailsGateway() {
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
