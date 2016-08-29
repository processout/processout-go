package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Projects manages the Project struct
type Projects struct {
	p *ProcessOut
}

type Project struct {
	// ID : ID of the project
	ID string `json:"id"`
	// Name : Name of the project
	Name string `json:"name"`
	// LogoURL : Name of the project
	LogoURL string `json:"logo_url"`
	// Email : Email of the project
	Email string `json:"email"`
	// CreatedAt : Date at which the project was created
	CreatedAt time.Time `json:"created_at"`
}

// dummyProject is a dummy function that's only
// here because some files need specific packages and some don't.
// It's easier to include it for every file. In case you couldn't
// tell, everything is generated.
func dummyProject() {
	type dummy struct {
		a bytes.Buffer
		b json.RawMessage
		c http.File
		d strings.Reader
		e time.Time
		f url.URL
	}
	errors.New("")
}
