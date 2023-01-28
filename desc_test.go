package httpstatusdesc

import (
	"testing"
)

func TestDesc(t *testing.T) {
	cases := []struct {
		name     string
		code     int
		expected string
	}{
		{"knowngood", 100, StatusContinue},
		{"unknown", -1, ""},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := Desc(tt.code)
			if actual != tt.expected {
				t.Fatalf("Expected %s but got %s", tt.expected, actual)
			}
		})
	}
}
