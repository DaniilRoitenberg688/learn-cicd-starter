package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
		Header string
		Output string
		Error  error
	}{
		"No token":         {"", "", ErrNoAuthHeaderIncluded},
		"Malformed header": {"lkahsdahsd", "", errors.New("malformed authorization header")},
		"Everything is ok": {"ApiKey token", "token", nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add("Authorization", tc.Header)
			out, err := GetAPIKey(header)
			if !reflect.DeepEqual(out, tc.Output) {
				t.Fatalf("expected: %v, got: %v", tc.Output, out)

			}
			if !reflect.DeepEqual(err, tc.Error) {
				t.Fatalf("expected %v, got: %v", tc.Error, err)
			}

		})
	}
}
