package testutil

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func AssertEqualStringCase(t *testing.T, inputs []string, answers []string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	if !assert.Equal(t, len(inputs), len(answers), "missing inputs or answers in test cases.") {
		return
	}

	if len(inputs) == 0 {
		return
	}

	if caseNum < 0 {
		caseNum += len(inputs) + 1
	}

	// TODO: time costs
	ok := true
	for i, input := range inputs {
		if caseNum > 0 && i+1 != caseNum {
			continue
		}
		mockReader := strings.NewReader(input)
		mockWriter := &bytes.Buffer{}
		solveFunc(mockReader, mockWriter)
		actualOutput := mockWriter.String()

		// trim space here, may have issue on special problems.
		answer := answers[i]
		answer = strings.TrimSpace(answer)
		actualOutput = strings.TrimSpace(actualOutput)

		_ok := assert.Equal(t, answer, actualOutput, "please check test case [%d]\nInput:\n%s", i+1, input)
		if !_ok {
			ok = _ok
		}
	}
	if !ok {
		t.Logf("ok? caseNum is [%d]", caseNum)
		return
	}

	if caseNum > 0 {
		t.Logf("case %d is passed.", caseNum)
		AssertEqualStringCase(t, inputs, answers, 0, solveFunc)
		return
	}

	t.Log("OK! SUBMIT!")
}

func AssertEqualFileCase(t *testing.T, dir string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	txtFilePaths, _ := filepath.Glob(filepath.Join(dir, "*.txt"))
	// ans1.txt, ..., in1.txt, ...
	if len(txtFilePaths) == 0 {
		return
	}

	var inputs, answers []string
	for _, path := range txtFilePaths[:len(txtFilePaths)/2] {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		answers = append(answers, string(data))
	}
	for _, path := range txtFilePaths[len(txtFilePaths)/2:] {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		inputs = append(inputs, string(data))
	}

	AssertEqualStringCase(t, inputs, answers, caseNum, solveFunc)
}

func AssertEqualCase(t *testing.T, rawText string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	if rawText[0] == '\n' {
		rawText = rawText[1:]
	}
	examples := strings.Split(rawText, "\ninputCopy\n")
	var inputs, answers []string
	for _, e := range examples {
		splits := strings.Split(e, "\noutputCopy\n")
		inputs = append(inputs, splits[0])
		answers = append(answers, splits[1])
	}

	AssertEqualStringCase(t, inputs, answers, caseNum, solveFunc)
}

func AssertEqual(t *testing.T, rawText string, solveFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, 0, solveFunc)
}
