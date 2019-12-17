package day4

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
			`4
-1
10
16
18`,
			`Age is not valid, setting age to 0.
You are young.
You are young.

You are young.
You are a teenager.

You are a teenager.
You are old.

You are old.
You are old.

`,
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
