package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_run(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	type testCase struct {
		input
		guess
		innerData []int // optional
	}
	// corner cases
	testCases := []testCase{
		{
			input: input{10},
			guess: guess{1e9},
		},
	}
	// small cases
	for i := 1; i <= 1000; i++ {
		testCases = append(testCases, testCase{
			input: input{10},
			guess: guess{i},
		})
	}
	// random cases
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		v := 1 + rand.Intn(1e9) // [1,1e9]
		testCases = append(testCases, testCase{
			input: input{10},
			guess: guess{v},
		})
	}

	// TODO config limits
	const (
		queryLimit    = 64
		minQueryValue = 1
		maxQueryValue = 1e18
	)
	checkQuery := func(caseNum int, tc testCase) func(qIn) qOut {
		n := tc.n
		numToGuess := tc.ans
		_queryCnt := 0
		return func(qi qIn) (resp qOut) {
			q := qi.q
			if caseNum == debugCaseNum {
				Println(qi)
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			if q < minQueryValue || q > maxQueryValue {
				panic("invalid query arguments")
			}
			// ...
			resp.ok = q >= n+numToGuess
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
		actualAns := run(tc.input, checkQuery(caseNum, tc))
		if !assert.EqualValues(t, expectedAns, actualAns, "WA %d", caseNum) {
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
