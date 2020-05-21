package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_run(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	type testCase struct {
		n   int
		ans int
	}
	// corner cases
	testCases := []testCase{
		{ans: 1e9 - 1},
		{ans: 1e9},
	}
	// small cases
	for i := 1; i <= 1000; i++ {
		testCases = append(testCases, testCase{ans: i})
	}
	// random cases
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		v := 1 + rand.Intn(1e9) // [1,1e9]
		testCases = append(testCases, testCase{ans: v})
	}

	const (
		queryLimit    = 64
		minQueryValue = 1
		maxQueryValue = 1e18
	)
	checkQuery := func(caseNum int, tc testCase) func(int64) bool {
		//n := tc.n
		//expectedAns := tc.ans
		queryCnt := 0
		return func(_q int64) (res bool) {
			q := int(_q)
			if caseNum == debugCaseNum {
				println(q)
			}
			queryCnt++
			if queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			if q < minQueryValue || q > maxQueryValue {
				panic("invalid query arguments")
			}
			// ...
			return false
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
		expectedAns := tc.ans
		actualAns := run(tc.n, checkQuery(caseNum, tc))
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
