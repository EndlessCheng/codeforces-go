package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/342/E
// https://codeforces.com/problemset/status/342/problem/E
func TestCF342E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
1 2
2 3
2 4
4 5
2 1
2 5
1 2
2 5
outputCopy
0
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF342E)
}

func TestRECF342E(_t *testing.T) {
	return
	testutil.DebugTLE = 0

	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 9)
		m := rg.Int(1,9)
		rg.NewLine()
		rg.TreeEdges(n,1)
		for i := 0; i < m; i++ {
			rg.Int(1,2)
			rg.Int(1,n)
			rg.NewLine()
		}
		return rg.String(), func(myOutput string) bool { return true }
	}
	testutil.CheckRunResultsInfWithTarget(_t, inputGenerator, 0, CF342E)
}
