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

// CardSchemeDetails represents the CardSchemeDetails API object
type CardSchemeDetails struct {
	// TransactionID is the scheme transaction ID associated with the card for transaction chaining (e.g. SCA)
	TransactionID *string `json:"transaction_id,omitempty"`
	// TransactionLinkID is the transaction Link Identifier for Mastercard transaction chaining
	TransactionLinkID *string `json:"transaction_link_id,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardSchemeDetails object and its
// children
func (s *CardSchemeDetails) SetClient(c *ProcessOut) *CardSchemeDetails {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardSchemeDetails) Prefill(c *CardSchemeDetails) *CardSchemeDetails {
	if c == nil {
		return s
	}

	s.TransactionID = c.TransactionID
	s.TransactionLinkID = c.TransactionLinkID

	return s
}

// dummyCardSchemeDetails is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardSchemeDetails() {
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
