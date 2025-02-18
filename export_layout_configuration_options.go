package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// ExportLayoutConfigurationOptions represents the ExportLayoutConfigurationOptions API object
type ExportLayoutConfigurationOptions struct {
	// Columns is the columns options for configuration.
	Columns *[]string `json:"columns,omitempty"`
	// Time is the time options for configuration.
	Time *ExportLayoutConfigurationConfigurationOptionsTime `json:"time,omitempty"`
	// Amount is the amount options for configuration.
	Amount *ExportLayoutConfigurationConfigurationOptionsAmount `json:"amount,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ExportLayoutConfigurationOptions object and its
// children
func (s *ExportLayoutConfigurationOptions) SetClient(c *ProcessOut) *ExportLayoutConfigurationOptions {
	if s == nil {
		return s
	}
	s.client = c
	if s.Time != nil {
		s.Time.SetClient(c)
	}
	if s.Amount != nil {
		s.Amount.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ExportLayoutConfigurationOptions) Prefill(c *ExportLayoutConfigurationOptions) *ExportLayoutConfigurationOptions {
	if c == nil {
		return s
	}

	s.Columns = c.Columns
	s.Time = c.Time
	s.Amount = c.Amount

	return s
}

// ExportLayoutConfigurationOptionsFetchParameters is the structure representing the
// additional parameters used to call ExportLayoutConfigurationOptions.Fetch
type ExportLayoutConfigurationOptionsFetchParameters struct {
	*Options
	*ExportLayoutConfigurationOptions
}

// Fetch allows you to fetch export layout configuration options.
func (s ExportLayoutConfigurationOptions) Fetch(exportType string, options ...ExportLayoutConfigurationOptionsFetchParameters) (*ExportLayoutConfigurationOptions, error) {
	if s.client == nil {
		panic("Please use the client.NewExportLayoutConfigurationOptions() method to create a new ExportLayoutConfigurationOptions object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ExportLayoutConfigurationOptionsFetchParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ExportLayoutConfigurationOptions)

	type Response struct {
		ExportLayoutConfigurationOptions *ExportLayoutConfigurationOptions `json:"export_layout_configuration_options"`
		HasMore                          bool                              `json:"has_more"`
		Success                          bool                              `json:"success"`
		Message                          string                            `json:"message"`
		Code                             string                            `json:"error_type"`
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

	path := "/exports/layouts/options/" + url.QueryEscape(exportType) + ""

	req, err := http.NewRequest(
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

	payload.ExportLayoutConfigurationOptions.SetClient(s.client)
	return payload.ExportLayoutConfigurationOptions, nil
}

// dummyExportLayoutConfigurationOptions is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyExportLayoutConfigurationOptions() {
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
