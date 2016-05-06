package processout

import (
	"testing"
)

func TestValidateWebhook(t *testing.T) {
	dummyKey := "key-5eec899cc1527dd498f269d6a7eb7d60ab77279c6fb7413d8f5c77b50433d37f"
	cases := []struct {
		ProjectSecret string
		RequestBody   []byte
		Signature     string
		ExpectError   bool
	}{
		{
			ProjectSecret: "",
			ExpectError:   true,
		}, {
			ProjectSecret: dummyKey,
			Signature:     "invalid base64",
			ExpectError:   true,
		}, {
			ProjectSecret: dummyKey,
			RequestBody:   []byte("test"),
			Signature:     "Q7DO+ZJl+eNMEOqdNQGSbSezn1fG1nRWHYuiNueoGfs=",
			ExpectError:   true,
		}, {
			ProjectSecret: dummyKey,
			RequestBody:   []byte("test"),
			Signature:     "2sRTf40uImtulbp1YgES2QOMYuSWyGEM6nhZ5QXQ5kM=",
			ExpectError:   false,
		},
	}

	for i, c := range cases {
		p := New("project-id", c.ProjectSecret)
		err := p.Webhooks.Validate(c.RequestBody, c.Signature)
		if (err != nil) != c.ExpectError {
			t.Errorf("Unexpected error value for test case %d", i)
		}
	}
}
