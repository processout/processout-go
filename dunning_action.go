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

// DunningAction represents the DunningAction API object
type DunningAction struct {
	// Action is the dunning action. Can be either retry, cancel, set_past_due or leave_unchanged
	Action *string `json:"action,omitempty"`
	// DelayInDays is the delay in days that should be waited before executing the next dunning action
	DelayInDays *int `json:"delay_in_days,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the DunningAction object and its
// children
func (s *DunningAction) SetClient(c *ProcessOut) *DunningAction {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *DunningAction) Prefill(c *DunningAction) *DunningAction {
	if c == nil {
		return s
	}

	s.Action = c.Action
	s.DelayInDays = c.DelayInDays

	return s
}

// dummyDunningAction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyDunningAction() {
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
