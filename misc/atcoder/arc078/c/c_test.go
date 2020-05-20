package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
)

func Test_run(t *testing.T) {
	testRun(t, 0)
}

func testRun(t *testing.T, debugCaseNum int) {
	// corner cases
	testCases := []int{
		1e9 - 1,
		1e9,
	}
	// small cases
	for i := 1; i <= 1000; i++ {
		testCases = append(testCases, i)
	}
	// random cases
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		v := 1 + rand.Intn(1e9) // [1,1e9]
		testCases = append(testCases, v)
	}

	const (
		queryLimit    = 64
		minQueryValue = 1
		maxQueryValue = 1e18
	)
	checkQuery := func(caseNum int, expectedAns int) func(int) bool {
		queryCnt := 0
		return func(q int) (res bool) {
			if caseNum == debugCaseNum {
				println(q)
			}
			if queryCnt == queryLimit {
				panic("query limit exceeded")
			}
			queryCnt++
			if q < minQueryValue || q > maxQueryValue {
				panic("invalid query args")
			}
			qs := strconv.Itoa(q)
			es := strconv.Itoa(expectedAns)
			return q <= expectedAns && qs <= es || q > expectedAns && qs > es
		}
	}

	// do test
	if debugCaseNum < 0 {
		debugCaseNum += len(testCases)
	}
	const failedCountLimit = 10
	failedCount := 0
	for i, expectedAns := range testCases {
		caseNum := i + 1
		if debugCaseNum != 0 && caseNum != debugCaseNum {
			continue
		}
		actualAns := run(checkQuery(caseNum, expectedAns))
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
// https://atcoder.jp/contests/arc078/tasks/arc078_c
// https://atcoder.jp/contests/arc078/submit?taskScreenName=arc078_c
