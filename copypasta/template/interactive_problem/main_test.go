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
	// corner cases
	testCases := []int{
		1,
		2,
		9,
		10,
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
	checkQuery := func(caseNum int, expectedAns int) func(int64) bool {
		queryCnt := 0
		return func(_q int64) bool {
			q := int(_q)
			if caseNum == debugCaseNum {
				println(q)
			}
			if queryCnt == queryLimit {
				panic("query limit exceeded")
			}
			queryCnt++
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
