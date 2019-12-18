package day11

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
			`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`,
			"19",
		},
		{
			`-1 -1 0 -9 -2 -2
-2 -1 -6 -8 -2 -5
-1 -1 -1 -2 -3 -4
-1 -9 -2 -4 -4 -5
-7 -3 -3 -2 -9 -9
-1 -3 -1 -2 -4 -5`,
			"-6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var stdin bytes.Buffer
			var stdout bytes.Buffer
			stdin.Write([]byte(tt.input))

			run(&stdin, &stdout)

			result, err := stdout.ReadString('%')
			if err != io.EOF {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
