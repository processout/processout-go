package processout

import (
	"bytes"
	"encoding/json"
	"errors"
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
	ID string `json:"id"`
	// Project is the project to which the card belongs
	Project *Project `json:"project"`
	// Brand is the brand of the card (Visa, Mastercard, ...)
	Brand string `json:"brand"`
	// Type is the type of the card (Credit, Debit, ...)
	Type string `json:"type"`
	// BankName is the name of the bank of the card
	BankName string `json:"bank_name"`
	// Level is the level of the card (Electron, Classic, Gold, ...)
	Level string `json:"level"`
	// Iin is the first 6 digits of the card
	Iin string `json:"iin"`
	// Last4Digits is the last 4 digits of the card
	Last4Digits string `json:"last_4_digits"`
	// ExpMonth is the expiry month
	ExpMonth int `json:"exp_month"`
	// ExpYear is the expiry year, in a 4 digits format
	ExpYear int `json:"exp_year"`
	// Metadata is the metadata related to the card, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox is the define whether or not the card is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the card was created
	CreatedAt time.Time `json:"created_at"`
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
	errors.New("")
}
