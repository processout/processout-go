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

// ExternalThreeDS represents the ExternalThreeDS API object
type ExternalThreeDS struct {
	// Xid is the threeDS v1 ID
	Xid *string `json:"xid,omitempty"`
	// TransStatus is the transaction status
	TransStatus *string `json:"trans_status,omitempty"`
	// Eci is the eCI
	Eci *string `json:"eci,omitempty"`
	// Cavv is the authentication value
	Cavv *string `json:"cavv,omitempty"`
	// DsTransID is the dS Transaction ID
	DsTransID *string `json:"ds_trans_id,omitempty"`
	// Version is the threeDS Message version
	Version *string `json:"version,omitempty"`
	// AuthenticationFlow is the authentication flow: one of `frictionless` or 'challenge`
	AuthenticationFlow *string `json:"authentication_flow,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExternalThreeDS object and its
// children
func (s *ExternalThreeDS) SetClient(c *ProcessOut) *ExternalThreeDS {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExternalThreeDS) Prefill(c *ExternalThreeDS) *ExternalThreeDS {
	if c == nil {
		return s
	}

	s.Xid = c.Xid
	s.TransStatus = c.TransStatus
	s.Eci = c.Eci
	s.Cavv = c.Cavv
	s.DsTransID = c.DsTransID
	s.Version = c.Version
	s.AuthenticationFlow = c.AuthenticationFlow

	return s
}

// dummyExternalThreeDS is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExternalThreeDS() {
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
