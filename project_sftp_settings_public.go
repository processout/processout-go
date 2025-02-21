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

// ProjectSFTPSettingsPublic represents the ProjectSFTPSettingsPublic API object
type ProjectSFTPSettingsPublic struct {
	// Enabled is the whether the SFTP settings are enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Endpoint is the sFTP server endpoint with port
	Endpoint *string `json:"endpoint,omitempty"`
	// Username is the sFTP server username
	Username *string `json:"username,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ProjectSFTPSettingsPublic object and its
// children
func (s *ProjectSFTPSettingsPublic) SetClient(c *ProcessOut) *ProjectSFTPSettingsPublic {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ProjectSFTPSettingsPublic) Prefill(c *ProjectSFTPSettingsPublic) *ProjectSFTPSettingsPublic {
	if c == nil {
		return s
	}

	s.Enabled = c.Enabled
	s.Endpoint = c.Endpoint
	s.Username = c.Username

	return s
}

// ProjectSFTPSettingsPublicFetchSftpSettingsParameters is the structure representing the
// additional parameters used to call ProjectSFTPSettingsPublic.FetchSftpSettings
type ProjectSFTPSettingsPublicFetchSftpSettingsParameters struct {
	*Options
	*ProjectSFTPSettingsPublic
}

// FetchSftpSettings allows you to fetch the SFTP settings for the project.
func (s ProjectSFTPSettingsPublic) FetchSftpSettings(ID string, options ...ProjectSFTPSettingsPublicFetchSftpSettingsParameters) (*ProjectSFTPSettingsPublic, error) {
	return s.FetchSftpSettingsWithContext(context.Background(), ID, options...)
}

// FetchSftpSettings allows you to fetch the SFTP settings for the project., passes the provided context to the request
func (s ProjectSFTPSettingsPublic) FetchSftpSettingsWithContext(ctx context.Context, ID string, options ...ProjectSFTPSettingsPublicFetchSftpSettingsParameters) (*ProjectSFTPSettingsPublic, error) {
	if s.client == nil {
		panic("Please use the client.NewProjectSFTPSettingsPublic() method to create a new ProjectSFTPSettingsPublic object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectSFTPSettingsPublicFetchSftpSettingsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ProjectSFTPSettingsPublic)

	type Response struct {
		ProjectSFTPSettingsPublic *ProjectSFTPSettingsPublic `json:"sftp_settings"`
		HasMore                   bool                       `json:"has_more"`
		Success                   bool                       `json:"success"`
		Message                   string                     `json:"message"`
		Code                      string                     `json:"error_type"`
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

	path := "/projects/" + url.QueryEscape(ID) + "/sftp-settings"

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

	payload.ProjectSFTPSettingsPublic.SetClient(s.client)
	return payload.ProjectSFTPSettingsPublic, nil
}

// dummyProjectSFTPSettingsPublic is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyProjectSFTPSettingsPublic() {
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
