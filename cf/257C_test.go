package cf

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawData := `2
2 0
0 2
outputCopy
90.0000000000
inputCopy
3
2 0
0 2
-2 2
outputCopy
135.0000000000
inputCopy
4
2 0
0 2
-2 0
0 -2
outputCopy
270.0000000000
inputCopy
2
2 1
1 2
outputCopy
36.8698976458`
	examples := strings.Split(rawData, "\ninputCopy\n")
	var inputs, outputs []string
	for _, e := range examples {
		splits := strings.Split(e, "\noutputCopy\n")
		inputs = append(inputs, splits[0])
		outputs = append(outputs, splits[1])
	}

	for i, input := range inputs {
		reader = strings.NewReader(input)
		buf := &bytes.Buffer{}
		writer = buf
		Ans257C()
		actualOutput := buf.String()
		if actualOutput != "" && actualOutput[len(actualOutput)-1] == '\n' {
			actualOutput = actualOutput[:len(actualOutput)-1]
		}
		assert.Equal(t, outputs[i], actualOutput)
	}
}
