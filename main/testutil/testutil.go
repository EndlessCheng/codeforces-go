package testutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func AssertEqualCase(t *testing.T, rawText string, useCase int, solFunc func(io.Reader, io.Writer)) {
	if rawText[0] == '\n' {
		rawText = rawText[1:]
	}
	examples := strings.Split(rawText, "\ninputCopy\n")
	var inputs, outputs []string
	for _, e := range examples {
		splits := strings.Split(e, "\noutputCopy\n")
		inputs = append(inputs, splits[0])
		outputs = append(outputs, splits[1])
	}

	// TODO: time costs
	ok := true
	for i, input := range inputs {
		if useCase >= 0 && i+1 != useCase {
			continue
		}
		mockReader := strings.NewReader(input)
		mockWriter := &bytes.Buffer{}
		solFunc(mockReader, mockWriter)
		actualOutput := mockWriter.String()
		if actualOutput != "" && actualOutput[len(actualOutput)-1] == '\n' {
			actualOutput = actualOutput[:len(actualOutput)-1]
		}
		_ok := assert.Equal(t, outputs[i], actualOutput, "Please check test case %d\nInput:\n%s", i+1, input)
		if !_ok {
			ok = _ok
		}
	}
	if ok {
		if useCase >= 0 {
			t.Skip("OK, now try to test all cases!")
		} else {
			t.Log("OK, submit with main()!")
		}
	}
}

func AssertEqual(t *testing.T, rawText string, solFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, -1, solFunc)
}
