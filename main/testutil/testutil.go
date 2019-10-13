package testutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func AssertEqualCase(t *testing.T, rawText string, caseNum int, solFunc func(io.Reader, io.Writer)) {
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
		if caseNum >= 0 && i+1 != caseNum {
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
		if caseNum >= 0 {
			t.Skip("OK! Now try to test all cases!")
		} else {
			t.Log("OK! Submit with main()!")
		}
	} else {
		t.Log("OK? Submit with main()!")
	}
}

func AssertEqual(t *testing.T, rawText string, solFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, -1, solFunc)
}
