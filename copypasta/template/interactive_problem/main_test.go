package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func init() { rand.Seed(time.Now().UnixNano()) }

func Test_run(t *testing.T) { testRun(t, 0) }

func testRun(t *testing.T, debugCaseNum int) {
	assert := assert.New(t)

	type testCase struct {
		input
		guess
		innerData []int
	}
	format := func(tc testCase) (s string) {
		s = Sprintf("%v", tc.innerData)
		//s = strings.Join(tc.innerData, "\n")
		return
	}

	testCases := []testCase{}
	for tc := 0; tc < 3e5; tc++ {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 4)         // 输入
		a := rg.IntSlice(n, 1, 4) // 猜测对象或内部数据
		testCases = append(testCases, testCase{
			input:     input{n},
			guess:     guess{a},
			innerData: a,
		})
	}

	queryChecker := func(caseNum int, tc testCase) func(req) resp {
		n := tc.n
		//a := append([]int(nil), tc.ans...)

		queryCnt, queryLimit := 0, 4*n

		return func(req req) (resp resp) {
			if caseNum == debugCaseNum {
				Print(req, " ")
				defer func() { Println(resp) }()
			}

			queryCnt++
			if queryCnt > queryLimit {
				panic("query limit exceeded")
			}

			// ...

			resp.v = -1
			return
		}
	}

	// 有些题目我们仅需要检查答案的合法性就行了
	ansChecker := func(caseNum int, tc testCase, actualAns guess) bool {

		return true
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
		actualAns := run(tc.input, queryChecker(caseNum, tc))
		if !assert.EqualValues(expectedAns, actualAns, "Wrong Answer %d\nInner Data:\n%s", caseNum, format(tc)) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
		if !assert.Truef(ansChecker(caseNum, tc, actualAns), "Wrong Answer %d\nMy Answer:\n%s\nInner Data:\n%s", caseNum, actualAns, format(tc)) {
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
