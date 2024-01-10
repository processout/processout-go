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

// Addon represents the Addon API object
type Addon struct {
	// ID is the iD of the addon
	ID *string `json:"id,omitempty"`
	// Project is the project to which the addon belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the addon belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Subscription is the subscription to which the addon belongs
	Subscription *Subscription `json:"subscription,omitempty"`
	// SubscriptionID is the iD of the subscription to which the addon belongs
	SubscriptionID *string `json:"subscription_id,omitempty"`
	// Plan is the plan used to create the addon, if any
	Plan *Plan `json:"plan,omitempty"`
	// PlanID is the iD of the plan used to create the addon, if any
	PlanID *string `json:"plan_id,omitempty"`
	// Type is the type of the addon. Can be either metered or recurring
	Type *string `json:"type,omitempty"`
	// Name is the name of the addon
	Name *string `json:"name,omitempty"`
	// Amount is the amount of the addon
	Amount *string `json:"amount,omitempty"`
	// Quantity is the quantity of the addon
	Quantity *int `json:"quantity,omitempty"`
	// Metadata is the metadata related to the addon, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Sandbox is the define whether or not the addon is in sandbox environment
	Sandbox *bool `json:"sandbox,omitempty"`
	// CreatedAt is the date at which the addon was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *Addon) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the Addon object and its
// children
func (s *Addon) SetClient(c *ProcessOut) *Addon {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Subscription != nil {
		s.Subscription.SetClient(c)
	}
	if s.Plan != nil {
		s.Plan.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *Addon) Prefill(c *Addon) *Addon {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Subscription = c.Subscription
	s.SubscriptionID = c.SubscriptionID
	s.Plan = c.Plan
	s.PlanID = c.PlanID
	s.Type = c.Type
	s.Name = c.Name
	s.Amount = c.Amount
	s.Quantity = c.Quantity
	s.Metadata = c.Metadata
	s.Sandbox = c.Sandbox
	s.CreatedAt = c.CreatedAt

	return s
}

// AddonFetchSubscriptionAddonsParameters is the structure representing the
// additional parameters used to call Addon.FetchSubscriptionAddons
type AddonFetchSubscriptionAddonsParameters struct {
	*Options
	*Addon
}

// FetchSubscriptionAddons allows you to get the addons applied to the subscription.
func (s Addon) FetchSubscriptionAddons(subscriptionID string, options ...AddonFetchSubscriptionAddonsParameters) (*Iterator, error) {
	if s.client == nil {
		panic("Please use the client.NewAddon() method to create a new Addon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AddonFetchSubscriptionAddonsParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Addon)

	type Response struct {
		Addons []*Addon `json:"addons"`

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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + "/addons"

	req, err := http.NewRequest(
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

	addonsList := []Identifiable{}
	for _, o := range payload.Addons {
		addonsList = append(addonsList, o.SetClient(s.client))
	}
	addonsIterator := &Iterator{
		pos:     -1,
		path:    path,
		data:    addonsList,
		options: opt.Options,
		decoder: func(b io.Reader, i interface{}) (bool, error) {
			r := struct {
				Data    json.RawMessage `json:"addons"`
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
	return addonsIterator, nil
}

// AddonCreateParameters is the structure representing the
// additional parameters used to call Addon.Create
type AddonCreateParameters struct {
	*Options
	*Addon
	Prorate       interface{} `json:"prorate"`
	ProrationDate interface{} `json:"proration_date"`
	Preview       interface{} `json:"preview"`
}

// Create allows you to create a new addon to the given subscription ID.
func (s Addon) Create(options ...AddonCreateParameters) (*Addon, error) {
	if s.client == nil {
		panic("Please use the client.NewAddon() method to create a new Addon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AddonCreateParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Addon)

	type Response struct {
		Addon   *Addon `json:"addon"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		PlanID        interface{} `json:"plan_id"`
		Type          interface{} `json:"type"`
		Name          interface{} `json:"name"`
		Amount        interface{} `json:"amount"`
		Quantity      interface{} `json:"quantity"`
		Metadata      interface{} `json:"metadata"`
		Prorate       interface{} `json:"prorate"`
		ProrationDate interface{} `json:"proration_date"`
		Preview       interface{} `json:"preview"`
	}{
		Options:       opt.Options,
		PlanID:        s.PlanID,
		Type:          s.Type,
		Name:          s.Name,
		Amount:        s.Amount,
		Quantity:      s.Quantity,
		Metadata:      s.Metadata,
		Prorate:       opt.Prorate,
		ProrationDate: opt.ProrationDate,
		Preview:       opt.Preview,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.SubscriptionID) + "/addons"

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

	payload.Addon.SetClient(s.client)
	return payload.Addon, nil
}

// AddonFindParameters is the structure representing the
// additional parameters used to call Addon.Find
type AddonFindParameters struct {
	*Options
	*Addon
}

// Find allows you to find a subscription's addon by its ID.
func (s Addon) Find(subscriptionID, addonID string, options ...AddonFindParameters) (*Addon, error) {
	if s.client == nil {
		panic("Please use the client.NewAddon() method to create a new Addon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AddonFindParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Addon)

	type Response struct {
		Addon   *Addon `json:"addon"`
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

	path := "/subscriptions/" + url.QueryEscape(subscriptionID) + "/addons/" + url.QueryEscape(addonID) + ""

	req, err := http.NewRequest(
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

	payload.Addon.SetClient(s.client)
	return payload.Addon, nil
}

// AddonSaveParameters is the structure representing the
// additional parameters used to call Addon.Save
type AddonSaveParameters struct {
	*Options
	*Addon
	Prorate             interface{} `json:"prorate"`
	ProrationDate       interface{} `json:"proration_date"`
	Preview             interface{} `json:"preview"`
	IncrementQuantityBy interface{} `json:"increment_quantity_by"`
}

// Save allows you to save the updated addon attributes.
func (s Addon) Save(options ...AddonSaveParameters) (*Addon, error) {
	if s.client == nil {
		panic("Please use the client.NewAddon() method to create a new Addon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AddonSaveParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Addon)

	type Response struct {
		Addon   *Addon `json:"addon"`
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		PlanID              interface{} `json:"plan_id"`
		Type                interface{} `json:"type"`
		Name                interface{} `json:"name"`
		Amount              interface{} `json:"amount"`
		Quantity            interface{} `json:"quantity"`
		Metadata            interface{} `json:"metadata"`
		Prorate             interface{} `json:"prorate"`
		ProrationDate       interface{} `json:"proration_date"`
		Preview             interface{} `json:"preview"`
		IncrementQuantityBy interface{} `json:"increment_quantity_by"`
	}{
		Options:             opt.Options,
		PlanID:              s.PlanID,
		Type:                s.Type,
		Name:                s.Name,
		Amount:              s.Amount,
		Quantity:            s.Quantity,
		Metadata:            s.Metadata,
		Prorate:             opt.Prorate,
		ProrationDate:       opt.ProrationDate,
		Preview:             opt.Preview,
		IncrementQuantityBy: opt.IncrementQuantityBy,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.SubscriptionID) + "/addons/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
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

	payload.Addon.SetClient(s.client)
	return payload.Addon, nil
}

// AddonDeleteParameters is the structure representing the
// additional parameters used to call Addon.Delete
type AddonDeleteParameters struct {
	*Options
	*Addon
	Prorate       interface{} `json:"prorate"`
	ProrationDate interface{} `json:"proration_date"`
	Preview       interface{} `json:"preview"`
}

// Delete allows you to delete an addon applied to a subscription.
func (s Addon) Delete(options ...AddonDeleteParameters) error {
	if s.client == nil {
		panic("Please use the client.NewAddon() method to create a new Addon object")
	}
	if len(options) > 1 {
		panic("The options parameter should only be provided once.")
	}

	opt := AddonDeleteParameters{}
	if len(options) == 1 {
		opt = options[0]
	}
	if opt.Options == nil {
		opt.Options = &Options{}
	}
	s.Prefill(opt.Addon)

	type Response struct {
		HasMore bool   `json:"has_more"`
		Success bool   `json:"success"`
		Message string `json:"message"`
		Code    string `json:"error_type"`
	}

	data := struct {
		*Options
		Prorate       interface{} `json:"prorate"`
		ProrationDate interface{} `json:"proration_date"`
		Preview       interface{} `json:"preview"`
	}{
		Options:       opt.Options,
		Prorate:       opt.Prorate,
		ProrationDate: opt.ProrationDate,
		Preview:       opt.Preview,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return errors.New(err, "", "")
	}

	path := "/subscriptions/" + url.QueryEscape(*s.SubscriptionID) + "/addons/" + url.QueryEscape(*s.ID) + ""

	req, err := http.NewRequest(
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

// dummyAddon is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyAddon() {
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
