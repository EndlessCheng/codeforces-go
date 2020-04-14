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

func AssertEqualStringCase(t *testing.T, testCases [][2]string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	if len(testCases) == 0 {
		return
	}

	// 例如，-1 表示最后一个测试用例
	if caseNum < 0 {
		caseNum += len(testCases) + 1
	}

	allPassed := true
	for curCaseNum, tc := range testCases {
		if caseNum > 0 && curCaseNum+1 != caseNum {
			continue
		}

		input := strings.TrimSpace(tc[0])
		expectedOutput := strings.TrimSpace(tc[1])

		mockReader := strings.NewReader(input)
		mockWriter := &bytes.Buffer{}
		solveFunc(mockReader, mockWriter)
		actualOutput := strings.TrimSpace(mockWriter.String())

		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		if !assert.Equal(t, expectedOutput, actualOutput, "please check test case [%d]\nInput:\n%s", curCaseNum+1, inputInfo) {
			allPassed = false
			handleOutput(actualOutput)
		}
	}
	if !allPassed {
		t.Logf("ok? caseNum is [%d]", caseNum)
		return
	}

	if caseNum > 0 {
		t.Logf("case %d is passed.", caseNum)
		// 单个用例通过，测试所有用例
		AssertEqualStringCase(t, testCases, 0, solveFunc)
		return
	}

	t.Log("OK! SUBMIT!")
}

func AssertEqualFileCase(t *testing.T, dir string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	inputFilePaths, err := filepath.Glob(filepath.Join(dir, "in*.txt"))
	if err != nil {
		t.Fatal(err)
	}
	answerFilePaths, err := filepath.Glob(filepath.Join(dir, "ans*.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if len(inputFilePaths) != len(answerFilePaths) {
		t.Fatal("missing sample files")
	}

	testCases := make([][2]string, len(inputFilePaths))
	for i, path := range inputFilePaths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][0] = string(data)
	}
	for i, path := range answerFilePaths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][1] = string(data)
	}

	AssertEqualStringCase(t, testCases, caseNum, solveFunc)
}

func AssertEqualCase(t *testing.T, rawText string, caseNum int, solveFunc func(io.Reader, io.Writer)) {
	rawText = strings.TrimSpace(rawText)
	if rawText == "" {
		t.Fatal("rawText is empty")
	}

	sepInput := "inputCopy"
	if !strings.Contains(rawText, sepInput) {
		sepInput = "input"
	}
	sepOutput := "outputCopy"
	if !strings.Contains(rawText, sepOutput) {
		sepOutput = "output"
	}

	testCases := [][2]string{}
	examples := strings.Split(rawText, sepInput)
	for _, s := range examples {
		s = strings.TrimSpace(s)
		if s != "" {
			splits := strings.Split(s, sepOutput)
			testCases = append(testCases, [2]string{splits[0], splits[1]})
		}
	}

	AssertEqualStringCase(t, testCases, caseNum, solveFunc)
}

func AssertEqual(t *testing.T, rawText string, solveFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, 0, solveFunc)
}

// 对拍
// solveFuncAC 为暴力逻辑或已 AC 逻辑，solveFunc 为当前测试的逻辑
func AssertEqualRunResults(t *testing.T, testCases [][2]string, caseNum int, solveFuncAC, solveFunc func(io.Reader, io.Writer)) {
	if len(testCases) == 0 {
		return
	}

	for curCaseNum, tc := range testCases {
		if caseNum > 0 && curCaseNum+1 != caseNum {
			continue
		}

		input := strings.TrimSpace(tc[0])
		mockReader := strings.NewReader(input)
		mockWriterAC := &bytes.Buffer{}
		solveFuncAC(mockReader, mockWriterAC)
		mockReader = strings.NewReader(input)
		mockWriter := &bytes.Buffer{}
		solveFunc(mockReader, mockWriter)

		actualOutputAC := mockWriterAC.String()
		actualOutput := mockWriter.String()

		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		assert.Equal(t, actualOutputAC, actualOutput, "please check test case [%d]\nInput:\n%s", curCaseNum+1, inputInfo)
	}
}
