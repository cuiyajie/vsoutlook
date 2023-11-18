package utils

import "testing"

func TestIsWord(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"ab0d", true},
		{"ab_$", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := IsWord(tt.input); got != tt.want {
				t.Errorf("IsWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
