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

// Projects manages the Project struct
type Projects struct {
	p *ProcessOut
}

type Project struct {
	// ID : ID of the project
	ID string `json:"id"`
	// Name : Name of the project
	Name string `json:"name"`
	// LogoURL : Name of the project
	LogoURL string `json:"logo_url"`
	// Email : Email of the project
	Email string `json:"email"`
	// CreatedAt : Date at which the project was created
	CreatedAt time.Time `json:"created_at"`
}

// GatewayConfigurations : Get all the gateway configurations of the project
func (s Projects) GatewayConfigurations(project *Project, options ...Options) ([]*GatewayConfiguration, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		GatewayConfigurations []*GatewayConfiguration `json:"gateway_configurations"`
		Success               bool                    `json:"success"`
		Message               string                  `json:"message"`
		Code                  string                  `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/projects/" + url.QueryEscape(project.ID) + "/gateway-configurations"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, newError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return nil, erri
	}
	return payload.GatewayConfigurations, nil
}

// Find : Find a project by its ID.
func (s Projects) Find(projectID string, options ...Options) (*Project, *Error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Project `json:"project"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, newError(err)
	}

	path := "/projects/" + url.QueryEscape(projectID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, newError(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, newError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, newError(err)
	}

	if !payload.Success {
		erri := newError(errors.New(payload.Message))
		erri.Code = payload.Code

		return nil, erri
	}
	return &payload.Project, nil
}

// dummyProject is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyProject() {
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
