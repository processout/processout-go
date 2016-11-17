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

// Product represents the Product API object
type Product struct {
	// Client is the ProcessOut client used to communicate with the API
	Client *ProcessOut
	// ID is the iD of the product
	ID string `json:"id"`
	// Project is the project to which the product belongs
	Project *Project `json:"project"`
	// URL is the uRL to which you may redirect your customer to proceed with the payment
	URL string `json:"url"`
	// Name is the name of the product
	Name string `json:"name"`
	// Amount is the amount of the product
	Amount string `json:"amount"`
	// Currency is the currency of the product
	Currency string `json:"currency"`
	// Metadata is the metadata related to the product, in the form of a dictionary (key-value pair)
	Metadata map[string]string `json:"metadata"`
	// RequestEmail is the choose whether or not to request the email during the checkout process
	RequestEmail bool `json:"request_email"`
	// RequestShipping is the choose whether or not to request the shipping address during the checkout process
	RequestShipping bool `json:"request_shipping"`
	// ReturnURL is the uRL where the customer will be redirected upon payment
	ReturnURL string `json:"return_url"`
	// CancelURL is the uRL where the customer will be redirected if the paymen was canceled
	CancelURL string `json:"cancel_url"`
	// Sandbox is the define whether or not the product is in sandbox environment
	Sandbox bool `json:"sandbox"`
	// CreatedAt is the date at which the product was created
	CreatedAt time.Time `json:"created_at"`
}

// CreateInvoice allows you to create a new invoice from the product.
func (s Product) CreateInvoice(options ...Options) (*Invoice, error) {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Invoice *Invoice `json:"invoice"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/products/" + url.QueryEscape(s.ID) + "/invoices"

	req, err := http.NewRequest(
		"POST",
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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Invoice, nil
}

// All allows you to get all the products.
func (s Product) All(options ...Options) ([]*Product, error) {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Products []*Product `json:"products"`

		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
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

	path := "/products"

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Products, nil
}

// Create allows you to create a new product.
func (s Product) Create(options ...Options) (*Product, error) {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Product *Product `json:"product"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":             s.Name,
		"amount":           s.Amount,
		"currency":         s.Currency,
		"metadata":         s.Metadata,
		"request_email":    s.RequestEmail,
		"request_shipping": s.RequestShipping,
		"return_url":       s.ReturnURL,
		"cancel_url":       s.CancelURL,
		"expand":           opt.Expand,
		"filter":           opt.Filter,
		"limit":            opt.Limit,
		"page":             opt.Page,
		"end_before":       opt.EndBefore,
		"start_after":      opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/products"

	req, err := http.NewRequest(
		"POST",
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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Product, nil
}

// Find allows you to find a product by its ID.
func (s Product) Find(productID string, options ...Options) (*Product, error) {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Product *Product `json:"product"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/products/" + url.QueryEscape(productID) + ""

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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Product, nil
}

// Save allows you to save the updated product attributes.
func (s Product) Save(options ...Options) (*Product, error) {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}

	opt := Options{}
	if len(options) == 1 {
		opt = options[0]
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	type Response struct {
		Product *Product `json:"product"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	body, err := json.Marshal(map[string]interface{}{
		"name":             s.Name,
		"amount":           s.Amount,
		"currency":         s.Currency,
		"metadata":         s.Metadata,
		"request_email":    s.RequestEmail,
		"request_shipping": s.RequestShipping,
		"return_url":       s.ReturnURL,
		"cancel_url":       s.CancelURL,
		"expand":           opt.Expand,
		"filter":           opt.Filter,
		"limit":            opt.Limit,
		"page":             opt.Page,
		"end_before":       opt.EndBefore,
		"start_after":      opt.StartAfter,
	})
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/products/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"PUT",
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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return nil, erri
	}

	return payload.Product, nil
}

// Delete allows you to delete the product.
func (s Product) Delete(options ...Options) error {
	if s.Client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
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

	path := "/products/" + url.QueryEscape(s.ID) + ""

	req, err := http.NewRequest(
		"DELETE",
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
		erri := errors.NewFromResponse(res.StatusCode, payload.Message,
			payload.Code)

		return erri
	}

	return nil
}

// dummyProduct is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyProduct() {
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
