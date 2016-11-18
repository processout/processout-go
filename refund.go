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

// Refund represents the Refund API object
type Refund struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the refund
	ID string `json:"id"`
	// Transaction is the transaction to which the refund is applied
	Transaction *Transaction `json:"transaction"`
	// Reason is the reason for the refund. Either customer_request, duplicate or fraud
	Reason string `json:"reason"`
	// Information is the custom details regarding the refund
	Information string `json:"information"`
	// Amount is the amount to be refunded. Must not be greater than the amount still available on the transaction
	Amount string `json:"amount"`
	// Metadata is the metadata related to the refund, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// Sandbox is the define whether or not the refund is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the refund was done
	CreatedAt *time.Time `json:"created_at"`
}

func (s *Refund) setClient(c *ProcessOut) {
	s.Client = c
	if s.Transaction != nil {
		s.Transaction.setClient(c)
	}
}

// Find allows you to find a transaction's refund by its ID.
func (s Refund) Find(transactionID, refundID string, options ...Options) (*Refund, error) {
	if s.Client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Refund  *Refund `json:"refund"`
		Success bool    `json:"success"`
		Message string  `json:"message"`
		Code    string  `json:"error_type"`
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

	path := "/transactions/" + url.QueryEscape(transactionID) + "/refunds/" + url.QueryEscape(refundID) + ""

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

	payload.Refund.setClient(s.Client)
	return payload.Refund, nil
}

// Apply allows you to apply a refund to a transaction.
func (s Refund) Apply(transactionID string, options ...Options) error {
	if s.Client == nil {
		panic("Please use the client.NewRefund() method to create a new Refund object")
	}

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
		Code    string `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"amount":      s.Amount,
		"metadata":    s.Metadata,
		"reason":      s.Reason,
		"information": s.Information,
		"expand":      opt.Expand,
		"filter":      opt.Filter,
		"limit":       opt.Limit,
		"page":        opt.Page,
		"end_before":  opt.EndBefore,
		"start_after": opt.StartAfter,
	})
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/transactions/" + url.QueryEscape(transactionID) + "/refunds"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.New(err, "", "")
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
		return errors.New(err, "", "")
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return errors.New(err, "", "")
	}

	if !payload.Success {
		erri := errors.NewFromResponse(res.StatusCode, payload.Code,
			payload.Message)

		return erri
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
	errors.New(nil, "", "")
}
