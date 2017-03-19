package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// APIRequest represents the APIRequest API object
type APIRequest struct {
	// ID is the iD of the API request
	ID *string `json:"id,omitempty"`
	// Project is the project used to send the API request
	Project *Project `json:"project,omitempty"`
	// APIVersion is the aPI version used to process the request
	APIVersion *APIVersion `json:"api_version,omitempty"`
	// IdempotencyKey is the idempotency key used to identify the request
	IdempotencyKey *string `json:"idempotency_key,omitempty"`
	// URL is the uRL called
	URL *string `json:"url,omitempty"`
	// Method is the hTTP verb used in the request (GET, POST etc)
	Method *string `json:"method,omitempty"`
	// Headers is the headers sent with the request (client to server)
	Headers *map[string]string `json:"headers,omitempty"`
	// Body is the body of the request (client to server)
	Body *string `json:"body,omitempty"`
	// ResponseCode is the response code (such as 200 for a successful request)
	ResponseCode *int `json:"response_code,omitempty"`
	// ResponseHeaders is the headers sent in the response (server to client)
	ResponseHeaders *map[string]string `json:"response_headers,omitempty"`
	// ResponseBody is the body of the response (client to server)
	ResponseBody *string `json:"response_body,omitempty"`
	// ResponseMs is the number of milliseconds needed to process the request
	ResponseMs *float64 `json:"response_ms,omitempty"`
	// Sandbox is the define whether or not the API request was made in the sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the API request was made
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *APIRequest) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the APIRequest object and its
// children
func (s *APIRequest) SetClient(c *ProcessOut) *APIRequest {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.APIVersion != nil {
		s.APIVersion.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *APIRequest) Prefill(c *APIRequest) *APIRequest {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.APIVersion = c.APIVersion
	s.IdempotencyKey = c.IdempotencyKey
	s.URL = c.URL
	s.Method = c.Method
	s.Headers = c.Headers
	s.Body = c.Body
	s.ResponseCode = c.ResponseCode
	s.ResponseHeaders = c.ResponseHeaders
	s.ResponseBody = c.ResponseBody
	s.ResponseMs = c.ResponseMs
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// APIRequestAllParameters is the structure representing the
// additional parameters used to call APIRequest.All
type APIRequestAllParameters struct {
	*Options
	*APIRequest
}

// All allows you to get all the API requests.
func (s APIRequest) All(options ...APIRequestAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewAPIRequest() method to create a new APIRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := APIRequestAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.APIRequest)

	type Response struct {
		ApiRequests []*APIRequest `json:"api_requests"`

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

	path := "/api-requests"

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

	APIRequestsList := []Identifiable{}
	for _, o := range payload.ApiRequests {
		APIRequestsList = append(APIRequestsList, o.SetClient(s.client))
	}
	APIRequestsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    APIRequestsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"api_requests"`
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
	return APIRequestsIterator, nil
}

// APIRequestFindParameters is the structure representing the
// additional parameters used to call APIRequest.Find
type APIRequestFindParameters struct {
	*Options
	*APIRequest
}

// Find allows you to find an API request by its ID.
func (s APIRequest) Find(APIRequestID string, options ...APIRequestFindParameters) (*APIRequest, error) {
	if s.client == nil {
		panic("Please use the client.NewAPIRequest() method to create a new APIRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := APIRequestFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.APIRequest)

	type Response struct {
		APIRequest *APIRequest `json:"api_request"`
		HasMore    bool        `json:"has_more"`
		Success    bool        `json:"success"`
		Message    string      `json:"message"`
		Code       string      `json:"error_type"`
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

	path := "/api-requests/{request_id}"

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

	payload.APIRequest.SetClient(s.client)
	return payload.APIRequest, nil
}

// dummyAPIRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAPIRequest() {
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
