package processout

const (
	// APIVersion is the version of the API this package uses
	APIVersion = "1.3.0.0"
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
	Gatewaies             *Gatewaies
	GatewayConfigurations *GatewayConfigurations
	Invoices              *Invoices
	CustomerActions       *CustomerActions
	Projects              *Projects
	Refunds               *Refunds
	Subscriptions         *Subscriptions
	TailoredInvoices      *TailoredInvoices
	Transactions          *Transactions
	Webhooks              *Webhooks
}

// Options represents the options available when doing a request to the
// ProcessOut API
type Options struct {
	IdempotencyKey string
	Expand         []string
	Filter         string
	Limit          uint64
	Page           uint64
}

// Error represents an error coming from the ProcessOut API. It inherits
// error and adds a Code field, corresponding to the error code coming from
// the API
type Error struct {
	error
	Code string
}

// newError creates a new Error from an error, with an empty code
func newError(err error) *Error {
	return &Error{
		error: err,
		Code:  "",
	}
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
	p.Gatewaies = &Gatewaies{p: p}
	p.GatewayConfigurations = &GatewayConfigurations{p: p}
	p.Invoices = &Invoices{p: p}
	p.CustomerActions = &CustomerActions{p: p}
	p.Projects = &Projects{p: p}
	p.Refunds = &Refunds{p: p}
	p.Subscriptions = &Subscriptions{p: p}
	p.TailoredInvoices = &TailoredInvoices{p: p}
	p.Transactions = &Transactions{p: p}
	p.Webhooks = &Webhooks{p: p}

	return p
}
