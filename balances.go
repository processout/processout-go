package processout

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
)

// Balances represents the Balances API object
type Balances struct {
	// Vouchers is the vouchers linked to the customer
	Vouchers         *[]*Balance `json:"vouchers,omitempty"`
	AvailableBalance *Balance    `json:"available_balance,omitempty"`
	NextAction       *NextAction `json:"next_action,omitempty"`

	client *ProcessOut
}

type NextAction struct {
	ActionType string `json:"type"`
	Value      string `json:"value"`
}

// SetClient sets the client for the Balances object and its
// children
func (s *Balances) SetClient(c *ProcessOut) *Balances {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Balances) Prefill(c *Balances) *Balances {
	if c == nil {
		return s
	}

	s.Vouchers = c.Vouchers

	return s
}

// BalancesFindParameters is the structure representing the
// additional parameters used to call Balances.Find
type BalancesFindParameters struct {
	*Options
	*Balances

	// Parameter used to indicate that the API call retrieving the token balance is authenticated.
	Authenticated bool
}

// Find allows you to fetch a customer token's balance
func (s Balances) Find(tokenID string, options ...BalancesFindParameters) (*Balances, error) {
	return s.FindWithContext(context.Background(), tokenID, options...)
}

// Find allows you to fetch a customer token's balance, passes the provided context to the request
func (s Balances) FindWithContext(ctx context.Context, tokenID string, options ...BalancesFindParameters) (*Balances, error) {
	if s.client == nil {
		panic("Please use the client.NewBalances() method to create a new Balances object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := BalancesFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}

	s.Prefill(opt.Balances)

	type Response struct {
		Balances *Balances `json:"balances"`
		HasMore  bool      `json:"has_more"`
		Success  bool      `json:"success"`
		Message  string    `json:"message"`
		Code     string    `json:"error_type"`
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

	path := "/balances/tokens/" + url.QueryEscape(tokenID) + ""
	u, err := url.Parse(Host + path)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	q := u.Query()
	if opt.Authenticated {
		q.Add("authenticated", "true")
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		u.String(),
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

	payload.Balances.SetClient(s.client)
	return payload.Balances, nil
}

// dummyBalances is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyBalances() {
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
