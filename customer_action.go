package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CustomerActions manages the CustomerAction struct
type CustomerActions struct {
	p *ProcessOut
}

type CustomerAction struct {
	// Type : Customer action type (such as url)
	Type string `json:"type"`
	// Value : Value of the customer action. If type is an URL, URL to which you should redirect your customer
	Value string `json:"value"`
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
	errors.New("")
}
