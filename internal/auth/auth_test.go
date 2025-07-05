package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expectedAPIKey := "THIS_IS_A_VERY_SECURE_API_KEY"
	header := http.Header{}

	// should succeed
	header.Set("Authorization", fmt.Sprintf("ApiKey %s", expectedAPIKey))
	actual, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("Failed with error: %v", err)
		return
	}

	if actual != expectedAPIKey {
		t.Error("API keys dont match")
	}

	// should fail
	header.Set("Authorization", "")
	actual, err = GetAPIKey(header)

	if err == nil {
		t.Error("Should have returned error but retured nil error")
	}

	// should fail pt2
	header.Set("Authorization", fmt.Sprintf("Bearer: %s %s", expectedAPIKey, expectedAPIKey))
	actual, err = GetAPIKey(header)

	if err == nil {
		t.Error("Should have returned error but retured nil error")
	}
}
