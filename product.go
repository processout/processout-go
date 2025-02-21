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

// Product represents the Product API object
type Product struct {
	// ID is the iD of the product
	ID *string `json:"id,omitempty"`
	// Project is the project to which the product belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the product belongs
	ProjectID *string `json:"project_id,omitempty"`
	// URL is the uRL to which you may redirect your customer to proceed with the payment
	URL *string `json:"url,omitempty"`
	// Name is the name of the product
	Name *string `json:"name,omitempty"`
	// Amount is the amount of the product
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency of the product
	Currency *string `json:"currency,omitempty"`
	// Metadata is the metadata related to the product, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// ReturnURL is the uRL where the customer will be redirected upon payment
	ReturnURL *string `json:"return_url,omitempty"`
	// CancelURL is the uRL where the customer will be redirected if the paymen was canceled
	CancelURL *string `json:"cancel_url,omitempty"`
	// Sandbox is the define whether or not the product is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the product was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Product) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Product object and its
// children
func (s *Product) SetClient(c *ProcessOut) *Product {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Product) Prefill(c *Product) *Product {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.URL = c.URL
	s.Name = c.Name
	s.Amount = c.Amount
	s.Currency = c.Currency
	s.Metadata = c.Metadata
	s.ReturnURL = c.ReturnURL
	s.CancelURL = c.CancelURL
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// ProductCreateInvoiceParameters is the structure representing the
// additional parameters used to call Product.CreateInvoice
type ProductCreateInvoiceParameters struct {
	*Options
	*Product
}

// CreateInvoice allows you to create a new invoice from the product.
func (s Product) CreateInvoice(options ...ProductCreateInvoiceParameters) (*Invoice, error) {
	return s.CreateInvoiceWithContext(context.Background(), options...)
}

// CreateInvoice allows you to create a new invoice from the product., passes the provided context to the request
func (s Product) CreateInvoiceWithContext(ctx context.Context, options ...ProductCreateInvoiceParameters) (*Invoice, error) {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductCreateInvoiceParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
		Invoice *Invoice `json:"invoice"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/products/" + url.QueryEscape(*s.ID) + "/invoices"

	req, err := http.NewRequestWithContext(
		ctx,
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

	payload.Invoice.SetClient(s.client)
	return payload.Invoice, nil
}

// ProductAllParameters is the structure representing the
// additional parameters used to call Product.All
type ProductAllParameters struct {
	*Options
	*Product
}

// All allows you to get all the products.
func (s Product) All(options ...ProductAllParameters) (*Iterator, error) {
	return s.AllWithContext(context.Background(), options...)
}

// All allows you to get all the products., passes the provided context to the request
func (s Product) AllWithContext(ctx context.Context, options ...ProductAllParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductAllParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
		Products []*Product `json:"products"`

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

	path := "/products"

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
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

	productsList := []Identifiable{}
	for _, o := range payload.Products {
		productsList = append(productsList, o.SetClient(s.client))
	}
	productsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    productsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"products"`
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
	return productsIterator, nil
}

// ProductCreateParameters is the structure representing the
// additional parameters used to call Product.Create
type ProductCreateParameters struct {
	*Options
	*Product
}

// Create allows you to create a new product.
func (s Product) Create(options ...ProductCreateParameters) (*Product, error) {
	return s.CreateWithContext(context.Background(), options...)
}

// Create allows you to create a new product., passes the provided context to the request
func (s Product) CreateWithContext(ctx context.Context, options ...ProductCreateParameters) (*Product, error) {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
		Product *Product `json:"product"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	data := struct {
		*Options
		Name      interface{} `json:"name"`
		Amount    interface{} `json:"amount"`
		Currency  interface{} `json:"currency"`
		Metadata  interface{} `json:"metadata"`
		ReturnURL interface{} `json:"return_url"`
		CancelURL interface{} `json:"cancel_url"`
	}{
		Options:   opt.Options,
		Name:      s.Name,
		Amount:    s.Amount,
		Currency:  s.Currency,
		Metadata:  s.Metadata,
		ReturnURL: s.ReturnURL,
		CancelURL: s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/products"

	req, err := http.NewRequestWithContext(
		ctx,
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

	payload.Product.SetClient(s.client)
	return payload.Product, nil
}

// ProductFindParameters is the structure representing the
// additional parameters used to call Product.Find
type ProductFindParameters struct {
	*Options
	*Product
}

// Find allows you to find a product by its ID.
func (s Product) Find(productID string, options ...ProductFindParameters) (*Product, error) {
	return s.FindWithContext(context.Background(), productID, options...)
}

// Find allows you to find a product by its ID., passes the provided context to the request
func (s Product) FindWithContext(ctx context.Context, productID string, options ...ProductFindParameters) (*Product, error) {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
		Product *Product `json:"product"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
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

	path := "/products/" + url.QueryEscape(productID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
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

	payload.Product.SetClient(s.client)
	return payload.Product, nil
}

// ProductSaveParameters is the structure representing the
// additional parameters used to call Product.Save
type ProductSaveParameters struct {
	*Options
	*Product
}

// Save allows you to save the updated product attributes.
func (s Product) Save(options ...ProductSaveParameters) (*Product, error) {
	return s.SaveWithContext(context.Background(), options...)
}

// Save allows you to save the updated product attributes., passes the provided context to the request
func (s Product) SaveWithContext(ctx context.Context, options ...ProductSaveParameters) (*Product, error) {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
		Product *Product `json:"product"`
		HasMore bool     `json:"has_more"`
		Success bool     `json:"success"`
		Message string   `json:"message"`
		Code    string   `json:"error_type"`
	}

	data := struct {
		*Options
		Name      interface{} `json:"name"`
		Amount    interface{} `json:"amount"`
		Currency  interface{} `json:"currency"`
		Metadata  interface{} `json:"metadata"`
		ReturnURL interface{} `json:"return_url"`
		CancelURL interface{} `json:"cancel_url"`
	}{
		Options:   opt.Options,
		Name:      s.Name,
		Amount:    s.Amount,
		Currency:  s.Currency,
		Metadata:  s.Metadata,
		ReturnURL: s.ReturnURL,
		CancelURL: s.CancelURL,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/products/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"PUT",
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

	payload.Product.SetClient(s.client)
	return payload.Product, nil
}

// ProductDeleteParameters is the structure representing the
// additional parameters used to call Product.Delete
type ProductDeleteParameters struct {
	*Options
	*Product
}

// Delete allows you to delete the product.
func (s Product) Delete(options ...ProductDeleteParameters) error {
	return s.DeleteWithContext(context.Background(), options...)
}

// Delete allows you to delete the product., passes the provided context to the request
func (s Product) DeleteWithContext(ctx context.Context, options ...ProductDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewProduct() method to create a new Product object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := ProductDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Product)

	type Response struct {
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
		return errors.New(err, "", "")
	}

	path := "/products/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequestWithContext(
		ctx,
		"DELETE",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	setupRequest(s.client, opt.Options, req)

	res, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return errors.NewNetworkError(err)
	}
	payload := &Response{}
	defer res.Body.Close()
	if res.StatusCode >= 500 {
		return errors.New(nil, "", "An unexpected error occurred while processing your request.. A lot of sweat is already flowing from our developers head!")
	}
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
		g io.Reader
	}
	errors.New(nil, "", "")
}
