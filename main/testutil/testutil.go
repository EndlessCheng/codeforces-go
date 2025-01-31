package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type ioFunc func(io.Reader, io.Writer)

func isTLE(f func()) bool {
	if DebugTLE == 0 || IsDebugging() {
		f()
		return false
	}

	done := make(chan struct{})
	timer := time.NewTimer(DebugTLE)
	defer timer.Stop()
	go func() {
		defer close(done)
		f()
	}()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}

func AssertEqualStringCaseWithPrefix(t *testing.T, testCases [][2]string, targetCaseNum int, runFunc ioFunc, prefix string) {
	if len(testCases) == 0 {
		t.Error("empty testcase")
		return
	}

	// 例如，-1 表示最后一个测试用例
	if targetCaseNum < 0 {
		targetCaseNum += len(testCases) + 1
	}

	allPassed := true
	for curCaseNum, tc := range testCases {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		t.Run(fmt.Sprintf("%sCase %d", prefix, curCaseNum+1), func(t *testing.T) {
			input := removeExtraSpace(tc[0])
			const maxInputSize = 150
			inputInfo := input
			if len(inputInfo) > maxInputSize { // 截断过长的输入
				inputInfo = inputInfo[:maxInputSize] + "..."
			}
			expectedOutput := removeExtraSpace(tc[1])

			mockReader := strings.NewReader(input)
			mockWriter := &strings.Builder{}
			_f := func() { runFunc(mockReader, mockWriter) }
			if targetCaseNum == 0 && isTLE(_f) {
				allPassed = false
				t.Errorf("Time Limit Exceeded %d\nInput:\n%s", curCaseNum+1, inputInfo)
				return
			} 
			if targetCaseNum != 0 {
				_f()
			}

			// 还有剩余未读入的内容
			if mockReader.Len() > 0 {
				t.Log("[警告] 有未读入的内容")
			}

			actualOutput := removeExtraSpace(mockWriter.String())
			if !assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d\nInput:\n%s", curCaseNum+1, inputInfo) {
				allPassed = false
				handleOutput(actualOutput)
			}
		})
	}

	// 若有测试用例未通过，则前面必然会打印一些信息，这里直接返回
	if !allPassed {
		return
	}

	// 若测试的是单个用例，则接着测试所有用例
	if targetCaseNum > 0 {
		t.Logf("case %d is passed", targetCaseNum)
		AssertEqualStringCase(t, testCases, 0, runFunc)
		return
	}
}

func AssertEqualStringCase(t *testing.T, testCases [][2]string, targetCaseNum int, runFunc ioFunc) {
	AssertEqualStringCaseWithPrefix(t, testCases, targetCaseNum, runFunc, "")
}

func AssertEqualFileCaseWithName(t *testing.T, dir, inName, ansName string, targetCaseNum int, runFunc ioFunc) {
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
		t.Log("[WARN] empty test file")
		return
	}

	testCases := make([][2]string, len(inputFilePaths))
	for i, path := range inputFilePaths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][0] = string(data)
	}
	for i, path := range answerFilePaths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		testCases[i][1] = string(data)
	}

	AssertEqualStringCase(t, testCases, targetCaseNum, runFunc)
}

func AssertEqualFileCase(t *testing.T, dir string, targetCaseNum int, runFunc ioFunc) {
	AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", targetCaseNum, runFunc)
}

func AssertEqualCase(t *testing.T, rawText string, targetCaseNum int, runFunc ioFunc) {
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

	AssertEqualStringCase(t, testCases, targetCaseNum, runFunc)
}

func AssertEqual(t *testing.T, rawText string, runFunc ioFunc) {
	AssertEqualCase(t, rawText, 0, runFunc)
}

// 对拍
// runFuncAC 为暴力逻辑或已 AC 逻辑，runFunc 为当前测试的逻辑
func AssertEqualRunResults(t *testing.T, inputs []string, targetCaseNum int, runFuncAC, runFunc ioFunc) {
	if len(inputs) == 0 {
		return
	}

	for curCaseNum, input := range inputs {
		if targetCaseNum > 0 && curCaseNum+1 != targetCaseNum {
			continue
		}

		input = removeExtraSpace(input)
		const maxInputSize = 150
		inputInfo := input
		if len(inputInfo) > maxInputSize { // 截断过长的输入
			inputInfo = inputInfo[:maxInputSize] + "..."
		}

		mockReader := strings.NewReader(input)
		mockWriterAC := &strings.Builder{}
		runFuncAC(mockReader, mockWriterAC)
		expectedOutput := removeExtraSpace(mockWriterAC.String())

		mockReader = strings.NewReader(input)
		mockWriter := &strings.Builder{}
		_f := func() { runFunc(mockReader, mockWriter) }
		if targetCaseNum == 0 && isTLE(_f) {
			t.Errorf("Time Limit Exceeded %d\nInput:\n%s", curCaseNum+1, inputInfo)
			continue
		} else if targetCaseNum != 0 {
			_f()
		}
		actualOutput := removeExtraSpace(mockWriter.String())

		t.Run(fmt.Sprintf("Case %d", curCaseNum+1), func(t *testing.T) {
			assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d\nInput:\n%s", curCaseNum+1, inputInfo)
		})
	}
}

// 无尽对拍模式
// inputGenerator 生成随机测试数据，runFuncAC 为暴力逻辑或已 AC 逻辑，runFunc 为当前测试的逻辑
func AssertEqualRunResultsInf(t *testing.T, inputGenerator func() string, runFuncAC, runFunc ioFunc) {
	for tc := 1; ; tc++ {
		input := inputGenerator()
		input = removeExtraSpace(input)

		mockReader := strings.NewReader(input)
		mockWriterAC := &strings.Builder{}
		runFuncAC(mockReader, mockWriterAC)
		expectedOutput := removeExtraSpace(mockWriterAC.String())

		mockReader = strings.NewReader(input)
		mockWriter := &strings.Builder{}
		if isTLE(func() { runFunc(mockReader, mockWriter) }) {
			t.Errorf("Time Limit Exceeded %d\nInput:\n%s", tc, input)
			continue
		}
		actualOutput := removeExtraSpace(mockWriter.String())

		if DisableLogInput {
			assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d", tc)
		} else {
			assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d\nInput:\n%s", tc, input)
		}

		// 每到 2 的幂次就打印检测了多少个测试数据
		if tc&(tc-1) == 0 {
			t.Logf("%d cases checked.", tc)
		}

		if Once {
			break
		}
	}
}

type OutputChecker func(string) bool

// 无尽验证模式
// inputGenerator 除了返回随机输入数据外，还需要返回一个闭包，这个闭包接收 runFunc 的输出结果，根据输入数据验证输出结果是否正确
func CheckRunResultsInfWithTarget(t *testing.T, inputGenerator func() (string, OutputChecker), targetCaseNum int, runFunc ioFunc) {
	tc := 1
	var input string
	var checker OutputChecker

	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Runtime Error %d\nInput:\n%s", tc, input)
		}
	}()

	for ; ; tc++ {
		input, checker = inputGenerator()
		if targetCaseNum > 0 && tc != targetCaseNum {
			continue
		}

		input = removeExtraSpace(input)
		mockReader := strings.NewReader(input)
		mockWriter := &strings.Builder{}
		if isTLE(func() { runFunc(mockReader, mockWriter) }) {
			t.Errorf("Time Limit Exceeded %d\nInput:\n%s", tc, input)
			continue
		}
		actualOutput := removeExtraSpace(mockWriter.String())

		ok := checker(actualOutput)

		if DisableLogInput {
			assert.Truef(t, ok, "Wrong Answer %d", tc)
		} else {
			assert.Truef(t, ok, "Wrong Answer %d\nInput:\n%s\nOutput:\n%s", tc, input, actualOutput)
		}

		if targetCaseNum > 0 {
			if ok {
				CheckRunResultsInfWithTarget(t, inputGenerator, 0, runFunc)
			}
			return
		}

		// 每到 2 的幂次就打印检测了多少个测试数据
		if tc&(tc-1) == 0 {
			t.Logf("%d cases checked.", tc)
		}

		if Once {
			break
		}
	}
}

func CheckRunResultsInf(t *testing.T, inputGenerator func() (string, OutputChecker), runFunc ioFunc) {
	CheckRunResultsInfWithTarget(t, inputGenerator, 0, runFunc)
}
