package processout

import (
	"fmt"
	"net/http"
	"time"
)

var (
	// RequestAPIVersion is the default version of the API used in requests 
	// made with this package
	RequestAPIVersion = "1.4.0.0"
	// Host is the URL where API requests are made
	Host = "https://api.processout.com"

	// DefaultClient sets the HTTP default client used for ProcessOut clients
	DefaultClient = &http.Client{
		Timeout:time.Second * 95,
	}
)

// ProcessOut wraps all the components of the package in a
// single structure
type ProcessOut struct {
	// APIVersion is the version of the API to use
	APIVersion string
	// UserAgent is the UserAgent that will be used to send the request
	UserAgent string
	// ProcessOut project ID
	projectID string
	// ProcessOut project secret key
	projectSecret string

	// HTTPClient used to make requests
	HTTPClient *http.Client
}

// Options represents the options available when doing a request to the
// ProcessOut API
type Options struct {
	IdempotencyKey string `json:"-"`
	Expand         []string `json:"expand"`
	Filter         string `json:"filter"`
	Limit          uint64 `json:"limit"`
	EndBefore      string `json:"end_before"`
	StartAfter     string `json:"start_after"`
	DisableLogging bool `json:"-"`
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    RequestAPIVersion,
		HTTPClient:    DefaultClient,
		projectID:     projectID,
		projectSecret: projectSecret,
	}

	return p
}

func setupRequest(client *ProcessOut, opt *Options, req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", client.APIVersion)
	req.Header.Set("User-Agent", "ProcessOut Go-Bindings/{{ (index .S.Libraries "go").Version }}")
	req.Header.Set("Accept", "application/json")
	if client.UserAgent != "" {
		req.Header.Set("User-Agent", client.UserAgent)
	}
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(client.projectID, client.projectSecret)

	v := req.URL.Query()
	v.Set("filter", opt.Filter)
	v.Set("limit", fmt.Sprint(opt.Limit))
	v.Set("end_before", opt.EndBefore)
	v.Set("start_after", opt.StartAfter)
	for _, e := range opt.Expand {
		v.Add("expand[]", e)
	}
	req.URL.RawQuery = v.Encode()
}

{{- range $k, $v := .V }}
// New{{ $v.Name }} creates a new {{ $v.Name }} object
func (c *ProcessOut) New{{ $v.Name }}(prefill ...*{{ $v.Name }}) *{{ $v.Name }} {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the {{ $v.Name }}, or none.")
	}
	if len(prefill) == 0 {
		return &{{ $v.Name }}{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}
{{- end }}
