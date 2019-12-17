package day6

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{
			`2
Hacker
Rank`,
			`Hce akr
Rn ak`,
		},
		{
			`1
Hacker`,
			`Hce akr`,
		},
		{
			`0`,
			``,
		},
		{
			`1
kÖä:;ß`,
			`kä; Ö:ß`,
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

func Test_extended_07(t *testing.T) {
	dat, _ := ioutil.ReadFile("input07.txt")
	input := string(dat)

	dat, _ = ioutil.ReadFile("output07.txt")
	expected := string(dat)

	var stdin bytes.Buffer
	var stdout bytes.Buffer
	stdin.Write([]byte(input))

	run(&stdin, &stdout)

	result, err := stdout.ReadString('-')
	if err != io.EOF {
		assert.NoError(t, err)
	}
	assert.Equal(t, expected, result)
}

func Test_extended_09(t *testing.T) {
	dat, _ := ioutil.ReadFile("input09.txt")
	input := string(dat)

	dat, _ = ioutil.ReadFile("output09.txt")
	expected := string(dat)

	var stdin bytes.Buffer
	var stdout bytes.Buffer
	stdin.Write([]byte(input))

	run(&stdin, &stdout)

	result, err := stdout.ReadString('-')
	if err != io.EOF {
		assert.NoError(t, err)
	}

	assert.Equal(t, expected, result)
}
