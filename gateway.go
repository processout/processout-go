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

// Gateways manages the Gateway struct
type Gatewaies struct {
	p *ProcessOut
}

type Gateway struct {
	// ID : ID of the gateway
	ID string `json:"id"`
	// Name : Name of the payment gateway
	Name string `json:"name"`
	// DisplayName : Name of the payment gateway that can be displayed
	DisplayName string `json:"display_name"`
	// LogoURL : Logo URL of the payment gateway
	LogoURL string `json:"logo_url"`
	// URL : URL of the payment gateway
	URL string `json:"url"`
	// Flows : Supported flow by the gateway (one-off, subscription or tokenization)
	Flows []string `json:"flows"`
	// Description : Description of the payment gateway
	Description string `json:"description"`
}

// dummyGateway is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyGateway() {
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
