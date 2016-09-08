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

// GatewayConfigurations manages the GatewayConfiguration struct
type GatewayConfigurations struct {
	p *ProcessOut
}

type GatewayConfiguration struct {
	// ID : ID of the gateway configuration
	ID string `json:"id"`
	// Project : Project to which the gateway configuration belongs
	Project *Project `json:"project"`
	// Gateway : Gateway that the configuration configures
	Gateway *Gateway `json:"gateway"`
	// Enabled : Define whether or not the gateway configuration is enabled
	Enabled bool `json:"enabled"`
	// PublicKeys : Public keys of the payment gateway configuration (key-value pair)
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
	errors.New("")
}
