package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1479/A
// https://codeforces.com/problemset/status/1479/problem/A
func Test_CF1479A(t *testing.T) {
	testCF1479A(t, 0)
}

func testCF1479A(t *testing.T, debugCaseNum int) {
	//rand.Seed(time.Now().UnixNano())
	type testCase struct {
		input79
		guess79
		innerData []int
	}
	testCases := []testCase{

	}
	for i := 0; i < 5e5; i++ {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 9)
		a := rg.Permutation(1, n)
		var ans int
		if a[0] < a[1] {
			ans = 0
		} else if a[n-2] > a[n-1] {
			ans = n - 1
		}
		for i := 1; i < n-1; i++ {
			if a[i-1] > a[i] && a[i] < a[i+1] {
				ans = i
				break
			}
		}
		ans++
		testCases = append(testCases, testCase{
			input79:   input79{n},
			guess79:   guess79{ans},
			innerData: a,
		})
	}

	const queryLimit = 100
	queryChecker := func(caseNum int, tc testCase) func(req79) resp79 {
		//n := tc.n
		//a := append([]int(nil), tc.ans...)
		_queryCnt := 0
		return func(req req79) (resp resp79) {
			if caseNum == debugCaseNum {
				Print(req, " ")
				defer func() { Println(resp) }()
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			resp.v = tc.innerData[req.i-1]
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
		actualAns := CF1479A(tc.input79, queryChecker(caseNum, tc))
		res := actualAns.ans
		res--
		check := func()bool {
			a := tc.innerData
			n := len(a)
			if res == 0   {
				return a[0] < a[1]
			}
			if res == n-1 {
				return a[n-2] > a[n-1]
			}
			return a[res-1] > a[res] && a[res]<a[res+1]
		}
		if !check() {
			Printf("Wrong Answer %d\n", caseNum)
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testCF1479A(t, 0)
	}
}
