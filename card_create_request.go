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

// CardCreateRequest represents the CardCreateRequest API object
type CardCreateRequest struct {
	// Device is the device used to create the card
	Device *Device `json:"device,omitempty"`
	// Name is the cardholder name
	Name *string `json:"name,omitempty"`
	// Number is the card PAN (raw)
	Number *string `json:"number,omitempty"`
	// ExpDay is the card expiration day. Used for Apple Pay
	ExpDay *string `json:"exp_day,omitempty"`
	// ExpMonth is the card expiration month
	ExpMonth *string `json:"exp_month,omitempty"`
	// ExpYear is the card expiration year
	ExpYear *string `json:"exp_year,omitempty"`
	// Cvc2 is the cVC2
	Cvc2 *string `json:"cvc2,omitempty"`
	// PreferredScheme is the preferred card scheme
	PreferredScheme *string `json:"preferred_scheme,omitempty"`
	// Metadata is the metadata related to the card, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// TokenType is the this field defines if the card was tokenized with a 3rd party tokenization method: applepay, googlepay
	TokenType *string `json:"token_type,omitempty"`
	// Eci is the eCI indicator. Used if the card was tokenized with a 3rd party tokenization method
	Eci *string `json:"eci,omitempty"`
	// Cryptogram is the cryptogram (Base64-encoded). Used if the card was tokenized with a 3rd party tokenization method
	Cryptogram *string `json:"cryptogram,omitempty"`
	// ApplepayResponse is the raw ApplePay card tokenization response. Used if the card was tokenized with a 3rd party tokenization method
	ApplepayResponse *string `json:"applepay_response,omitempty"`
	// ApplepayMid is the applePay MID. Used if the card was tokenized with a 3rd party tokenization method
	ApplepayMid *string `json:"applepay_mid,omitempty"`
	// PaymentToken is the google payment token. Used if the card was tokenized with a 3rd party tokenization method
	PaymentToken *string `json:"payment_token,omitempty"`
	// Contact is the cardholder contact information
	Contact *CardContact `json:"contact,omitempty"`
	// Shipping is the cardholder shipping information
	Shipping *CardShipping `json:"shipping,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the CardCreateRequest object and its
// children
func (s *CardCreateRequest) SetClient(c *ProcessOut) *CardCreateRequest {
	if s == nil {
		return s
	}
	s.client = c
	if s.Device != nil {
		s.Device.SetClient(c)
	}
	if s.Contact != nil {
		s.Contact.SetClient(c)
	}
	if s.Shipping != nil {
		s.Shipping.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *CardCreateRequest) Prefill(c *CardCreateRequest) *CardCreateRequest {
	if c == nil {
		return s
	}

	s.Device = c.Device
	s.Name = c.Name
	s.Number = c.Number
	s.ExpDay = c.ExpDay
	s.ExpMonth = c.ExpMonth
	s.ExpYear = c.ExpYear
	s.Cvc2 = c.Cvc2
	s.PreferredScheme = c.PreferredScheme
	s.Metadata = c.Metadata
	s.TokenType = c.TokenType
	s.Eci = c.Eci
	s.Cryptogram = c.Cryptogram
	s.ApplepayResponse = c.ApplepayResponse
	s.ApplepayMid = c.ApplepayMid
	s.PaymentToken = c.PaymentToken
	s.Contact = c.Contact
	s.Shipping = c.Shipping

	return s
}

// CardCreateRequestCreateParameters is the structure representing the
// additional parameters used to call CardCreateRequest.Create
type CardCreateRequestCreateParameters struct {
	*Options
	*CardCreateRequest
}

// Create allows you to create a new card.
func (s CardCreateRequest) Create(options ...CardCreateRequestCreateParameters) (*CardCreateRequest, error) {
	if s.client == nil {
		panic("Please use the client.NewCardCreateRequest() method to create a new CardCreateRequest object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := CardCreateRequestCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.CardCreateRequest)

	type Response struct {
		CardCreateRequest *CardCreateRequest `json:"card"`
		HasMore           bool               `json:"has_more"`
		Success           bool               `json:"success"`
		Message           string             `json:"message"`
		Code              string             `json:"error_type"`
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
		"POST",
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

	payload.CardCreateRequest.SetClient(s.client)
	return payload.CardCreateRequest, nil
}

// dummyCardCreateRequest is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCardCreateRequest() {
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
