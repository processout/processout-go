package processout

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// ExportLayout represents the ExportLayout API object
type ExportLayout struct {
	// ID is the iD of the export layout
	ID *string `json:"id,omitempty"`
	// Project is the project to which the export layout belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the export layout belongs
	ProjectID *string `json:"project_id,omitempty"`
	// CreatedAt is the date at which the export layout was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Name is the name of the export layout.
	Name *string `json:"name,omitempty"`
	// Type is the type of the export layout.
	Type *string `json:"type,omitempty"`
	// IsDefault is the whether the export layout is the default one for the project. It will be used for automatic exports.
	IsDefault *bool `json:"is_default,omitempty"`
	// Configuration is the configuration of the export layout.
	Configuration *ExportLayoutConfiguration `json:"configuration,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *ExportLayout) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the ExportLayout object and its
// children
func (s *ExportLayout) SetClient(c *ProcessOut) *ExportLayout {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Configuration != nil {
		s.Configuration.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayout) Prefill(c *ExportLayout) *ExportLayout {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.CreatedAt = c.CreatedAt
	s.Name = c.Name
	s.Type = c.Type
	s.IsDefault = c.IsDefault
	s.Configuration = c.Configuration

	return s
}

// ExportLayoutAllParameters is the structure representing the
// additional parameters used to call ExportLayout.All
type ExportLayoutAllParameters struct {
	*Options
	*ExportLayout
}

// All allows you to get all the export layouts.
func (s ExportLayout) All(options ...ExportLayoutAllParameters) (*Iterator, error) {
	return s.AllWithContext(context.Background(), options...)
}

// All allows you to get all the export layouts., passes the provided context to the request
func (s ExportLayout) AllWithContext(ctx context.Context, options ...ExportLayoutAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

	type Response struct {
		ExportLayouts []*ExportLayout `json:"export_layouts"`

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

	path := "/exports/layouts"

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
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

	exportLayoutsList := []Identifiable{}
	for _, o := range payload.ExportLayouts {
		exportLayoutsList = append(exportLayoutsList, o.SetClient(s.client))
	}
	exportLayoutsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    exportLayoutsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"export_layouts"`
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
	return exportLayoutsIterator, nil
}

// ExportLayoutFindParameters is the structure representing the
// additional parameters used to call ExportLayout.Find
type ExportLayoutFindParameters struct {
	*Options
	*ExportLayout
}

// Find allows you to find an export layout by its ID.
func (s ExportLayout) Find(exportLayoutID string, options ...ExportLayoutFindParameters) (*ExportLayout, error) {
	return s.FindWithContext(context.Background(), exportLayoutID, options...)
}

// Find allows you to find an export layout by its ID., passes the provided context to the request
func (s ExportLayout) FindWithContext(ctx context.Context, exportLayoutID string, options ...ExportLayoutFindParameters) (*ExportLayout, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

	type Response struct {
		ExportLayout *ExportLayout `json:"export_layout"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
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

	path := "/exports/layouts/" + url.QueryEscape(exportLayoutID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
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

	payload.ExportLayout.SetClient(s.client)
	return payload.ExportLayout, nil
}

// ExportLayoutFindDefaultParameters is the structure representing the
// additional parameters used to call ExportLayout.FindDefault
type ExportLayoutFindDefaultParameters struct {
	*Options
	*ExportLayout
}

// FindDefault allows you to find the default export layout for given export type.
func (s ExportLayout) FindDefault(exportType string, options ...ExportLayoutFindDefaultParameters) (*ExportLayout, error) {
	return s.FindDefaultWithContext(context.Background(), exportType, options...)
}

// FindDefault allows you to find the default export layout for given export type., passes the provided context to the request
func (s ExportLayout) FindDefaultWithContext(ctx context.Context, exportType string, options ...ExportLayoutFindDefaultParameters) (*ExportLayout, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutFindDefaultParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

	type Response struct {
		ExportLayout *ExportLayout `json:"export_layout"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
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

	path := "/exports/layouts/default/" + url.QueryEscape(exportType) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
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

	payload.ExportLayout.SetClient(s.client)
	return payload.ExportLayout, nil
}

// ExportLayoutCreateParameters is the structure representing the
// additional parameters used to call ExportLayout.Create
type ExportLayoutCreateParameters struct {
	*Options
	*ExportLayout
}

// Create allows you to create a new export layout.
func (s ExportLayout) Create(options ...ExportLayoutCreateParameters) (*ExportLayout, error) {
	return s.CreateWithContext(context.Background(), options...)
}

// Create allows you to create a new export layout., passes the provided context to the request
func (s ExportLayout) CreateWithContext(ctx context.Context, options ...ExportLayoutCreateParameters) (*ExportLayout, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

	type Response struct {
		ExportLayout *ExportLayout `json:"export_layout"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	data := struct {
		*Options
		Name          interface{} `json:"name"`
		Type          interface{} `json:"type"`
		IsDefault     interface{} `json:"is_default"`
		Configuration interface{} `json:"configuration"`
	}{
		Options:       opt.Options,
		Name:          s.Name,
		Type:          s.Type,
		IsDefault:     s.IsDefault,
		Configuration: s.Configuration,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/exports/layouts"

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
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

	payload.ExportLayout.SetClient(s.client)
	return payload.ExportLayout, nil
}

// ExportLayoutUpdateParameters is the structure representing the
// additional parameters used to call ExportLayout.Update
type ExportLayoutUpdateParameters struct {
	*Options
	*ExportLayout
}

// Update allows you to update the export layout.
func (s ExportLayout) Update(exportLayoutID string, options ...ExportLayoutUpdateParameters) (*ExportLayout, error) {
	return s.UpdateWithContext(context.Background(), exportLayoutID, options...)
}

// Update allows you to update the export layout., passes the provided context to the request
func (s ExportLayout) UpdateWithContext(ctx context.Context, exportLayoutID string, options ...ExportLayoutUpdateParameters) (*ExportLayout, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutUpdateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

	type Response struct {
		ExportLayout *ExportLayout `json:"export_layout"`
		HasMore      bool          `json:"has_more"`
		Success      bool          `json:"success"`
		Message      string        `json:"message"`
		Code         string        `json:"error_type"`
	}

	data := struct {
		*Options
		Name          interface{} `json:"name"`
		IsDefault     interface{} `json:"is_default"`
		Configuration interface{} `json:"configuration"`
	}{
		Options:       opt.Options,
		Name:          s.Name,
		IsDefault:     s.IsDefault,
		Configuration: s.Configuration,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/exports/layouts/" + url.QueryEscape(exportLayoutID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.NewNetworkError(err)
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

	payload.ExportLayout.SetClient(s.client)
	return payload.ExportLayout, nil
}

// ExportLayoutDeleteParameters is the structure representing the
// additional parameters used to call ExportLayout.Delete
type ExportLayoutDeleteParameters struct {
	*Options
	*ExportLayout
}

// Delete allows you to delete the export layout.
func (s ExportLayout) Delete(exportLayoutID string, options ...ExportLayoutDeleteParameters) error {
	return s.DeleteWithContext(context.Background(), exportLayoutID, options...)
}

// Delete allows you to delete the export layout., passes the provided context to the request
func (s ExportLayout) DeleteWithContext(ctx context.Context, exportLayoutID string, options ...ExportLayoutDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewExportLayout() method to create a new ExportLayout object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayout)

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

	path := "/exports/layouts/" + url.QueryEscape(exportLayoutID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
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

// dummyExportLayout is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayout() {
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
