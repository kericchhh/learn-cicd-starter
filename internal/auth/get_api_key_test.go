package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("valid API key", func(t *testing.T) {
		headers := http.Header{"Authorization": {"ApiKey test-key"}}

		key, err := GetAPIKey(headers)
		if err != nil {
			t.Fatalf("GetAPIKey() error = %v, want nil", err)
		}
		if key != "test-key" {
			t.Errorf("GetAPIKey() key = %q, want %q", key, "test-key")
		}
	})

	t.Run("missing header", func(t *testing.T) {
		key, err := GetAPIKey(http.Header{})
		if !errors.Is(err, ErrNoAuthHeaderIncluded) {
			t.Errorf("GetAPIKey() error = %v, want %v", err, ErrNoAuthHeaderIncluded)
		}
		if key != "" {
			t.Errorf("GetAPIKey() key = %q, want empty", key)
		}
	})

	t.Run("malformed header", func(t *testing.T) {
		headers := http.Header{"Authorization": {"Bearer test-key"}}

		key, err := GetAPIKey(headers)
		if err == nil {
			t.Error("GetAPIKey() error = nil, want an error")
		}
		if key != "" {
			t.Errorf("GetAPIKey() key = %q, want empty", key)
		}
	})
}
