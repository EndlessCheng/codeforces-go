package testutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func AssertEqual(t *testing.T, rawText string, solFunc func(io.Reader, io.Writer)) {
	examples := strings.Split(rawText, "\ninputCopy\n")
	var inputs, outputs []string
	for _, e := range examples {
		splits := strings.Split(e, "\noutputCopy\n")
		inputs = append(inputs, splits[0])
		outputs = append(outputs, splits[1])
	}

	for i, input := range inputs {
		buf := &bytes.Buffer{}
		solFunc(strings.NewReader(input), buf)
		actualOutput := buf.String()
		if actualOutput != "" && actualOutput[len(actualOutput)-1] == '\n' {
			actualOutput = actualOutput[:len(actualOutput)-1]
		}
		assert.Equal(t, outputs[i], actualOutput, "Please check test case %d", i+1)
	}
}
