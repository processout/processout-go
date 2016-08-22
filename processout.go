package processout

const (
	// APIVersion is the version of the API this package uses
	APIVersion = "1.2.0.0"
)

var (
	// Host is the URL where API requests are made
	Host = "https://api.processout.com"
)

// ProcessOut wraps all the components of the package in a
// single structure
type ProcessOut struct {
	// APIVersion is the version of the API to use
	APIVersion string
	// ProcessOut project ID
	projectID string
	// ProcessOut project secret key
	projectSecret string

	Activities            *Activities
	AuthorizationRequests *AuthorizationRequests
	Customers             *Customers
	Tokens                *Tokens
	Events                *Events
	Invoices              *Invoices
	RecurringInvoices     *RecurringInvoices
	TailoredInvoices      *TailoredInvoices
	Transactions          *Transactions
	Webhooks              *Webhooks
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    APIVersion,
		projectID:     projectID,
		projectSecret: projectSecret,
	}
	p.Activities = &Activities{p: p}
	p.AuthorizationRequests = &AuthorizationRequests{p: p}
	p.Customers = &Customers{p: p}
	p.Tokens = &Tokens{p: p}
	p.Events = &Events{p: p}
	p.Invoices = &Invoices{p: p}
	p.RecurringInvoices = &RecurringInvoices{p: p}
	p.TailoredInvoices = &TailoredInvoices{p: p}
	p.Transactions = &Transactions{p: p}
	p.Webhooks = &Webhooks{p: p}

	return p
}