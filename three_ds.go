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

// ThreeDS represents the ThreeDS API object
type ThreeDS struct {
	// Version is the version of the 3DS
	Version *string `json:"version,omitempty"`
	// Status is the current status of the authentication
	Status *string `json:"status,omitempty"`
	// Fingerprinted is the true if a fingerprint has occured
	Fingerprinted *bool `json:"fingerprinted,omitempty"`
	// Challenged is the true if a challenge has occured
	Challenged *bool `json:"challenged,omitempty"`
	// AresTransStatus is the ares transaction status
	AresTransStatus *string `json:"ares_trans_status,omitempty"`
	// CresTransStatus is the cres transaction status
	CresTransStatus *string `json:"cres_trans_status,omitempty"`
	// DsTransID is the universally unique transaction identifier assigned by the DS to identify a single transaction
	DsTransID *string `json:"ds_trans_id,omitempty"`
	// FingerprintCompletionIndicator is the indicates whether the 3DS fingerprint successfully completed
	FingerprintCompletionIndicator *string `json:"fingerprint_completion_indicator,omitempty"`
	// ServerTransID is the universally unique transaction identifier assigned by the 3DS Server to identify a single transaction
	ServerTransID *string `json:"server_trans_id,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ThreeDS object and its
// children
func (s *ThreeDS) SetClient(c *ProcessOut) *ThreeDS {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ThreeDS) Prefill(c *ThreeDS) *ThreeDS {
	if c == nil {
		return s
	}

	s.Version = c.Version
	s.Status = c.Status
	s.Fingerprinted = c.Fingerprinted
	s.Challenged = c.Challenged
	s.AresTransStatus = c.AresTransStatus
	s.CresTransStatus = c.CresTransStatus
	s.DsTransID = c.DsTransID
	s.FingerprintCompletionIndicator = c.FingerprintCompletionIndicator
	s.ServerTransID = c.ServerTransID

	return s
}

// dummyThreeDS is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyThreeDS() {
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
