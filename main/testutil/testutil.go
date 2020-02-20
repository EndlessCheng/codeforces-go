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

	// 例如，-1 表示最后一个测试用例
	if caseNum < 0 {
		caseNum += len(inputs) + 1
	}

	allPassed := true
	for curCase, input := range inputs {
		if caseNum > 0 && curCase+1 != caseNum {
			continue
		}

		input = strings.TrimSpace(input)
		mockReader := strings.NewReader(input)
		mockWriter := &bytes.Buffer{}
		solveFunc(mockReader, mockWriter)
		actualOutput := mockWriter.String()

		expectedOutput := strings.TrimSpace(answers[curCase])
		actualOutput = strings.TrimSpace(actualOutput)
		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		if !assert.Equal(t, expectedOutput, actualOutput, "please check test case [%d]\nInput:\n%s", curCase+1, inputInfo) {
			allPassed = false
		}
	}
	if !allPassed {
		t.Logf("ok? caseNum is [%d]", caseNum)
		return
	}

	if caseNum > 0 {
		t.Logf("case %d is passed.", caseNum)
		// 单个用例通过，测试所有用例
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
	rawText = strings.TrimSpace(rawText)
	if rawText == "" {
		t.Fatal("rawText is empty")
	}

	examples := strings.Split(rawText, "inputCopy")
	var inputs, answers []string
	for _, e := range examples {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}
		splits := strings.Split(e, "outputCopy")
		in := strings.TrimSpace(splits[0])
		out := strings.TrimSpace(splits[1])
		inputs = append(inputs, in)
		answers = append(answers, out)
	}

	AssertEqualStringCase(t, inputs, answers, caseNum, solveFunc)
}

func AssertEqual(t *testing.T, rawText string, solveFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, 0, solveFunc)
}
