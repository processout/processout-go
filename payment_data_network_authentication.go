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

// PaymentDataNetworkAuthentication represents the PaymentDataNetworkAuthentication API object
type PaymentDataNetworkAuthentication struct {
	// Cavv is the authentication CAVV
	Cavv *string `json:"cavv,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the PaymentDataNetworkAuthentication object and its
// children
func (s *PaymentDataNetworkAuthentication) SetClient(c *ProcessOut) *PaymentDataNetworkAuthentication {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PaymentDataNetworkAuthentication) Prefill(c *PaymentDataNetworkAuthentication) *PaymentDataNetworkAuthentication {
	if c == nil {
		return s
	}

	s.Cavv = c.Cavv

	return s
}

// dummyPaymentDataNetworkAuthentication is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentDataNetworkAuthentication() {
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
