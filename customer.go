package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// Customers manages the Customer struct
type Customers struct {
	p *ProcessOut
}

type Customer struct {
	// Address1 : Main address of the customer
	Address1 string `json:"address1"`
	// Address2 : Secondary address of the customer
	Address2 string `json:"address2"`
	// City : Shipping city of the customer
	City string `json:"city"`
	// CountryCode : Shipping country code of the customer
	CountryCode string `json:"country_code"`
	// FirstName : First name of the customer
	FirstName string `json:"first_name"`
	// ID : Id of the customer
	ID string `json:"id"`
	// LastName : Last name of the customer
	LastName string `json:"last_name"`
	// State : Shipping state of the customer
	State string `json:"state"`
	// Zip : Shipping ZIP code of the customer
	Zip string `json:"zip"`
}

// Tokens : Get all the authorization tokens of the customer.
func (c Customers) Tokens(customer *Customer) ([]*CustomerToken, error) {
	type Response struct {
		Tokens  []*CustomerToken `json:"tokens"`
		Success bool             `json:"success"`
		Message string           `json:"message"`
	}

	path := "customers/{id}/tokens"
	path = strings.Replace(path, "{id}", customer.ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return payload.Tokens, nil
}

// Authorize : Authorize (create) a new customer token.
func (c Customers) Authorize(customer *Customer, gatewayName, name, token string) (*CustomerToken, error) {

	type Request struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	}

	type Response struct {
		CustomerToken `json:"token"`
		Success       bool   `json:"success"`
		Message       string `json:"message"`
	}

	body, err := json.Marshal(&Request{
		Name:  name,
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	path := "/customers/{id}/gateways/{gateway_name}/tokens"
	path = strings.Replace(path, "{id}", customer.ID, -1)
	path = strings.Replace(path, "{gateway_name}", gatewayName, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", c.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return &payload.CustomerToken, nil
}

// Find : Get the customer data.
func (c Customers) Find(ID string) (*Customer, error) {
	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	path := "/customers/{id}"
	path = strings.Replace(path, "{id}", ID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return &payload.Customer, nil
}

// Save : Update the customer data.
func (c Customers) Save(customer *Customer) (*Customer, error) {
	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	body, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}

	path := "/customers/{id}"
	path = strings.Replace(path, "{id}", customer.ID, -1)

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", c.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return &payload.Customer, nil
}

// Delete : Delete the customer.
func (c Customers) Delete(customer *Customer) error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/customers/{id}"
	path = strings.Replace(path, "{id}", customer.ID, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", c.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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

// All : Get the customers list belonging to the project.
func (c Customers) All() ([]*Customer, error) {
	type Response struct {
		Customers []*Customer `json:"customers"`
		Success   bool        `json:"success"`
		Message   string      `json:"message"`
	}

	path := "/customers"

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return payload.Customers, nil
}

// Create : Create a new customer.
func (c Customers) Create(customer *Customer) (*Customer, error) {
	type Response struct {
		Customer `json:"customer"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
	}

	body, err := json.Marshal(customer)
	if err != nil {
		return nil, err
	}

	path := "/customers"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", c.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return &payload.Customer, nil
}

// FindToken : Find a specific customer token.
func (c Customers) FindToken(customer *Customer, tokenID string) (*CustomerToken, error) {
	type Response struct {
		CustomerToken `json:"token"`
		Success       bool   `json:"success"`
		Message       string `json:"message"`
	}

	path := "customers/{id}/tokens/{token_id}"
	path = strings.Replace(path, "{id}", customer.ID, -1)
	path = strings.Replace(path, "{token_id}", tokenID, -1)

	req, err := http.NewRequest(
		"GET",
		Host+path,
		nil,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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
	return &payload.CustomerToken, nil
}

// Revoke : Revoke (delete) a specific customer token.
func (c Customers) Revoke(customer *Customer, tokenID string) error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "customers/{id}/tokens/{token_id}"
	path = strings.Replace(path, "{id}", customer.ID, -1)
	path = strings.Replace(path, "{token_id}", tokenID, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", c.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.p.projectID, c.p.projectSecret)

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

// dummyCustomer is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyCustomer() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
