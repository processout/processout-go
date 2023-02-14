package processout

import (
	"fmt"
	"net/http"
	"time"
)

var (
	// RequestAPIVersion is the default version of the API used in requests
	// made with this package
	RequestAPIVersion = "1.4.0.0"
	// Host is the URL where API requests are made
	Host = "https://api.processout.com"

	// DefaultClient sets the HTTP default client used for ProcessOut clients
	DefaultClient = &http.Client{
		Timeout: time.Second * 95,
	}
)

// ProcessOut wraps all the components of the package in a
// single structure
type ProcessOut struct {
	// APIVersion is the version of the API to use
	APIVersion string
	// UserAgent is the UserAgent that will be used to send the request
	UserAgent string
	// ProcessOut project ID
	projectID string
	// ProcessOut project secret key
	projectSecret string

	// HTTPClient used to make requests
	HTTPClient *http.Client
}

// Options represents the options available when doing a request to the
// ProcessOut API
type Options struct {
	IdempotencyKey string   `json:"-"`
	Expand         []string `json:"expand"`
	Filter         string   `json:"filter"`
	Limit          uint64   `json:"limit"`
	EndBefore      string   `json:"end_before"`
	StartAfter     string   `json:"start_after"`
	DisableLogging bool     `json:"-"`
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    RequestAPIVersion,
		HTTPClient:    DefaultClient,
		projectID:     projectID,
		projectSecret: projectSecret,
	}

	return p
}

func setupRequest(client *ProcessOut, opt *Options, req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-Version", client.APIVersion)
	req.Header.Set("User-Agent", "ProcessOut Go-Bindings/v4.29.0")
	req.Header.Set("Accept", "application/json")
	if client.UserAgent != "" {
		req.Header.Set("User-Agent", client.UserAgent)
	}
	if opt.IdempotencyKey != "" {
		req.Header.Set("Idempotency-Key", opt.IdempotencyKey)
	}
	if opt.DisableLogging {
		req.Header.Set("Disable-Logging", "true")
	}
	req.SetBasicAuth(client.projectID, client.projectSecret)

	v := req.URL.Query()
	v.Set("filter", opt.Filter)
	v.Set("limit", fmt.Sprint(opt.Limit))
	v.Set("end_before", opt.EndBefore)
	v.Set("start_after", opt.StartAfter)
	for _, e := range opt.Expand {
		v.Add("expand[]", e)
	}
	req.URL.RawQuery = v.Encode()
}

// NewActivity creates a new Activity object
func (c *ProcessOut) NewActivity(prefill ...*Activity) *Activity {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Activity, or none.")
	}
	if len(prefill) == 0 {
		return &Activity{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAddon creates a new Addon object
func (c *ProcessOut) NewAddon(prefill ...*Addon) *Addon {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Addon, or none.")
	}
	if len(prefill) == 0 {
		return &Addon{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAPIRequest creates a new APIRequest object
func (c *ProcessOut) NewAPIRequest(prefill ...*APIRequest) *APIRequest {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the APIRequest, or none.")
	}
	if len(prefill) == 0 {
		return &APIRequest{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAPIVersion creates a new APIVersion object
func (c *ProcessOut) NewAPIVersion(prefill ...*APIVersion) *APIVersion {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the APIVersion, or none.")
	}
	if len(prefill) == 0 {
		return &APIVersion{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewApplePayAlternativeMerchantCertificates creates a new ApplePayAlternativeMerchantCertificates object
func (c *ProcessOut) NewApplePayAlternativeMerchantCertificates(prefill ...*ApplePayAlternativeMerchantCertificates) *ApplePayAlternativeMerchantCertificates {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the ApplePayAlternativeMerchantCertificates, or none.")
	}
	if len(prefill) == 0 {
		return &ApplePayAlternativeMerchantCertificates{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewAlternativeMerchantCertificate creates a new AlternativeMerchantCertificate object
func (c *ProcessOut) NewAlternativeMerchantCertificate(prefill ...*AlternativeMerchantCertificate) *AlternativeMerchantCertificate {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the AlternativeMerchantCertificate, or none.")
	}
	if len(prefill) == 0 {
		return &AlternativeMerchantCertificate{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewBalances creates a new Balances object
func (c *ProcessOut) NewBalances(prefill ...*Balances) *Balances {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Balances, or none.")
	}
	if len(prefill) == 0 {
		return &Balances{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewBalance creates a new Balance object
func (c *ProcessOut) NewBalance(prefill ...*Balance) *Balance {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Balance, or none.")
	}
	if len(prefill) == 0 {
		return &Balance{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCard creates a new Card object
func (c *ProcessOut) NewCard(prefill ...*Card) *Card {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Card, or none.")
	}
	if len(prefill) == 0 {
		return &Card{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCardInformation creates a new CardInformation object
func (c *ProcessOut) NewCardInformation(prefill ...*CardInformation) *CardInformation {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CardInformation, or none.")
	}
	if len(prefill) == 0 {
		return &CardInformation{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCoupon creates a new Coupon object
func (c *ProcessOut) NewCoupon(prefill ...*Coupon) *Coupon {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Coupon, or none.")
	}
	if len(prefill) == 0 {
		return &Coupon{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCustomer creates a new Customer object
func (c *ProcessOut) NewCustomer(prefill ...*Customer) *Customer {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Customer, or none.")
	}
	if len(prefill) == 0 {
		return &Customer{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewToken creates a new Token object
func (c *ProcessOut) NewToken(prefill ...*Token) *Token {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Token, or none.")
	}
	if len(prefill) == 0 {
		return &Token{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewDiscount creates a new Discount object
func (c *ProcessOut) NewDiscount(prefill ...*Discount) *Discount {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Discount, or none.")
	}
	if len(prefill) == 0 {
		return &Discount{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewEvent creates a new Event object
func (c *ProcessOut) NewEvent(prefill ...*Event) *Event {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Event, or none.")
	}
	if len(prefill) == 0 {
		return &Event{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewGateway creates a new Gateway object
func (c *ProcessOut) NewGateway(prefill ...*Gateway) *Gateway {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Gateway, or none.")
	}
	if len(prefill) == 0 {
		return &Gateway{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewGatewayConfiguration creates a new GatewayConfiguration object
func (c *ProcessOut) NewGatewayConfiguration(prefill ...*GatewayConfiguration) *GatewayConfiguration {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the GatewayConfiguration, or none.")
	}
	if len(prefill) == 0 {
		return &GatewayConfiguration{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoice creates a new Invoice object
func (c *ProcessOut) NewInvoice(prefill ...*Invoice) *Invoice {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Invoice, or none.")
	}
	if len(prefill) == 0 {
		return &Invoice{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceTax creates a new InvoiceTax object
func (c *ProcessOut) NewInvoiceTax(prefill ...*InvoiceTax) *InvoiceTax {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceTax, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceTax{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceExternalFraudTools creates a new InvoiceExternalFraudTools object
func (c *ProcessOut) NewInvoiceExternalFraudTools(prefill ...*InvoiceExternalFraudTools) *InvoiceExternalFraudTools {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceExternalFraudTools, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceExternalFraudTools{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceRisk creates a new InvoiceRisk object
func (c *ProcessOut) NewInvoiceRisk(prefill ...*InvoiceRisk) *InvoiceRisk {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceRisk, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceRisk{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceDevice creates a new InvoiceDevice object
func (c *ProcessOut) NewInvoiceDevice(prefill ...*InvoiceDevice) *InvoiceDevice {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceDevice, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceDevice{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceShipping creates a new InvoiceShipping object
func (c *ProcessOut) NewInvoiceShipping(prefill ...*InvoiceShipping) *InvoiceShipping {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceShipping, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceShipping{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewInvoiceDetail creates a new InvoiceDetail object
func (c *ProcessOut) NewInvoiceDetail(prefill ...*InvoiceDetail) *InvoiceDetail {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceDetail, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceDetail{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewCustomerAction creates a new CustomerAction object
func (c *ProcessOut) NewCustomerAction(prefill ...*CustomerAction) *CustomerAction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CustomerAction, or none.")
	}
	if len(prefill) == 0 {
		return &CustomerAction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewDunningAction creates a new DunningAction object
func (c *ProcessOut) NewDunningAction(prefill ...*DunningAction) *DunningAction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the DunningAction, or none.")
	}
	if len(prefill) == 0 {
		return &DunningAction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPayout creates a new Payout object
func (c *ProcessOut) NewPayout(prefill ...*Payout) *Payout {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Payout, or none.")
	}
	if len(prefill) == 0 {
		return &Payout{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPayoutItem creates a new PayoutItem object
func (c *ProcessOut) NewPayoutItem(prefill ...*PayoutItem) *PayoutItem {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the PayoutItem, or none.")
	}
	if len(prefill) == 0 {
		return &PayoutItem{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPlan creates a new Plan object
func (c *ProcessOut) NewPlan(prefill ...*Plan) *Plan {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Plan, or none.")
	}
	if len(prefill) == 0 {
		return &Plan{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewProduct creates a new Product object
func (c *ProcessOut) NewProduct(prefill ...*Product) *Product {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Product, or none.")
	}
	if len(prefill) == 0 {
		return &Product{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewProject creates a new Project object
func (c *ProcessOut) NewProject(prefill ...*Project) *Project {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Project, or none.")
	}
	if len(prefill) == 0 {
		return &Project{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewRefund creates a new Refund object
func (c *ProcessOut) NewRefund(prefill ...*Refund) *Refund {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Refund, or none.")
	}
	if len(prefill) == 0 {
		return &Refund{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewSubscription creates a new Subscription object
func (c *ProcessOut) NewSubscription(prefill ...*Subscription) *Subscription {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Subscription, or none.")
	}
	if len(prefill) == 0 {
		return &Subscription{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewTransaction creates a new Transaction object
func (c *ProcessOut) NewTransaction(prefill ...*Transaction) *Transaction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Transaction, or none.")
	}
	if len(prefill) == 0 {
		return &Transaction{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewThreeDS creates a new ThreeDS object
func (c *ProcessOut) NewThreeDS(prefill ...*ThreeDS) *ThreeDS {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the ThreeDS, or none.")
	}
	if len(prefill) == 0 {
		return &ThreeDS{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPaymentDataThreeDSRequest creates a new PaymentDataThreeDSRequest object
func (c *ProcessOut) NewPaymentDataThreeDSRequest(prefill ...*PaymentDataThreeDSRequest) *PaymentDataThreeDSRequest {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the PaymentDataThreeDSRequest, or none.")
	}
	if len(prefill) == 0 {
		return &PaymentDataThreeDSRequest{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPaymentDataNetworkAuthentication creates a new PaymentDataNetworkAuthentication object
func (c *ProcessOut) NewPaymentDataNetworkAuthentication(prefill ...*PaymentDataNetworkAuthentication) *PaymentDataNetworkAuthentication {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the PaymentDataNetworkAuthentication, or none.")
	}
	if len(prefill) == 0 {
		return &PaymentDataNetworkAuthentication{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewPaymentDataThreeDSAuthentication creates a new PaymentDataThreeDSAuthentication object
func (c *ProcessOut) NewPaymentDataThreeDSAuthentication(prefill ...*PaymentDataThreeDSAuthentication) *PaymentDataThreeDSAuthentication {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the PaymentDataThreeDSAuthentication, or none.")
	}
	if len(prefill) == 0 {
		return &PaymentDataThreeDSAuthentication{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewTransactionOperation creates a new TransactionOperation object
func (c *ProcessOut) NewTransactionOperation(prefill ...*TransactionOperation) *TransactionOperation {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the TransactionOperation, or none.")
	}
	if len(prefill) == 0 {
		return &TransactionOperation{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewWebhook creates a new Webhook object
func (c *ProcessOut) NewWebhook(prefill ...*Webhook) *Webhook {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Webhook, or none.")
	}
	if len(prefill) == 0 {
		return &Webhook{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}

// NewWebhookEndpoint creates a new WebhookEndpoint object
func (c *ProcessOut) NewWebhookEndpoint(prefill ...*WebhookEndpoint) *WebhookEndpoint {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the WebhookEndpoint, or none.")
	}
	if len(prefill) == 0 {
		return &WebhookEndpoint{
			client: c,
		}
	}

	prefill[0].client = c
	return prefill[0]
}
