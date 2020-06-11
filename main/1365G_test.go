package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_run(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	type testCase struct {
		input1365
		guess1365
		a []int64
	}
	testCases := []testCase{
		{
			input1365: input1365{3},
			guess1365: guess1365{[]int64{6, 5, 3}},
			a:         []int64{1, 2, 4},
		},
	}

	const (
		queryLimit    = 13
		minQueryValue = 1
	)
	checkQuery := func(caseNum int, tc testCase) func(req1365) resp1365 {
		n := tc.n
		a := tc.a
		_queryCnt := 0
		return func(req req1365) (resp resp1365) {
			q := req.q
			if caseNum == debugCaseNum {
				Println(req)
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			if len(q) < minQueryValue || len(q) > n {
				panic("invalid query arguments")
			}
			or := int64(0)
			for _, id := range q {
				or |= a[id-1]
			}
			resp.or = or
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
		expectedAns := tc.guess1365
		actualAns := run1365(tc.input1365, checkQuery(caseNum, tc))
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
