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

// PayoutItemAmountBreakdowns represents the PayoutItemAmountBreakdowns API object
type PayoutItemAmountBreakdowns struct {
	// SchemeFee is the amount relating to scheme fee
	SchemeFee *string `json:"scheme_fee,omitempty"`
	// InterchangeFee is the amount relating to interchange fee
	InterchangeFee *string `json:"interchange_fee,omitempty"`
	// GatewayFee is the amount relating to gateway fee
	GatewayFee *string `json:"gateway_fee,omitempty"`
	// MarkupFee is the amount relating to markup fee
	MarkupFee *string `json:"markup_fee,omitempty"`
	// AcquirerFee is the amount relating to acquirer fee
	AcquirerFee *string `json:"acquirer_fee,omitempty"`
	// OtherFee is the amount relating to other fee
	OtherFee *string `json:"other_fee,omitempty"`

	client *ProcessOut
}

// SetClient sets the client for the PayoutItemAmountBreakdowns object and its
// children
func (s *PayoutItemAmountBreakdowns) SetClient(c *ProcessOut) *PayoutItemAmountBreakdowns {
	if s == nil {
		return s
	}
	s.client = c

	return s
}

// Prefil prefills the object with data provided in the parameter
func (s *PayoutItemAmountBreakdowns) Prefill(c *PayoutItemAmountBreakdowns) *PayoutItemAmountBreakdowns {
	if c == nil {
		return s
	}

	s.SchemeFee = c.SchemeFee
	s.InterchangeFee = c.InterchangeFee
	s.GatewayFee = c.GatewayFee
	s.MarkupFee = c.MarkupFee
	s.AcquirerFee = c.AcquirerFee
	s.OtherFee = c.OtherFee

	return s
}

// dummyPayoutItemAmountBreakdowns is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyPayoutItemAmountBreakdowns() {
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
