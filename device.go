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

// Device represents the Device API object
type Device struct {
	// RequestOrigin is the request origin. Use "backend" if the request is not coming directly from the frontend
	RequestOrigin *string `json:"request_origin,omitempty"`
	// ID is the device identifier
	ID *string `json:"id,omitempty"`
	// Channel is the device channel. Possible values: "web", "ios", "android", "other"
	Channel *string `json:"channel,omitempty"`
	// IpAddress is the device IP address. Use if request origin is "backend"
	IpAddress *string `json:"ip_address,omitempty"`
	// UserAgent is the device user agent. Use if request origin is "backend"
	UserAgent *string `json:"user_agent,omitempty"`
	// HeaderAccept is the device accept header. Use if request origin is "backend"
	HeaderAccept *string `json:"header_accept,omitempty"`
	// HeaderReferer is the device referer header. Use if request origin is "backend"
	HeaderReferer *string `json:"header_referer,omitempty"`
	// AppColorDepth is the device color depth. Use if request origin is "backend"
	AppColorDepth *int `json:"app_color_depth,omitempty"`
	// AppJavaEnabled is the device Java enabled. Use if request origin is "backend"
	AppJavaEnabled *bool `json:"app_java_enabled,omitempty"`
	// AppLanguage is the device language. Use if request origin is "backend"
	AppLanguage *string `json:"app_language,omitempty"`
	// AppScreenHeight is the device screen height. Use if request origin is "backend"
	AppScreenHeight *int `json:"app_screen_height,omitempty"`
	// AppScreenWidth is the device screen width. Use if request origin is "backend"
	AppScreenWidth *int `json:"app_screen_width,omitempty"`
	// AppTimezoneOffset is the device timezone offset. Use if request origin is "backend"
	AppTimezoneOffset *int `json:"app_timezone_offset,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Device) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Device object and its
// children
func (s *Device) SetClient(c *ProcessOut) *Device {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Device) Prefill(c *Device) *Device {
	if c == nil {
		return s
	}

	s.RequestOrigin = c.RequestOrigin
	s.ID = c.ID
	s.Channel = c.Channel
	s.IpAddress = c.IpAddress
	s.UserAgent = c.UserAgent
	s.HeaderAccept = c.HeaderAccept
	s.HeaderReferer = c.HeaderReferer
	s.AppColorDepth = c.AppColorDepth
	s.AppJavaEnabled = c.AppJavaEnabled
	s.AppLanguage = c.AppLanguage
	s.AppScreenHeight = c.AppScreenHeight
	s.AppScreenWidth = c.AppScreenWidth
	s.AppTimezoneOffset = c.AppTimezoneOffset

	return s
}

// dummyDevice is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyDevice() {
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
