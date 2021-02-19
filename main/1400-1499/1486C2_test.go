package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_CF1486C2(t *testing.T) {
	testCF1486C2(t, 0)
}

func testCF1486C2(t *testing.T, debugCaseNum int) {
	//rand.Seed(time.Now().UnixNano())
	type testCase struct {
		input86
		guess86
		innerData []int
	}
	testCases := []testCase{
		{
			input86:   input86{5},
			guess86:   guess86{1},
			innerData: []int{5, 1, 4, 2, 3},
		},
	}
	for i := 0; i < 1e5; i++ {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 9) // 输入
		a := rg.Permutation(1, n)
		b := append([]int(nil), a...)
		sort.Sort(sort.Reverse(sort.IntSlice(b)))
		p := 0
		for i, v := range a {
			if v == b[0] {
				p = i + 1
				break
			}
		}
		testCases = append(testCases, testCase{
			input86:   input86{n},
			guess86:   guess86{p},
			innerData: a,
		})
	}

	const queryLimit = 20
	queryChecker := func(caseNum int, tc testCase) func(req86) resp86 {
		//n := tc.n
		//a := append([]int(nil), tc.ans...)
		_queryCnt := 0
		return func(req req86) (resp resp86) {
			if caseNum == debugCaseNum {
				Print(req, " ")
				defer func() { Println(resp) }()
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			l, r := req.l, req.r
			if l >= r {
				panic("invalid request")
			}
			b := append([]int(nil), tc.innerData[l-1:r]...)
			sort.Sort(sort.Reverse(sort.IntSlice(b)))
			v := b[1]
			for i, w := range tc.innerData {
				if w == v {
					resp.v = i + 1
					break
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
		expectedAns := tc.guess86
		actualAns := CF1486C2(tc.input86, queryChecker(caseNum, tc))
		if !assert.EqualValues(t, expectedAns, actualAns, "Wrong Answer %d\nInner Data:\n%v", caseNum, tc.innerData) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testCF1486C2(t, 0)
	}
}
