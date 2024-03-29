// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://codeforces.com/problemset/problem/375/D
// https://codeforces.com/problemset/status/375/problem/D
func Test_cf375D(t *testing.T) {
	testCases := [][2]string{
		{
			`8 5
1 2 2 3 3 2 3 3
1 2
1 5
2 3
2 4
5 6
5 7
5 8
1 2
1 3
1 4
2 3
5 3`,
			`2
2
1
0
1`,
		},
		{
			`4 1
1 2 3 4
1 2
2 3
3 4
1 1`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf375D)
}

func TestCheck_cf375D(_t *testing.T) {
	return
	assert := assert.New(_t)
	_ = assert
	testutil.DebugTLE = 0
	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 5)
		m := rg.Int(1,1)
		rg.NewLine()
		rg.IntSlice(n, 1, 5)
		rg.TreeEdges(n, 1)
		for i := 0; i < m; i++ {
			rg.Int(1, n)
			rg.Int(1, 5)
			rg.NewLine()
		}
		return rg.String(), func(myOutput string) (_b bool) {
			return true
		}
	}

	testutil.CheckRunResultsInfWithTarget(_t, inputGenerator, 0, cf375D)
}
