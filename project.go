package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v4/errors"
)

// Project represents the Project API object
type Project struct {
	// ID is the iD of the project
	ID *string `json:"id,omitempty"`
	// SupervisorProject is the project used to create this project
	SupervisorProject *Project `json:"supervisor_project,omitempty"`
	// SupervisorProjectID is the iD of the project used to create this project
	SupervisorProjectID *string `json:"supervisor_project_id,omitempty"`
	// APIVersion is the current API version of the project
	APIVersion *APIVersion `json:"api_version,omitempty"`
	// Name is the name of the project
	Name *string `json:"name,omitempty"`
	// LogoURL is the name of the project
	LogoURL *string `json:"logo_url,omitempty"`
	// Email is the email of the project
	Email *string `json:"email,omitempty"`
	// DefaultCurrency is the default currency of the project, used to compute analytics amounts
	DefaultCurrency *string `json:"default_currency,omitempty"`
	// PrivateKey is the private key of the project. Only returned when creating a project
	PrivateKey *string `json:"private_key,omitempty"`
	// DunningConfiguration is the dunning configuration of the project
	DunningConfiguration *[]*DunningAction `json:"dunning_configuration,omitempty"`
	// CreatedAt is the date at which the project was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Project) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Project object and its
// children
func (s *Project) SetClient(c *ProcessOut) *Project {
	if s == nil {
		return s
	}
	s.client = c
	if s.SupervisorProject != nil {
		s.SupervisorProject.SetClient(c)
	}
	if s.APIVersion != nil {
		s.APIVersion.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Project) Prefill(c *Project) *Project {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.SupervisorProject = c.SupervisorProject
	s.SupervisorProjectID = c.SupervisorProjectID
	s.APIVersion = c.APIVersion
	s.Name = c.Name
	s.LogoURL = c.LogoURL
	s.Email = c.Email
	s.DefaultCurrency = c.DefaultCurrency
	s.PrivateKey = c.PrivateKey
	s.DunningConfiguration = c.DunningConfiguration
	s.CreatedAt = c.CreatedAt

	return s
}

// ProjectFetchGatewayConfigurationsParameters is the structure representing the
// additional parameters used to call Project.FetchGatewayConfigurations
type ProjectFetchGatewayConfigurationsParameters struct {
	*Options
	*Project
}

// FetchGatewayConfigurations allows you to get all the gateway configurations of the project
func (s Project) FetchGatewayConfigurations(options ...ProjectFetchGatewayConfigurationsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectFetchGatewayConfigurationsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		GatewayConfigurations []*GatewayConfiguration `json:"gateway_configurations"`

		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
	}{
		Options: opt.Options,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/projects/" + url.QueryEscape(*s.ID) + "/gateway-configurations"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	gatewayConfigurationsList := []Identifiable{}
	for _, o := range payload.GatewayConfigurations {
		gatewayConfigurationsList = append(gatewayConfigurationsList, o.SetClient(s.client))
	}
	gatewayConfigurationsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    gatewayConfigurationsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"gateway_configurations"`
				HasMore bool            `json:"has_more"`
			}{}
			if err := json.NewDecoder(b).Decode(&r); err != nil {
				return false, err
			}
			if err := json.Unmarshal(r.Data, i); err != nil {
				return false, err
			}
			return r.HasMore, nil
		},
		client:      s.client,
		hasMoreNext: payload.HasMore,
		hasMorePrev: true,
	}
	return gatewayConfigurationsIterator, nil
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
