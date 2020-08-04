package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func AssertEqualStringCase(t *testing.T, testCases [][2]string, caseNum int, runFunc func(io.Reader, io.Writer)) {
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

		input := removeExtraSpace(tc[0])
		expectedOutput := removeExtraSpace(tc[1])

		mockReader := strings.NewReader(input)
		mockWriter := &strings.Builder{}
		runFunc(mockReader, mockWriter)
		actualOutput := removeExtraSpace(mockWriter.String())

		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		if !assert.Equal(t, expectedOutput, actualOutput, "WA %d\nInput:\n%s", curCaseNum+1, inputInfo) {
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
		AssertEqualStringCase(t, testCases, 0, runFunc)
		return
	}

	t.Log("OK! SUBMIT!")
}

func AssertEqualFileCaseWithName(t *testing.T, dir, inName, ansName string, caseNum int, runFunc func(io.Reader, io.Writer)) {
	inputFilePaths, err := filepath.Glob(filepath.Join(dir, inName))
	if err != nil {
		t.Fatal(err)
	}
	answerFilePaths, err := filepath.Glob(filepath.Join(dir, ansName))
	if err != nil {
		t.Fatal(err)
	}
	if len(inputFilePaths) != len(answerFilePaths) {
		t.Fatal("missing sample files")
	}
	if len(inputFilePaths) == 0 {
		t.Log("empty test file")
		return
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

	AssertEqualStringCase(t, testCases, caseNum, runFunc)
}

func AssertEqualFileCase(t *testing.T, dir string, caseNum int, runFunc func(io.Reader, io.Writer)) {
	AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", caseNum, runFunc)
}

func AssertEqualCase(t *testing.T, rawText string, caseNum int, runFunc func(io.Reader, io.Writer)) {
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

	AssertEqualStringCase(t, testCases, caseNum, runFunc)
}

func AssertEqual(t *testing.T, rawText string, runFunc func(io.Reader, io.Writer)) {
	AssertEqualCase(t, rawText, 0, runFunc)
}

// 对拍
// solveFuncAC 为暴力逻辑或已 AC 逻辑，solveFunc 为当前测试的逻辑
func AssertEqualRunResults(t *testing.T, inputs []string, caseNum int, runFuncAC, runFunc func(io.Reader, io.Writer)) {
	if len(inputs) == 0 {
		return
	}

	for curCaseNum, input := range inputs {
		if caseNum > 0 && curCaseNum+1 != caseNum {
			continue
		}

		input = removeExtraSpace(input)
		mockReader := strings.NewReader(input)
		mockWriterAC := &strings.Builder{}
		runFuncAC(mockReader, mockWriterAC)
		mockReader = strings.NewReader(input)
		mockWriter := &strings.Builder{}
		runFunc(mockReader, mockWriter)

		actualOutputAC := removeExtraSpace(mockWriterAC.String())
		actualOutput := removeExtraSpace(mockWriter.String())

		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize {
			inputInfo = inputInfo[:maxInputSize] + "..."
		}
		assert.Equal(t, actualOutputAC, actualOutput, "WA %d\nInput:\n%s", curCaseNum+1, inputInfo)
	}
}

// 无尽对拍模式
func AssertEqualRunResultsInf(t *testing.T, inputGenerator func() string, runFuncAC, runFunc func(io.Reader, io.Writer)) {
	const needPrint = runtime.GOOS == "darwin"

	for tc := 1; ; tc++ {
		input := inputGenerator()
		input = removeExtraSpace(input)
		mockReader := strings.NewReader(input)
		mockWriterAC := &strings.Builder{}
		runFuncAC(mockReader, mockWriterAC)
		mockReader = strings.NewReader(input)
		mockWriter := &strings.Builder{}
		//t0 := time.Now()
		runFunc(mockReader, mockWriter)
		//fmt.Println(time.Since(t0))

		actualOutputAC := removeExtraSpace(mockWriterAC.String())
		actualOutput := removeExtraSpace(mockWriter.String())
		if !assert.Equal(t, actualOutputAC, actualOutput, "WA %d\nInput:\n%s", tc, input) && needPrint {
			fmt.Printf("[CASE %d]\n", tc)
			fmt.Println("[AC]", actualOutputAC)
			fmt.Println("[WA]", actualOutput)
			fmt.Println(input)
			fmt.Println()
		}

		if tc%1e5 == 0 {
			s := fmt.Sprintf("%d cases passed.", tc)
			if needPrint {
				fmt.Println(s)
			}
			t.Log(s)
		}

		//break
	}
}
