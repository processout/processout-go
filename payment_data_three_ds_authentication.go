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

// PaymentDataThreeDSAuthentication represents the PaymentDataThreeDSAuthentication API object
type PaymentDataThreeDSAuthentication struct {
	// XID is the authentication XID
	XID *string `json:"XID,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the PaymentDataThreeDSAuthentication object and its
// children
func (s *PaymentDataThreeDSAuthentication) SetClient(c *ProcessOut) *PaymentDataThreeDSAuthentication {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PaymentDataThreeDSAuthentication) Prefill(c *PaymentDataThreeDSAuthentication) *PaymentDataThreeDSAuthentication {
	if c == nil {
		return s
	}

	s.XID = c.XID

	return s
}

// dummyPaymentDataThreeDSAuthentication is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentDataThreeDSAuthentication() {
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
