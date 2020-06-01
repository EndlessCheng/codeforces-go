package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCF1363D(t *testing.T) {
	testCF1363D(t, 0)
}

func testCF1363D(t *testing.T, debugCaseNum int) {
	type testCase struct {
		input1363
		guess1363
		a []int
	}
	testCases := []testCase{
		{
			input1363: input1363{4, 1, [][]int{{3, 2}}},
			guess1363: guess1363{[]int{4}},
			a:         []int{1, 2, 3, 4},
		},
		{
			input1363: input1363{4, 2, [][]int{{1, 3}, {2, 4}}},
			guess1363: guess1363{[]int{4, 3}},
			a:         []int{1, 2, 3, 4},
		},
		{
			input1363: input1363{4, 4, [][]int{{1}, {2}, {3}, {4}}},
			guess1363: guess1363{[]int{4, 4, 4, 3}},
			a:         []int{1, 2, 3, 4},
		},
		{
			input1363: input1363{4, 3, [][]int{{1}, {2}, {3, 4}}},
			guess1363: guess1363{[]int{4, 4, 2}},
			a:         []int{1, 2, 3, 4},
		},
	}

	const (
		queryLimit    = 12
		minQueryValue = 1
	)
	checkQuery := func(caseNum int, tc testCase) func(qIn1363) qOut1363 {
		_queryCnt := 0
		return func(qi qIn1363) (resp qOut1363) {
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
		expectedAns := tc.guess1363
		actualAns := CF1363D(tc.input1363, checkQuery(caseNum, tc))
		if !assert.EqualValues(t, expectedAns, actualAns, "WA %d", caseNum) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testCF1363D(t, 0)
	}
}
