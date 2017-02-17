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

// CardInformation represents the CardInformation API object
type CardInformation struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// Iin is the first 6 digits of the card
	Iin string `json:"iin,omitempty"`
	// Scheme is the scheme of the card, such as visa or mastercard
	Scheme string `json:"scheme,omitempty"`
	// Type is the type of the card (Credit, Debit, ...)
	Type string `json:"type,omitempty"`
	// BankName is the name of the bank of the card
	BankName string `json:"bank_name,omitempty"`
	// Brand is the level of the card (Electron, Classic, Gold, ...)
	Brand string `json:"brand,omitempty"`
	// Country is the country that issued the card
	Country string `json:"country,omitempty"`
}

// SetClient sets the client for the CardInformation object and its
// children
func (s *CardInformation) SetClient(c *ProcessOut) {
	if s == nil {
		return
	}
	s.Client = c
}

// Fetch allows you to fetch card information from the IIN.
func (s CardInformation) Fetch(iin string, options ...Options) (*CardInformation, error) {
	if s.Client == nil {
		panic("Please use the client.NewCardInformation() method to create a new CardInformation object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		CardInformation *CardInformation `json:"coupon"`
		Success         bool             `json:"success"`
		Message         string           `json:"message"`
		Code            string           `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
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
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.Client.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(s.Client.projectID, s.Client.projectSecret)

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

	payload.CardInformation.SetClient(s.Client)
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
	}
	errors.New(nil, "", "")
}
