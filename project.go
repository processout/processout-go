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

// ProjectRegeneratePrivateKeyParameters is the structure representing the
// additional parameters used to call Project.RegeneratePrivateKey
type ProjectRegeneratePrivateKeyParameters struct {
	*Options
	*Project
}

// RegeneratePrivateKey allows you to regenerate the project private key. Make sure to store the new private key and use it in any future request.
func (s Project) RegeneratePrivateKey(options ...ProjectRegeneratePrivateKeyParameters) (*Project, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectRegeneratePrivateKeyParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		Project *Project `json:"project"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/private-keys"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Project.SetClient(s.client)
	return payload.Project, nil
}

// ProjectFetchParameters is the structure representing the
// additional parameters used to call Project.Fetch
type ProjectFetchParameters struct {
	*Options
	*Project
}

// Fetch allows you to fetch the current project information.
func (s Project) Fetch(options ...ProjectFetchParameters) (*Project, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectFetchParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		Project *Project `json:"project"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/projects/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Project.SetClient(s.client)
	return payload.Project, nil
}

// ProjectSaveParameters is the structure representing the
// additional parameters used to call Project.Save
type ProjectSaveParameters struct {
	*Options
	*Project
}

// Save allows you to save the updated project's attributes.
func (s Project) Save(options ...ProjectSaveParameters) (*Project, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		Project *Project `json:"project"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/projects/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Project.SetClient(s.client)
	return payload.Project, nil
}

// ProjectDeleteParameters is the structure representing the
// additional parameters used to call Project.Delete
type ProjectDeleteParameters struct {
	*Options
	*Project
}

// Delete allows you to delete the project. Be careful! Executing this request will prevent any further interaction with the API that uses this project.
func (s Project) Delete(options ...ProjectDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/projects/{project_id}"

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
	}

	return nil
}

// ProjectFetchSupervisedParameters is the structure representing the
// additional parameters used to call Project.FetchSupervised
type ProjectFetchSupervisedParameters struct {
	*Options
	*Project
}

// FetchSupervised allows you to get all the supervised projects.
func (s Project) FetchSupervised(options ...ProjectFetchSupervisedParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectFetchSupervisedParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		Projects []*Project `json:"projects"`

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

	path := "/supervised-projects"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	projectsList := []Identifiable{}
	for _, o := range payload.Projects {
		projectsList = append(projectsList, o.SetClient(s.client))
	}
	projectsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    projectsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"projects"`
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
		hasMorePrev: false,
	}
	return projectsIterator, nil
}

// ProjectCreateSupervisedParameters is the structure representing the
// additional parameters used to call Project.CreateSupervised
type ProjectCreateSupervisedParameters struct {
	*Options
	*Project
	ApplepaySettings interface{} `json:"applepay_settings"`
}

// CreateSupervised allows you to create a new supervised project.
func (s Project) CreateSupervised(options ...ProjectCreateSupervisedParameters) (*Project, error) {
	if s.client == nil {
		panic("Please use the client.NewProject() method to create a new Project object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectCreateSupervisedParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Project)

	type Response struct {
		Project *Project `json:"project"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	data := struct {
		*Options
		ID                   interface{} `json:"id"`
		Name                 interface{} `json:"name"`
		DefaultCurrency      interface{} `json:"default_currency"`
		DunningConfiguration interface{} `json:"dunning_configuration"`
		ApplepaySettings     interface{} `json:"applepay_settings"`
	}{
		Options:              opt.Options,
		ID:                   s.ID,
		Name:                 s.Name,
		DefaultCurrency:      s.DefaultCurrency,
		DunningConfiguration: s.DunningConfiguration,
		ApplepaySettings:     opt.ApplepaySettings,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/supervised-projects"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return nil, errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return nil, erri
	}

	payload.Project.SetClient(s.client)
	return payload.Project, nil
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
