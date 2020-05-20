package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func Test_run(t *testing.T) {
	const _debugCaseNum = 0
	testCases := []int{
		1,
		2,
		9,
		10,
		1e9 - 1,
		1e9,
	}
	// append some random data
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
			if caseNum == _debugCaseNum {
				println(q)
			}
			if queryCnt == queryLimit {
				panic("query limit exceeded")
			}
			queryCnt++
			if q < minQueryValue || q > maxQueryValue {
				panic("invalid query args")
			}
			// ...
			return false
		}
	}

	// do test
	for i, expectedAns := range testCases {
		caseNum := i + 1
		if caseNum == _debugCaseNum {
			print()
		}
		actualAns := run(checkQuery(caseNum, expectedAns))
		assert.EqualValues(t, expectedAns, actualAns, "WA %d", caseNum)
	}
}
