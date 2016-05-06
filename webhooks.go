package processout

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// Webhooks manages webhooks
type Webhooks struct {
	p *ProcessOut
}

// ValidateWebhook authenticates the origin of a webhook from ProcessOut. It
// takes the raw body of the request and the MAC from the request headers.
func (w Webhooks) Validate(requestBody []byte, signature string) error {
	mac := hmac.New(sha256.New, []byte(w.p.projectSecret))
	mac.Write(requestBody)
	expectedMAC := mac.Sum(nil)
	actualMAC, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	if !hmac.Equal(actualMAC, expectedMAC) {
		return errors.New("Invalid message authentication code")
	}
	return nil
}
