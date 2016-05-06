package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// PaymentGateways manages the PaymentGateway struct
type PaymentGateways struct {
	p *ProcessOut
}

type PaymentGateway struct {
	// Beta : Determine if the gateway's integration is still in beta
	Beta bool `json:"beta"`
	// DisplayName : Name of the payment gateway to be displayed
	DisplayName string `json:"display_name"`
	// Name : Internal name of the payment gateway
	Name string `json:"name"`
	// PublicKeys :
	PublicKeys []*PaymentGatewayPublicKey `json:"public_keys"`
	// Settings : Settings of the payment gateway, in the form of a dictionary
	Settings map[string]string `json:"settings"`
}

// Save : Update or set the payment gateway settings.
func (p PaymentGateways) Save(paymentGateway *PaymentGateway, gatewayName string) (*PaymentGateway, error) {
	type Response struct {
		PaymentGateway `json:"gateway"`
		Success        bool   `json:"success"`
		Message        string `json:"message"`
	}

	body, err := json.Marshal(paymentGateway)
	if err != nil {
		return nil, err
	}

	path := "/gateways/{gateway_name}"
	path = strings.Replace(path, "{gateway_name}", gatewayName, -1)

	req, err := http.NewRequest(
		"PUT",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", p.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(p.p.projectID, p.p.projectSecret)

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
	return &payload.PaymentGateway, nil
}

// Delete : Remove the payment gateway and its settings from the current project.
func (p PaymentGateways) Delete(gatewayName string) error {
	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	path := "/gateways/{gateway_name}"
	path = strings.Replace(path, "{gateway_name}", gatewayName, -1)

	req, err := http.NewRequest(
		"DELETE",
		Host+path,
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("API-Version", p.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(p.p.projectID, p.p.projectSecret)

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

// dummyPaymentGateway is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPaymentGateway() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
	}
	errors.New("")
}
