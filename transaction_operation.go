package processout

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/processout.v5/errors"
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
	// GatewayConfiguration is the gateway configuration that was used to process the operation
	GatewayConfiguration *GatewayConfiguration `json:"gateway_configuration,omitempty"`
	// GatewayConfigurationID is the iD of the gateway configuration that was used to process the operation
	GatewayConfigurationID *string `json:"gateway_configuration_id,omitempty"`
	// Amount is the amount of the operation
	Amount *string `json:"amount,omitempty"`
	// Currency is the currency of the operation
	Currency *string `json:"currency,omitempty"`
	// IsAttempt is the true if the operation is an attempt, false otherwise
	IsAttempt *bool `json:"is_attempt,omitempty"`
	// HasFailed is the true if the operation has failed, false otherwise
	HasFailed *bool `json:"has_failed,omitempty"`
	// IsAccountable is the true if the operation amount can be accounted for, false otherwise
	IsAccountable *bool `json:"is_accountable,omitempty"`
	// Type is the type of the operation, such as authorization, capture, refund or void
	Type *string `json:"type,omitempty"`
	// GatewayOperationID is the iD of the operation done through the PSP
	GatewayOperationID *string `json:"gateway_operation_id,omitempty"`
	// Arn is the acquirer Routing Number, can be used to track a payment or refund at the issuer
	Arn *string `json:"arn,omitempty"`
	// ErrorCode is the error code returned when attempting the operation, if any
	ErrorCode *string `json:"error_code,omitempty"`
	// ErrorMessage is the error message returned when attempting the operation, if any
	ErrorMessage *string `json:"error_message,omitempty"`
	// GatewayData is the additionnal context saved when processing the transaction on the specific PSP
	GatewayData *map[string]string `json:"gateway_data,omitempty"`
	// PaymentDataThreeDSRequest is the threeDS request payment data (read-only)
	PaymentDataThreeDSRequest *PaymentDataThreeDSRequest `json:"payment_data_three_d_s_request,omitempty"`
	// PaymentDataThreeDSAuthentication is the 3-D Secure authentication payment data (read-only)
	PaymentDataThreeDSAuthentication *PaymentDataThreeDSAuthentication `json:"payment_data_three_d_s_authentication,omitempty"`
	// PaymentDataNetworkAuthentication is the network authentication payment data (read-only)
	PaymentDataNetworkAuthentication *PaymentDataNetworkAuthentication `json:"payment_data_network_authentication,omitempty"`
	// InitialSchemeTransactionID is the initial scheme ID that was referenced in the request
	InitialSchemeTransactionID *string `json:"initial_scheme_transaction_id,omitempty"`
	// SchemeID is the the ID assigned to the transaction by the scheme in the last successful authorization
	SchemeID *string `json:"scheme_id,omitempty"`
	// ProcessedWithNetworkToken is the indicates whether the transaction was processed with a network token instead of raw card details
	ProcessedWithNetworkToken *bool `json:"processed_with_network_token,omitempty"`
	// PaymentType is the payment type of the transaction
	PaymentType *string `json:"payment_type,omitempty"`
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
	if s.GatewayConfiguration != nil {
		s.GatewayConfiguration.SetClient(c)
	}
	if s.PaymentDataThreeDSRequest != nil {
		s.PaymentDataThreeDSRequest.SetClient(c)
	}
	if s.PaymentDataThreeDSAuthentication != nil {
		s.PaymentDataThreeDSAuthentication.SetClient(c)
	}
	if s.PaymentDataNetworkAuthentication != nil {
		s.PaymentDataNetworkAuthentication.SetClient(c)
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
	s.GatewayConfiguration = c.GatewayConfiguration
	s.GatewayConfigurationID = c.GatewayConfigurationID
	s.Amount = c.Amount
	s.Currency = c.Currency
	s.IsAttempt = c.IsAttempt
	s.HasFailed = c.HasFailed
	s.IsAccountable = c.IsAccountable
	s.Type = c.Type
	s.GatewayOperationID = c.GatewayOperationID
	s.Arn = c.Arn
	s.ErrorCode = c.ErrorCode
	s.ErrorMessage = c.ErrorMessage
	s.GatewayData = c.GatewayData
	s.PaymentDataThreeDSRequest = c.PaymentDataThreeDSRequest
	s.PaymentDataThreeDSAuthentication = c.PaymentDataThreeDSAuthentication
	s.PaymentDataNetworkAuthentication = c.PaymentDataNetworkAuthentication
	s.InitialSchemeTransactionID = c.InitialSchemeTransactionID
	s.SchemeID = c.SchemeID
	s.ProcessedWithNetworkToken = c.ProcessedWithNetworkToken
	s.PaymentType = c.PaymentType
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
