package processout

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// Card represents the Card API object
type Card struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the card
	ID string `json:"id,omitempty"`
	// Project is the project to which the card belongs
	Project *Project `json:"project,omitempty"`
	// Scheme is the scheme of the card, such as visa or mastercard
	Scheme string `json:"scheme,omitempty"`
	// Type is the type of the card (Credit, Debit, ...)
	Type string `json:"type,omitempty"`
	// BankName is the name of the bank of the card
	BankName string `json:"bank_name,omitempty"`
	// Brand is the level of the card (Electron, Classic, Gold, ...)
	Brand string `json:"brand,omitempty"`
	// Iin is the first 6 digits of the card
	Iin string `json:"iin,omitempty"`
	// Last4Digits is the last 4 digits of the card
	Last4Digits string `json:"last_4_digits,omitempty"`
	// ExpMonth is the expiry month
	ExpMonth int `json:"exp_month,omitempty"`
	// ExpYear is the expiry year, in a 4 digits format
	ExpYear int `json:"exp_year,omitempty"`
	// Metadata is the metadata related to the card, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the card is in sandbox environment
	Sandbox bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the card was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// SetClient sets the client for the Card object and its
// children
func (s *Card) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
}

// dummyCard is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCard() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New(nil, "", "")
}
