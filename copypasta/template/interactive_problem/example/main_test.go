package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	type testCase struct {
		input
		guess
		a []int
	}
	testCases := []testCase{
		{
			input: input{4, 1, [][]int{{3, 2}}},
			guess: guess{[]int{4}},
			a:     []int{1, 2, 3, 4},
		},
		{
			input: input{4, 2, [][]int{{1, 3}, {2, 4}}},
			guess: guess{[]int{4, 3}},
			a:     []int{1, 2, 3, 4},
		},
		{
			input: input{4, 4, [][]int{{1}, {2}, {3}, {4}}},
			guess: guess{[]int{4, 4, 4, 3}},
			a:     []int{1, 2, 3, 4},
		},
		{
			input: input{4, 3, [][]int{{1}, {2}, {3, 4}}},
			guess: guess{[]int{4, 4, 2}},
			a:     []int{1, 2, 3, 4},
		},
	}

	const (
		queryLimit    = 12
		minQueryValue = 1
	)
	checkQuery := func(caseNum int, tc testCase) func(qIn) qOut {
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
			if len(q) < minQueryValue || len(q) > tc.n {
				panic("invalid query arguments")
			}
			for _, id := range q {
				if id < minQueryValue || id > tc.n {
					panic("invalid query arguments")
				}
				if tc.a[id-1] > resp.max {
					resp.max = tc.a[id-1]
				}
			}
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
