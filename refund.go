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

// Refunds manages the Refund struct
type Refunds struct {
	p *ProcessOut
}

type Refund struct {
	// ID : ID of the refund
	ID string `json:"id"`
	// Transaction : Transaction to which the refund is applied
	Transaction *Transaction `json:"transaction"`
	// Reason : Reason for the refund. Either customer_request, duplicate or fraud
	Reason string `json:"reason"`
	// Information : Custom details regarding the refund
	Information string `json:"information"`
	// Amount : Amount to be refunded. Must not be greater than the amount still available on the transaction
	Amount string `json:"amount"`
	// Metadata : Metadata related to the refund, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox : Define whether or not the refund is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the refund was done
	CreatedAt time.Time `json:"created_at"`
}

// Find : Find a transaction's refund by its ID.
func (s Refunds) Find(transactionID, refundID string, options ...Options) (*Refund, error) {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Refund  `json:"refund"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"expand": opt.Expand,
	})
	if err != nil {
		return nil, err
	}

	path := "/transactions/" + url.QueryEscape(transactionID) + "/refunds/" + url.QueryEscape(refundID) + ""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Refund, nil
}

// Apply : Apply a refund to a transaction.
func (s Refunds) Apply(refund *Refund, transactionID string, options ...Options) error {
	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"amount":      refund.Amount,
		"metadata":    refund.Metadata,
		"reason":      refund.Reason,
		"information": refund.Information,
		"expand":      opt.Expand,
	})
	if err != nil {
		return err
	}

	path := "/transactions/{transactions_id}/refunds"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Accept", "application/json")
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	req.SetBasicAuth(s.p.projectID, s.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return err
	}

	if !payload.Success {
		return errors.New(payload.Message)
	}
	return nil
}

// dummyRefund is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyRefund() {
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
