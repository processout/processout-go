package processout

var (
	// APIVersion is the version of the API this package uses
	APIVersion = "1.3.0.0"
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
}

// Options represents the options available when doing a request to the
// ProcessOut API
type Options struct {
	IdempotencyKey string
	Expand         []string
	Filter         string
	Limit          uint64
	Page           uint64
	EndBefore      string
	StartAfter     string
	DisableLogging bool
}

// New creates a new struct *ProcessOut with the given API credentials. It
// initializes all the resources available so they can be used immediately.
func New(projectID, projectSecret string) *ProcessOut {
	p := &ProcessOut{
		APIVersion:    APIVersion,
		projectID:     projectID,
		projectSecret: projectSecret,
	}

	return p
}

// NewActivity creates a new Activity object
func (p *ProcessOut) NewActivity(prefill ...*Activity) *Activity {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Activity, or none.")
	}
	if len(prefill) == 0 {
		return &Activity{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewAuthorizationRequest creates a new AuthorizationRequest object
func (p *ProcessOut) NewAuthorizationRequest(prefill ...*AuthorizationRequest) *AuthorizationRequest {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the AuthorizationRequest, or none.")
	}
	if len(prefill) == 0 {
		return &AuthorizationRequest{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewCard creates a new Card object
func (p *ProcessOut) NewCard(prefill ...*Card) *Card {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Card, or none.")
	}
	if len(prefill) == 0 {
		return &Card{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewCardInformation creates a new CardInformation object
func (p *ProcessOut) NewCardInformation(prefill ...*CardInformation) *CardInformation {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CardInformation, or none.")
	}
	if len(prefill) == 0 {
		return &CardInformation{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewCoupon creates a new Coupon object
func (p *ProcessOut) NewCoupon(prefill ...*Coupon) *Coupon {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Coupon, or none.")
	}
	if len(prefill) == 0 {
		return &Coupon{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewCustomer creates a new Customer object
func (p *ProcessOut) NewCustomer(prefill ...*Customer) *Customer {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Customer, or none.")
	}
	if len(prefill) == 0 {
		return &Customer{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewToken creates a new Token object
func (p *ProcessOut) NewToken(prefill ...*Token) *Token {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Token, or none.")
	}
	if len(prefill) == 0 {
		return &Token{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewDiscount creates a new Discount object
func (p *ProcessOut) NewDiscount(prefill ...*Discount) *Discount {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Discount, or none.")
	}
	if len(prefill) == 0 {
		return &Discount{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewEvent creates a new Event object
func (p *ProcessOut) NewEvent(prefill ...*Event) *Event {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Event, or none.")
	}
	if len(prefill) == 0 {
		return &Event{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewGateway creates a new Gateway object
func (p *ProcessOut) NewGateway(prefill ...*Gateway) *Gateway {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Gateway, or none.")
	}
	if len(prefill) == 0 {
		return &Gateway{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewGatewayConfiguration creates a new GatewayConfiguration object
func (p *ProcessOut) NewGatewayConfiguration(prefill ...*GatewayConfiguration) *GatewayConfiguration {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the GatewayConfiguration, or none.")
	}
	if len(prefill) == 0 {
		return &GatewayConfiguration{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewInvoice creates a new Invoice object
func (p *ProcessOut) NewInvoice(prefill ...*Invoice) *Invoice {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Invoice, or none.")
	}
	if len(prefill) == 0 {
		return &Invoice{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewInvoiceDetail creates a new InvoiceDetail object
func (p *ProcessOut) NewInvoiceDetail(prefill ...*InvoiceDetail) *InvoiceDetail {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the InvoiceDetail, or none.")
	}
	if len(prefill) == 0 {
		return &InvoiceDetail{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewCustomerAction creates a new CustomerAction object
func (p *ProcessOut) NewCustomerAction(prefill ...*CustomerAction) *CustomerAction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the CustomerAction, or none.")
	}
	if len(prefill) == 0 {
		return &CustomerAction{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewPlan creates a new Plan object
func (p *ProcessOut) NewPlan(prefill ...*Plan) *Plan {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Plan, or none.")
	}
	if len(prefill) == 0 {
		return &Plan{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewProduct creates a new Product object
func (p *ProcessOut) NewProduct(prefill ...*Product) *Product {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Product, or none.")
	}
	if len(prefill) == 0 {
		return &Product{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewProject creates a new Project object
func (p *ProcessOut) NewProject(prefill ...*Project) *Project {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Project, or none.")
	}
	if len(prefill) == 0 {
		return &Project{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewRefund creates a new Refund object
func (p *ProcessOut) NewRefund(prefill ...*Refund) *Refund {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Refund, or none.")
	}
	if len(prefill) == 0 {
		return &Refund{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewSubscription creates a new Subscription object
func (p *ProcessOut) NewSubscription(prefill ...*Subscription) *Subscription {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Subscription, or none.")
	}
	if len(prefill) == 0 {
		return &Subscription{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewTransaction creates a new Transaction object
func (p *ProcessOut) NewTransaction(prefill ...*Transaction) *Transaction {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Transaction, or none.")
	}
	if len(prefill) == 0 {
		return &Transaction{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}

// NewWebhook creates a new Webhook object
func (p *ProcessOut) NewWebhook(prefill ...*Webhook) *Webhook {
	if len(prefill) > 1 {
		panic("You may only provide one structure used to prefill the Webhook, or none.")
	}
	if len(prefill) == 0 {
		return &Webhook{
			Client: p,
		}
	}

	prefill[0].Client = p
	return prefill[0]
}
