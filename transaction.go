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

// Transactions manages the Transaction struct
type Transactions struct {
	p *ProcessOut
}

type Transaction struct {
	// ID : ID of the transaction
	ID string `json:"id"`
	// Status : Status of the transaction
	Status string `json:"status"`
	// Fee : ProcessOut fee applied on the transaction
	Fee string `json:"fee"`
	// Sandbox : Define whether or not the transaction is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt : Date at which the transaction was created
	CreatedAt time.Time `json:"created_at"`
}


// All : Get all the transactions.
func (s Transactions) All() ([]*Transaction, error) {

	type Response struct {
		Transactions []*Transaction `json:"transactions"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/transactions"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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
	return payload.Transactions, nil
}

// Find : Find a transaction by its ID.
func (s Transactions) Find(transactionID string) (*Transaction, error) {

	type Response struct {
		Transaction `json:"transaction"`
		Success bool `json:"success"`
		Message string `json:"message"`
	}

	 _ , err := json.Marshal(map[string]interface{}{

	})
	if err != nil {
		return nil, err
	}

	path := "/transactions/"+url.QueryEscape(transactionID)+""

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
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
	return &payload.Transaction, nil
}


// dummyTransaction is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyTransaction() {
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
