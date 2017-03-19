package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v3/errors"
)

// TransactionOperation represents the TransactionOperation API object
type TransactionOperation struct {
	// ID is the iD of the transaction operation
	ID *string `json:"id,omitempty"`
	// Transaction is the transaction to which the operation belongs
	Transaction *Transaction `json:"transaction,omitempty"`
	// TransactionID is the iD of the transaction to which the operation belongs
	TransactionID *string `json:"transaction_id,omitempty"`
	// Token is the token that was used by the operation, if any
	Token *Token `json:"token,omitempty"`
	// TokenID is the iD of the token was used by the operation, if any
	TokenID *string `json:"token_id,omitempty"`
	// Card is the card that was used by the operation, if any
	Card *Card `json:"card,omitempty"`
	// CardID is the iD of the card that was used by the operation, if any
	CardID *string `json:"card_id,omitempty"`
	// Amount is the amount of the operation
	Amount *string `json:"amount,omitempty"`
	// IsAttempt is the true if the operation is an attempt, false otherwise
	IsAttempt *bool `json:"is_attempt,omitempty"`
	// HasFailed is the true if the operation has failed, false otherwise
	HasFailed *bool `json:"has_failed,omitempty"`
	// IsAccountable is the true if the operation amount can be accounted for, false otherwise
	IsAccountable *bool `json:"is_accountable,omitempty"`
	// Type is the type of the operation, such as authorization, capture, refund or void
	Type *string `json:"type,omitempty"`
	// ErrorCode is the error code returned when attempting the operation, if any
	ErrorCode *string `json:"error_code,omitempty"`
	// Metadata is the metadata related to the operation, in the form of a dictionary (key-value pair)
	Metadata *map[string]string `json:"metadata,omitempty"`
	// GatewayFee is the gateway fee generated by the operation
	GatewayFee *string `json:"gateway_fee,omitempty"`
	// CreatedAt is the date at which the operation was created
	CreatedAt *time.Time `json:"created_at,omitempty"`

	client *ProcessOut
}

// GetID implements the  Identiable interface
func (s *TransactionOperation) GetID() string {
	if s.ID == nil {
		return ""
	}

	return *s.ID
}

// SetClient sets the client for the TransactionOperation object and its
// children
func (s *TransactionOperation) SetClient(c *ProcessOut) *TransactionOperation {
	if s == nil {
		return s
	}
	s.client = c
	if s.Transaction != nil {
		s.Transaction.SetClient(c)
	}
	if s.Token != nil {
		s.Token.SetClient(c)
	}
	if s.Card != nil {
		s.Card.SetClient(c)
	}

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *TransactionOperation) Prefill(c *TransactionOperation) *TransactionOperation {
	if c == nil {
		return s
	}

	s.ID = c.ID
	s.Transaction = c.Transaction
	s.TransactionID = c.TransactionID
	s.Token = c.Token
	s.TokenID = c.TokenID
	s.Card = c.Card
	s.CardID = c.CardID
	s.Amount = c.Amount
	s.IsAttempt = c.IsAttempt
	s.HasFailed = c.HasFailed
	s.IsAccountable = c.IsAccountable
	s.Type = c.Type
	s.ErrorCode = c.ErrorCode
	s.Metadata = c.Metadata
	s.GatewayFee = c.GatewayFee
	s.CreatedAt = c.CreatedAt

	return s
}

// dummyTransactionOperation is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyTransactionOperation() {
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
