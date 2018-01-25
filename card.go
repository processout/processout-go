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

// Card represents the Card API object
type Card struct {
	// ID is the iD of the card
	ID *string `json:"id,omitempty"`
	// Project is the project to which the card belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the card belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Token is the token linked to the card, which can be used to process payments
	Token *Token `json:"token,omitempty"`
	// Scheme is the scheme of the card, such as visa or mastercard
	Scheme *string `json:"scheme,omitempty"`
	// Type is the type of the card (Credit, Debit, ...)
	Type *string `json:"type,omitempty"`
	// BankName is the name of the bank of the card
	BankName *string `json:"bank_name,omitempty"`
	// Brand is the level of the card (Electron, Classic, Gold, ...)
	Brand *string `json:"brand,omitempty"`
	// Iin is the first 6 digits of the card
	Iin *string `json:"iin,omitempty"`
	// Last4Digits is the last 4 digits of the card
	Last4Digits *string `json:"last_4_digits,omitempty"`
	// ExpMonth is the expiry month
	ExpMonth *int `json:"exp_month,omitempty"`
	// ExpYear is the expiry year, in a 4 digits format
	ExpYear *int `json:"exp_year,omitempty"`
	// CvcCheck is the status of the CVC check initially made on the card when the CVC was provided
	CvcCheck *string `json:"cvc_check,omitempty"`
	// AvsCheck is the status of the AVS check initially made on the card when the AVS was provided
	AvsCheck *string `json:"avs_check,omitempty"`
	// Name is the name of the card holder
	Name *string `json:"name,omitempty"`
	// Address1 is the address line of the card holder
	Address1 *string `json:"address1,omitempty"`
	// Address2 is the secondary address line of the card holder
	Address2 *string `json:"address2,omitempty"`
	// City is the city of the card holder
	City *string `json:"city,omitempty"`
	// State is the state of the card holder
	State *string `json:"state,omitempty"`
	// CountryCode is the country code of the card holder (ISO-3166, 2 characters format)
	CountryCode *string `json:"country_code,omitempty"`
	// Zip is the zIP code of the card holder
	Zip *string `json:"zip,omitempty"`
	// Metadata is the metadata related to the card, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ExpiresSoon is the contains true if the card will expire soon, false otherwise
	ExpiresSoon *bool `json:"expires_soon,omitempty"`
	// Sandbox is the define whether or not the card is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the card was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Card) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Card object and its
// children
func (s *Card) SetClient(c *ProcessOut) *Card {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Token != nil {
		s.Token.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Card) Prefill(c *Card) *Card {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Token = c.Token
	s.Scheme = c.Scheme
	s.Type = c.Type
	s.BankName = c.BankName
	s.Brand = c.Brand
	s.Iin = c.Iin
	s.Last4Digits = c.Last4Digits
	s.ExpMonth = c.ExpMonth
	s.ExpYear = c.ExpYear
	s.CvcCheck = c.CvcCheck
	s.AvsCheck = c.AvsCheck
	s.Name = c.Name
	s.Address1 = c.Address1
	s.Address2 = c.Address2
	s.City = c.City
	s.State = c.State
	s.CountryCode = c.CountryCode
	s.Zip = c.Zip
	s.Metadata = c.Metadata
	s.ExpiresSoon = c.ExpiresSoon
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// CardAllParameters is the structure representing the
// additional parameters used to call Card.All
type CardAllParameters struct {
	*Options
	*Card
}

// All allows you to get all the cards.
func (s Card) All(options ...CardAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewCard() method to create a new Card object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CardAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Card)

	type Response struct {
		Cards []*Card `json:"cards"`

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

	path := "/cards"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	cardsList := []Identifiable{}
	for _, o := range payload.Cards {
		cardsList = append(cardsList, o.SetClient(s.client))
	}
	cardsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    cardsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"cards"`
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
		hasMorePrev: false,
	}
	return cardsIterator, nil
}

// CardFindParameters is the structure representing the
// additional parameters used to call Card.Find
type CardFindParameters struct {
	*Options
	*Card
}

// Find allows you to find a card by its ID.
func (s Card) Find(cardID string, options ...CardFindParameters) (*Card, error) {
	if s.client == nil {
		panic("Please use the client.NewCard() method to create a new Card object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CardFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Card)

	type Response struct {
		Card    *Card  `json:"card"`
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

	path := "/cards/" + url.QueryEscape(cardID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, errors.New(err, "", "")
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
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

	payload.Card.SetClient(s.client)
	return payload.Card, nil
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
