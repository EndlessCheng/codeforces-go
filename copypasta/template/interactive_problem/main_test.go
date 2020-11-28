package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	//rand.Seed(time.Now().UnixNano())
	type testCase struct {
		input
		guess
		//innerData []int
	}
	testCases := []testCase{

	}
	for i := 0; i < 1e5; i++ {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 9)         // 输入
		a := rg.IntSlice(n, 1, 9) // 需要猜测的对象
		testCases = append(testCases, testCase{
			input: input{n},
			guess: guess{a},
		})
	}

	const queryLimit = 1000
	queryChecker := func(caseNum int, tc testCase) func(req) resp {
		//n := tc.n
		//a := append([]int(nil), tc.ans...)
		_queryCnt := 0
		return func(req req) (resp resp) {
			if caseNum == debugCaseNum {
				Print(req, " ")
				defer func() { Println(resp) }()
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			// ...

			resp.v = -1
			return
		}
	}

	// do test
	if debugCaseNum < 0 {
		debugCaseNum += len(testCases)
	}
	const failedCountLimit = 10
	failedCount := 0
	for i, tc := range testCases {
		caseNum := i + 1
		if debugCaseNum != 0 && caseNum != debugCaseNum {
			continue
		}
		expectedAns := tc.guess
		actualAns := run(tc.input, queryChecker(caseNum, tc))
		if !assert.EqualValues(t, expectedAns, actualAns, "Wrong Answer %d", caseNum) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testRun(t, 0)
	}
}
