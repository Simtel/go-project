package armisimtel

import (
	"errors"
	"net/http/httptest"
	"testing"
)

func TestBind(t *testing.T) {
	tests := []struct {
		name     string
		payload  DomainPayload
		expected error
	}{
		{
			name:     "Empty name",
			payload:  DomainPayload{Name: ""},
			expected: errors.New("name is required"),
		},
		{
			name:     "Valid name",
			payload:  DomainPayload{Name: "TestName"},
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/", nil)

			err := test.payload.Bind(req)

			if (err != nil) && err.Error() != test.expected.Error() {
				t.Errorf("expected error: %v, got: %v", test.expected, err)
			}
			if (err == nil) && (test.expected != nil) {
				t.Errorf("expected error: %v, got: %v", test.expected, err)
			}
		})
	}
}
