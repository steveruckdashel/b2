package b2

import (
	"os"
	"testing"
)

func TestAuthorize(t *testing.T) {
	acountID := os.Getenv("B2_ID")
	secretKey := os.Getenv("B2_KEY")
	b2, err := Authorize(acountID, secretKey)
	if err != nil {
		t.Fatal(err)
	}
	if b2 != nil {
		t.Logf("%+v", b2)
	}
}
