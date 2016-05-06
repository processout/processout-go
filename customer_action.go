package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// CustomerActions manages the CustomerAction struct
type CustomerActions struct {
	p *ProcessOut
}

type CustomerAction struct {
	// URL : URL to which you may redirect the customer
	URL string `json:"url"`
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
	}
	errors.New("")
}
