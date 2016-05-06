package processout

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

// Projects manages the Project struct
type Projects struct {
	p *ProcessOut
}

type Project struct {
	// Email : Email of the project
	Email string `json:"email"`
	// ID : Unique ID of the project
	ID string `json:"id"`
	// LogoURL : URL of the project's logo
	LogoURL string `json:"logo_url"`
	// Name : Name of the project
	Name string `json:"name"`
	// SecretKey : Secret key of the project
	SecretKey string `json:"secret_key"`
}

// CreateSupervised : Create a new supervised project which will belong to current project.
func (p Projects) CreateSupervised(project *Project) (*Project, error) {
	type Response struct {
		Project `json:"project"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	body, err := json.Marshal(project)
	if err != nil {
		return nil, err
	}

	path := "/projects/supervised"

	req, err := http.NewRequest(
		"POST",
		Host+path,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("API-Version", p.p.APIVersion)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(p.p.projectID, p.p.projectSecret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := &Response{}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(payload)
	if err != nil {
		return nil, err
	}

	if !payload.Success {
		return nil, errors.New(payload.Message)
	}
	return &payload.Project, nil
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
	}
	errors.New("")
}
