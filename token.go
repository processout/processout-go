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

// Tokens manages the Token struct
type Tokens struct {
	p *ProcessOut
}

type Token struct {
	// ID : ID of the customer token
	ID string `json:"id"`
	// Customer : Customer linked to the token
	Customer *Customer `json:"customer"`
	// CustomerID : ID of the customer token
	CustomerID string `json:"customer_id"`
	// Name : Name of the customer token
	Name string `json:"name"`
	// IsRecurringInvoice : Define whether or not the customer token is used on a recurring invoice
	IsRecurringInvoice string `json:"is_recurring_invoice"`
	// CreatedAt : Date at which the customer token was created
	CreatedAt time.Time `json:"created_at"`
}

// Delete : Delete a specific customer's token by its ID.
func (s Tokens) Delete(token *Token) (*Token, error) {

	type Response struct {
		Token   `json:"token"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	_, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	path := "/customers/" + url.QueryEscape(token.CustomerID) + "/tokens/" + url.QueryEscape(token.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", s.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
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
	return &payload.Token, nil
}

// dummyToken is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyToken() {
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
