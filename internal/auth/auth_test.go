package auth_test

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 1234")
	key, err := auth.GetAPIKey(headers)
	if err == nil {
		fmt.Println("")
	}
	want := "1234"
	// heres a comment
	if !reflect.DeepEqual(want, key) {
		t.Fatalf("expected: %v, got: %v", want, key)
	}
}
