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
			input: input{4},
			guess: guess{[]int{1, 2, 1e9 - 1, 1e9}},
		},
	}
	// small cases
	for i := 1; i <= 1000; i++ {
		testCases = append(testCases, testCase{
			input: input{1},
			guess: guess{[]int{i}},
		})
	}
	// random cases
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		v := rand.Intn(1e9) + 1 // [1,1e9]
		testCases = append(testCases, testCase{
			input: input{1},
			guess: guess{[]int{v}},
		})
	}

	// TODO config limits
	const (
		queryLimit    = 64
		minQueryValue = 1
		maxQueryValue = 1e18
	)
	checkQuery := func(caseNum int, tc testCase) func(req) resp {
		//n := tc.n
		_queryCnt := 0
		return func(req req) (resp resp) {
			q := req.q
			if caseNum == debugCaseNum {
				Println(req)
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			if len(q) < minQueryValue || len(q) > maxQueryValue {
				panic("invalid query arguments")
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
