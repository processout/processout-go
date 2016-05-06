package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// PaymentGatewayPublicKeys manages the PaymentGatewayPublicKey struct
type PaymentGatewayPublicKeys struct {
	p *ProcessOut
}

type PaymentGatewayPublicKey struct {
	// Key : Key name of the public key
	Key string `json:"key"`
	// Value : Key value of the public key
	Value string `json:"value"`
}

// dummyPaymentGatewayPublicKey is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentGatewayPublicKey() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
