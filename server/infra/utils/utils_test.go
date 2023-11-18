package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandomeCode(t *testing.T) {
	code := SmsCode()
	assert.Equal(t, len(code), 6)
}

func TestTimeFormat(t *testing.T) {
	fmt.Println(time.Now().Format("2006-01-02"))
}

func TestExt(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"alfiler_cotter.prt.4", ".prt.4"},
		{"alfiler_cotter.prt", ".prt"},
		{"test.abcd", ".abcd"},
		{"abcd", ""},
		{"abcd.", "."},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Ext(tt.input); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
