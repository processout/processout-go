package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// CustomerTokens manages the CustomerToken struct
type CustomerTokens struct {
	p *ProcessOut
}

type CustomerToken struct {
	// Gateway : Name of the payment gateway this token was created on
	Gateway string `json:"gateway"`
	// ID : Id of the customer token
	ID string `json:"id"`
	// Name : Name of the token to be displayed
	Name string `json:"name"`
}

// dummyCustomerToken is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomerToken() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
