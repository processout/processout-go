package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Cards manages the Card struct
type Cards struct {
	p *ProcessOut
}

type Card struct {
	// ID : ID of the card
	ID string `json:"id"`
	// Project : Project to which the card belongs
	Project *Project `json:"project"`
	// Brand : Brand of the card (Visa, Mastercard, ...)
	Brand string `json:"brand"`
	// Type : Type of the card (Credit, Debit, ...)
	Type string `json:"type"`
	// BankName : Name of the bank of the card
	BankName string `json:"bank_name"`
	// Level : Level of the card (Electron, Classic, Gold, ...)
	Level string `json:"level"`
	// Iin : First 6 digits of the card
	Iin string `json:"iin"`
	// Last4Digits : Last 4 digits of the card
	Last4Digits string `json:"last_4_digits"`
	// ExpMonth : Expiry month
	ExpMonth int `json:"exp_month"`
	// ExpYear : Expiry year, in a 4 digits format
	ExpYear int `json:"exp_year"`
	// Metadata : Metadata related to the card, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox : Define whether or not the card is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the card was created
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
