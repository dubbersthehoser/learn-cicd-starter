package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input       map[string]string
		expect      string
		shouldError bool
	}

	cases := []test{
		test{
			input: map[string]string{
				"Authorization": "ApiKey lol",
			},
			expect: "lol",
		},
		test{
			input: map[string]string{
				"Authorization": "ApiKey",
			},
			expect: "",
			shouldError: true,
		},
		test{
			input: map[string]string{
				"noAuth": "wop wop...",
			},
			expect: "",
			shouldError: true,
		},
	}


	for i, _case := range cases {
		header := http.Header{}
		for k, v := range _case.input {
			header.Set(k, v)
		}
		expect := _case.expect
		actual, err := GetAPIKey(header)
		if _case.shouldError {
			if err == nil {
				t.Errorf("case %d, expected error, got nil", i)
			}
			continue
		}
		if err != nil {
			t.Errorf("case %d, unexpected error: %s", i, err)
			continue
		}

		if actual != expect {
			t.Errorf("case %d, expect %s, got %s", i, expect, actual)
		}
	}
}
