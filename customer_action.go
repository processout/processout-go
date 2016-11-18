package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// CustomerAction represents the CustomerAction API object
type CustomerAction struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// Type is the customer action type (such as url)
	Type string `json:"type"`
	// Value is the value of the customer action. If type is an URL, URL to which you should redirect your customer
	Value string `json:"value"`
}

func (s *CustomerAction) setClient(c *ProcessOut) {
	s.Client = c
}

// dummyCustomerAction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomerAction() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New(nil, "", "")
}
