package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//func init() { rand.Seed(time.Now().UnixNano()) }

func Test_run23(t *testing.T) { testRun23(t, 0) }

func testRun23(t *testing.T, debugCaseNum int) {
	assert := assert.New(t)
/*

   ..#.
   #...
   #.#.
   ....
*/
	type testCase struct {
		input23
		guess23
		innerData []string
	}
	format := func(tc testCase) string { return strings.Join(tc.innerData, "\n") }

	testCases := []testCase{}
	for tc := 0; tc < 1e5; tc++ {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 6) // 输入
		a := make([][]byte, n)
		for i := range a {
			a[i] = bytes.Repeat([]byte{'#'}, n)
		}

		x, y := 0, 0
		for x < n-1 || y < n-1 {
			a[x][y] = '.'
			if x == n-1 {
				y++
			} else if y == n-1 {
				x++
			} else if rg.Int(0, 1) == 0 {
				x++
			} else {
				y++
			}
		}
		a[n-1][n-1] = '.'

		clear := rg.Int(0, n*n)
		for i := 0; i < clear; i++ {
			a[rg.Int(0, n-1)][rg.Int(0, n-1)] = '.'
		}

		b := make([]string, n)
		for i, r := range a {
			b[i] = string(r)
		}
		testCases = append(testCases, testCase{
			input23:   input23{n},
			guess23:   guess23{},
			innerData: b,
		})
	}

	queryChecker := func(caseNum int, tc testCase) func(req23) resp23 {
		n := tc.n
		queryCnt, queryLimit := 0, 4*n

		return func(req req23) (resp resp23) {
			if caseNum == debugCaseNum {
				Print(req, " ")
				defer func() { Println(resp) }()
			}

			queryCnt++
			if queryCnt > queryLimit {
				panic("query limit exceeded")
			}

			var f func(x, y int) bool
			f = func(x, y int) bool {
				if x == n || y == n || tc.innerData[x][y] == '#' {
					return false
				}
				if x == req.tx-1 && y == req.ty-1 {
					return true
				}
				return f(x+1, y) || f(x, y+1)
			}
			if f(req.sx-1, req.sy-1) {
				resp.s = "YES"
			} else {
				resp.s = "NO"
			}
			return
		}
	}

	ansChecker := func(caseNum int, tc testCase, actualAns guess23) bool {
		x, y := 0, 0
		for _, b := range actualAns.ans {
			if b == 'D' {
				x++
			} else {
				y++
			}
			if tc.innerData[x][y] == '#' {
				return false
			}
		}
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
		actualAns := CF1023E(tc.input23, queryChecker(caseNum, tc))
		if !assert.Truef(ansChecker(caseNum, tc, actualAns), "Wrong Answer %d\nMy Answer:\n%s\nInner Data:\n%s", caseNum, actualAns, format(tc)) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testRun23(t, 0)
	}
}
