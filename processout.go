package processout

const (
	// APIVersion is the version of the API this package uses
	APIVersion = "1.1.0.0"
	// Host is the URL where API requests are made
	Host = "https://api.processout.com/"
)

// ProcessOut wraps all the components of the package in a
// single structure
type ProcessOut struct {
	// APIVersion is the version of the API to use
	APIVersion string
	// ProcessOut's project id
	projectID string
	// ProcessOut's project secret key
	projectSecret string

	Customers                *Customers
	CustomerActions          *CustomerActions
	CustomerTokens           *CustomerTokens
	Events                   *Events
	Invoices                 *Invoices
	PaymentGateways          *PaymentGateways
	PaymentGatewayPublicKeys *PaymentGatewayPublicKeys
	Projects                 *Projects
	RecurringInvoices        *RecurringInvoices
	TailoredInvoices         *TailoredInvoices
	Webhooks                 *Webhooks
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    APIVersion,
		projectID:     projectID,
		projectSecret: projectSecret,
	}
	p.Customers = &Customers{p: p}
	p.CustomerActions = &CustomerActions{p: p}
	p.CustomerTokens = &CustomerTokens{p: p}
	p.Events = &Events{p: p}
	p.Invoices = &Invoices{p: p}
	p.PaymentGateways = &PaymentGateways{p: p}
	p.PaymentGatewayPublicKeys = &PaymentGatewayPublicKeys{p: p}
	p.Projects = &Projects{p: p}
	p.RecurringInvoices = &RecurringInvoices{p: p}
	p.TailoredInvoices = &TailoredInvoices{p: p}
	p.Webhooks = &Webhooks{p: p}

	return p
}
