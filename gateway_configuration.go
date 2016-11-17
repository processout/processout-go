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

// GatewayConfiguration represents the GatewayConfiguration API object
type GatewayConfiguration struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the gateway configuration
	ID string `json:"id"`
	// Project is the project to which the gateway configuration belongs
	Project *Project `json:"project"`
	// Gateway is the gateway that the configuration configures
	Gateway *Gateway `json:"gateway"`
	// Enabled is the define whether or not the gateway configuration is enabled
	Enabled bool `json:"enabled"`
	// PublicKeys is the public keys of the payment gateway configuration (key-value pair)
	PublicKeys map[string]string `json:"public_keys"`
}

// dummyGatewayConfiguration is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyGatewayConfiguration() {
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
