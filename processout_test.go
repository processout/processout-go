package processout

import "testing"
import "net/http"
import "bytes"

func getClient() *ProcessOut {
	return New("test-proj_gAO1Uu0ysZJvDuUpOGPkUBeE3pGalk3x",
		"key_sandbox_mah31RDFqcDxmaS7MvhDbJfDJvjtsFTB")
}

func TestNew(t *testing.T) {
	p := New("project-id", "project-secret")
	if p.APIVersion != RequestAPIVersion {
		t.Errorf("Wrong API version")
	}
	if p.projectID != "project-id" {
		t.Errorf("Wrong project ID")
	}
	if p.projectSecret != "project-secret" {
		t.Errorf("Wrong project secret")
	}
}

func TestCreateFetchInvoice(t *testing.T) {
	p := getClient()

	iv, err := p.NewInvoice(&Invoice{
		Name:     String("test invoice"),
		Amount:   String("9.99"),
		Currency: String("EUR"),
	}).Create()
	if err != nil {
		t.Errorf("The invoice could not be created: %s", err.Error())
	}
	if iv.ID == nil || *iv.ID == "" {
		t.Errorf("The created invoice ID was empty")
	}

	iv2, err := p.NewInvoice().Find(*iv.ID)
	if err != nil {
		t.Errorf("The invoice could not be fetched: %s", err.Error())
	}
	if *iv.ID != *iv2.ID {
		t.Errorf("The invoice IDs did not match")
	}
}

func TestCaptureInvoice(t *testing.T) {
	p := getClient()
	iv, err := p.NewInvoice(&Invoice{
		Name:     String("test invoice"),
		Amount:   String("9.99"),
		Currency: String("EUR"),
	}).Create()
	if err != nil {
		t.Errorf("The invoice could not be created: %s", err.Error())
	}

	req, _ := http.NewRequest("POST", "https://processout.com?token=test-valid", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	gr := NewGatewayRequest("sandbox", req)
	tr, err := iv.Capture(gr.String())
	if err != nil {
		t.Errorf("The invoice should have been captured, but got: %s", err.Error())
	}
	if tr.Status == nil || *tr.Status != "completed" {
		t.Errorf("The transaction should have been completed, but got: %s", *tr.Status)
	}

	// Check the expand
	tr, err = tr.Find(*tr.ID, TransactionFindParameters{
		Options: &Options{
			Expand: []string{"gateway_configuration"},
		},
	})
	if err != nil {
		t.Errorf("The invoice should have been captured, but got: %s", err.Error())
	}
	if tr.GatewayConfiguration == nil || tr.GatewayConfiguration.ID == nil || *tr.GatewayConfiguration.ID == "" {
		t.Errorf("The transaction gateway configuration was expanded even though we expanded it")
	}
}

func TestGetCustomers(t *testing.T) {
	p := getClient()

	_, err := p.NewCustomer().All()
	if err != nil {
		t.Errorf("The customers list could not be fetched: %s", err.Error())
	}
}

func TestCreateCustomerSubscription(t *testing.T) {
	p := getClient()

	cust, err := p.NewCustomer().Create()
	if err != nil {
		t.Errorf("The customer could not be created: %s", err.Error())
	}
	if cust.ID == nil || *cust.ID == "" {
		t.Errorf("The customer ID should not be empty")
	}

	sub, err := p.NewSubscription(&Subscription{
		CustomerID: cust.ID,
		Name:       String("great subscription"),
		Amount:     String("9.99"),
		Currency:   String("USD"),
		Interval:   String("1d"),
	}).Create()
	if err != nil {
		t.Errorf("The subscription could not be created: %s", err.Error())
	}
	if sub.ID == nil || *sub.ID == "" {
		t.Errorf("The subscription ID should not be empty")
	}
}

func TestCreateCustomerPrefill(t *testing.T) {
	p := getClient()

	tmpl := &Customer{
		Email: String("john@smith.com"),
	}

	cust, err := p.NewCustomer(tmpl).Create(CustomerCreateParameters{
		Options: &Options{
			Expand: []string{"project"},
		},
	})
	if err != nil {
		t.Errorf("The subscription could not be created: %s", err.Error())
	}
	if *cust.Email != *tmpl.Email {
		t.Errorf("The email should be %s but was %s", *tmpl.Email, *cust.Email)
	}
}

func TestCreateCustomerParameters(t *testing.T) {
	p := getClient()

	tmpl := &Customer{
		Email: String("john@smith.com"),
	}

	cust, err := p.NewCustomer().Create(CustomerCreateParameters{
		Options: &Options{
			Expand: []string{"project"},
		},
		Customer: tmpl,
	})
	if err != nil {
		t.Errorf("The subscription could not be created: %s", err.Error())
	}
	if *cust.Email != *tmpl.Email {
		t.Errorf("The email should be %s but was %s", *tmpl.Email, *cust.Email)
	}
}

func TestExpandCustomerProject(t *testing.T) {
	p := getClient()

	cust, _ := p.NewCustomer().Create(CustomerCreateParameters{
		Options: &Options{
			Expand: []string{"project"},
		}})
	if cust.Project == nil {
		t.Errorf("The customer project should be expanded")
	}
}

func TestPaginateCustomersNext(t *testing.T) {
	p := getClient()

	custs, _ := p.NewCustomer().All(CustomerAllParameters{
		Options: &Options{
			Limit: 10,
		},
	})

	seenIDs := []string{}
	i := 0
	for custs.Next() {
		i++

		cust := custs.Get().(*Customer)
		for _, s := range seenIDs {
			if s == *cust.ID {
				t.Fatalf("the customer with ID %s was already found in last iteration", s)
			}
		}
		seenIDs = append(seenIDs, *cust.ID)

		if i > 11 {
			break
		}
	}

	if i < 11 {
		t.Errorf("the iteration count should have been greater than 10")
	}
	if err := custs.Error(); err != nil {
		t.Errorf("there shouldn't have been any error, but got %s", err.Error())
	}
}

func TestPaginateCustomersPrev(t *testing.T) {
	p := getClient()

	custs, _ := p.NewCustomer().All(CustomerAllParameters{
		Options: &Options{
			Limit: 10,
		},
	})
	for custs.Prev() {
		t.Errorf("There shouldnt have been any iteration")
	}
	if err := custs.Error(); err != nil {
		t.Errorf("There shouldn't have been any error, but got %s", err.Error())
	}
}

func TestCreateSupervisedProject(t *testing.T) {
	p := getClient()

	project, err := p.NewProject(&Project{
		Name:            String("test supervised project"),
		DefaultCurrency: String("USD"),
	}).CreateSupervised()
	if err != nil {
		t.Errorf("There shouldn't have been any error, but got %s", err.Error())
	}

	if project.PrivateKey == nil || *project.PrivateKey == "" {
		t.Error("The project private key was empty")
	}
}
