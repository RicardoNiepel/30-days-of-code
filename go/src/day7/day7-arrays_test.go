package day7

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			"4\n1 4 3 2",
			`2 3 4 1`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var stdin bytes.Buffer
			var stdout bytes.Buffer
			stdin.Write([]byte(tt.input))

			run(&stdin, &stdout)

			result, err := stdout.ReadString('-')
			if err != io.EOF {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
