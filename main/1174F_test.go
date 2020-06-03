package main

import (
	. "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCF1174F(t *testing.T) {
	testCF1174F(t, 0)
}

func testCF1174F(t *testing.T, debugCaseNum int) {
	type testCase struct {
		input1174
		guess1174
	}
	testCases := []testCase{
		{
			input1174: input1174{5, [][]int{1: {2, 3}, 2: {1}, 3: {1, 4, 5}, 4: {3}, 5: {3}}},
			guess1174: guess1174{5},
		},
	}

	const (
		queryLimit    = 36
		minQueryValue = 1
	)
	checkQuery := func(caseNum int, tc testCase) func(qIn1174) qOut1174 {
		n, g := tc.n, tc.g
		x := tc.ans
		const mx = 20
		pa := make([][mx]int, n+1)
		dep := make([]int, n+1)
		var dfs func(v, p, d int)
		dfs = func(v, p, d int) {
			pa[v][0] = p
			dep[v] = d
			for _, w := range g[v] {
				if w != p {
					dfs(w, v, d+1)
				}
			}
		}
		dfs(1, 0, 0)
		for k := 0; k+1 < mx; k++ {
			for v := range pa {
				if p := pa[v][k]; p != -1 {
					pa[v][k+1] = pa[p][k]
				} else {
					pa[v][k+1] = -1
				}
			}
		}
		uptoDep := func(v, d int) int {
			for k := 0; k < mx; k++ {
				if (dep[v]-d)>>k&1 == 1 {
					v = pa[v][k]
				}
			}
			return v
		}
		_lca := func(v, w int) int {
			if dep[v] > dep[w] {
				v, w = w, v
			}
			w = uptoDep(w, dep[v])
			if v == w {
				return v
			}
			for k := mx - 1; k >= 0; k-- {
				if pa[v][k] != pa[w][k] {
					v, w = pa[v][k], pa[w][k]
				}
			}
			return pa[v][0]
		}
		_d := func(v, w int) int { return dep[v] + dep[w] - dep[_lca(v, w)]<<1 }
		down := func(u, v int) int {
			if dep[u] >= dep[v] {
				return -1
			}
			v = uptoDep(v, dep[u]+1)
			if pa[v][0] == u {
				return v
			}
			return -1
		}
		_queryCnt := 0
		return func(qi qIn1174) (resp qOut1174) {
			tp, v := qi.tp, qi.v
			if caseNum == debugCaseNum {
				Println(qi)
			}
			_queryCnt++
			if _queryCnt > queryLimit {
				panic("query limit exceeded")
			}
			if v < minQueryValue || v > n {
				panic("invalid query arguments")
			}
			if tp == "d" {
				resp.v = _d(v, x)
			} else if tp == "s" {
				w := down(v, x)
				if w == -1 {
					panic("invalid query arguments")
				}
				resp.v = w
			} else {
				panic("invalid query type")
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
		expectedAns := tc.guess1174
		actualAns := CF1174F(tc.input1174, checkQuery(caseNum, tc))
		if !assert.EqualValues(t, expectedAns, actualAns, "WA %d", caseNum) {
			failedCount++
			if failedCount > failedCountLimit {
				t.Fatal("too many wrong cases, terminated")
			}
		}
	}

	if debugCaseNum != 0 && failedCount == 0 {
		testCF1174F(t, 0)
	}
}
