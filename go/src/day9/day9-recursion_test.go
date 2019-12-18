package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		input          int32
		expectedOutput int32
	}{
		{
			3,
			6,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.input), func(t *testing.T) {

			result := factorial(tt.input)

			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
