package processout

import (
	"testing"
)

func TestNew(t *testing.T) {
	p := New("project-id", "project-secret")
	if p.APIVersion != APIVersion {
		t.Errorf("Wrong API version")
	}
	if p.projectID != "project-id" {
		t.Errorf("Wrong project ID")
	}
	if p.projectSecret != "project-secret" {
		t.Errorf("Wrong project secret")
	}
	if p.Invoices == nil {
		t.Errorf("Missing managers")
	}
}
