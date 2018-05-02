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

// PayoutItem represents the PayoutItem API object
type PayoutItem struct {
	// ID is the iD of the payout item
	ID *string `json:"id,omitempty"`
	// Project is the project to which the payout item belongs
	Project *Project `json:"project,omitempty"`
	// ProjectID is the iD of the project to which the payout item belongs
	ProjectID *string `json:"project_id,omitempty"`
	// Payout is the payout to which the item belongs
	Payout *Payout `json:"payout,omitempty"`
	// PayoutID is the iD of the payout to which the item belongs
	PayoutID *string `json:"payout_id,omitempty"`
	// Transaction is the transaction linked to this payout item. Can be null
	Transaction *Transaction `json:"transaction,omitempty"`
	// TransactionID is the iD of the transaction linked to this payout item. Can be null
	TransactionID *string `json:"transaction_id,omitempty"`
	// Type is the type of the payout item
	Type *string `json:"type,omitempty"`
	// GatewayID is the iD of the payout item from the payment gateway
	GatewayID *string `json:"gateway_id,omitempty"`
	// Fee is the fee linked to this specific payout item. Can be null or 0.
	Fee *string `json:"fee,omitempty"`
	// Metadata is the metadata related to the payout item, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// CreatedAt is the date at which the payout item was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *PayoutItem) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the PayoutItem object and its
// children
func (s *PayoutItem) SetClient(c *ProcessOut) *PayoutItem {
	if s == nil {
		return s
	}
	s.client = c
	if s.Project != nil {
		s.Project.SetClient(c)
	}
	if s.Payout != nil {
		s.Payout.SetClient(c)
	}
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PayoutItem) Prefill(c *PayoutItem) *PayoutItem {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Project = c.Project
	s.ProjectID = c.ProjectID
	s.Payout = c.Payout
	s.PayoutID = c.PayoutID
	s.Transaction = c.Transaction
	s.TransactionID = c.TransactionID
	s.Type = c.Type
	s.GatewayID = c.GatewayID
	s.Fee = c.Fee
	s.Metadata = c.Metadata
	s.CreatedAt = c.CreatedAt

	return s
}

// dummyPayoutItem is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPayoutItem() {
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
