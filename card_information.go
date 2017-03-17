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

// CardInformation represents the CardInformation API object
type CardInformation struct {
	// Iin is the first 6 digits of the card
	Iin string `json:"iin,omitempty"`
	// Scheme is the scheme of the card, such as visa or mastercard
	Scheme string `json:"scheme,omitempty"`
	// Type is the type of the card (Credit, Debit, ...)
	Type *string `json:"type,omitempty"`
	// BankName is the name of the bank of the card
	BankName *string `json:"bank_name,omitempty"`
	// Brand is the level of the card (Electron, Classic, Gold, ...)
	Brand *string `json:"brand,omitempty"`
	// Country is the country that issued the card
	Country *string `json:"country,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardInformation object and its
// children
func (s *CardInformation) SetClient(c *ProcessOut) *CardInformation {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardInformation) Prefill(c *CardInformation) *CardInformation {
	if c == nil {
		return s
	}

	s.Iin = c.Iin
	s.Scheme = c.Scheme
	s.Type = c.Type
	s.BankName = c.BankName
	s.Brand = c.Brand
	s.Country = c.Country

	return s
}

// CardInformationFetchParameters is the structure representing the
// additional parameters used to call CardInformation.Fetch
type CardInformationFetchParameters struct {
	*Options
	*CardInformation
}

// Fetch allows you to fetch card information from the IIN.
func (s CardInformation) Fetch(iin string, options ...CardInformationFetchParameters) (*CardInformation, error) {
	if s.client == nil {
		panic("Please use the client.NewCardInformation() method to create a new CardInformation object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CardInformationFetchParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.CardInformation)

	type Response struct {
		CardInformation *CardInformation `json:"card_information"`
		HasMore         bool             `json:"has_more"`
		Success         bool             `json:"success"`
		Message         string           `json:"message"`
		Code            string           `json:"error_type"`
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

	path := "/iins/" + url.QueryEscape(iin) + ""

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

	payload.CardInformation.SetClient(s.client)
	return payload.CardInformation, nil
}

// dummyCardInformation is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardInformation() {
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
