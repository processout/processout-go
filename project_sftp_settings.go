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

// ProjectSFTPSettings represents the ProjectSFTPSettings API object
type ProjectSFTPSettings struct {
	// Endpoint is the sFTP server endpoint, port is required
	Endpoint *string `json:"endpoint,omitempty"`
	// Username is the sFTP server username
	Username *string `json:"username,omitempty"`
	// Password is the sFTP server password, required when no 'private_key' is passed
	Password *string `json:"password,omitempty"`
	// PrivateKey is the sFTP server private key, required when no 'password' is passed
	PrivateKey *string `json:"private_key,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the ProjectSFTPSettings object and its
// children
func (s *ProjectSFTPSettings) SetClient(c *ProcessOut) *ProjectSFTPSettings {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *ProjectSFTPSettings) Prefill(c *ProjectSFTPSettings) *ProjectSFTPSettings {
	if c == nil {
		return s
	}

	s.Endpoint = c.Endpoint
	s.Username = c.Username
	s.Password = c.Password
	s.PrivateKey = c.PrivateKey

	return s
}

// ProjectSFTPSettingsSaveSftpSettingsParameters is the structure representing the
// additional parameters used to call ProjectSFTPSettings.SaveSftpSettings
type ProjectSFTPSettingsSaveSftpSettingsParameters struct {
	*Options
	*ProjectSFTPSettings
}

// SaveSftpSettings allows you to save the SFTP settings for the project.
func (s ProjectSFTPSettings) SaveSftpSettings(ID string, options ...ProjectSFTPSettingsSaveSftpSettingsParameters) error {
	if s.client == nil {
		panic("Please use the client.NewProjectSFTPSettings() method to create a new ProjectSFTPSettings object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectSFTPSettingsSaveSftpSettingsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ProjectSFTPSettings)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Endpoint   interface{} `json:"endpoint"`
		Username   interface{} `json:"username"`
		Password   interface{} `json:"password"`
		PrivateKey interface{} `json:"private_key"`
	}{
		Options:    opt.Options,
		Endpoint:   s.Endpoint,
		Username:   s.Username,
		Password:   s.Password,
		PrivateKey: s.PrivateKey,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/projects/" + url.QueryEscape(ID) + "/sftp-settings"

	req, err := http.NewRequest(
		"PUT",
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

// ProjectSFTPSettingsDeleteSftpSettingsParameters is the structure representing the
// additional parameters used to call ProjectSFTPSettings.DeleteSftpSettings
type ProjectSFTPSettingsDeleteSftpSettingsParameters struct {
	*Options
	*ProjectSFTPSettings
}

// DeleteSftpSettings allows you to delete the SFTP settings for the project.
func (s ProjectSFTPSettings) DeleteSftpSettings(ID string, options ...ProjectSFTPSettingsDeleteSftpSettingsParameters) error {
	if s.client == nil {
		panic("Please use the client.NewProjectSFTPSettings() method to create a new ProjectSFTPSettings object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProjectSFTPSettingsDeleteSftpSettingsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.ProjectSFTPSettings)

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

	path := "/projects/" + url.QueryEscape(ID) + "/sftp-settings"

	req, err := http.NewRequest(
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

// dummyProjectSFTPSettings is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyProjectSFTPSettings() {
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
