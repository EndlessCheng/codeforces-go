package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://codeforces.com/contest/1641/problem/B
// https://codeforces.com/problemset/status/1641/problem/B
func TestCF1641B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2
5 7
2
5 5
6
1 3 1 2 2 3
6
3 2 1 1 2 3
outputCopy
-1
0
1
2
4
1 3
5 3
5 3
10 3
2
8 6 
5
0 3
8 3
5 3 
6 2 
7 1
4
2 6 6 2
inputCopy
1
6
2 1 3 2 3 1
outputCopy
3
4 1
5 3
9 1
3
6 4 2
inputCopy
1
6
1 3 3 2 1 2
outputCopy

inputCopy
1
2
3 3
outputCopy
`
	testutil.AssertEqualCase(t, rawText, -1, CF1641B)
}

func TestCheckCF1641B(t *testing.T) {
	//return
	assert := assert.New(t)
	_ = assert

	//testutil.DebugTLE = 0

	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 9)
		rg.NewLine()
		a := rg.IntSlice(n, 1, 9)
		return rg.String(), func(myOutput string) (_b bool) {
			in := strings.NewReader(myOutput)

			var nop ,sz int
			Fscan(in, &nop)

			cnt := map[int]int{}
			for _, v := range a {
				cnt[v]++
			}
			for _, c := range cnt {
				if c&1 > 0 {
					if !assert.EqualValues(-1, nop) {
						return
					}
				}
			}

			for i := 0; i < nop; i++ {
				var p, c int
				Fscan(in, &p, &c)
				a = append(a[:p], append([]int{c, c}, a[p:]...)...)
			}

			Fscan(in, &sz)
			for i := 0; i < sz; i++ {
				var k int
				Fscan(in, &k)
				b := a[:k]
				a=a[k:]
				if !assert.EqualValues(0, len(b)&1) {
					return
				}
				for i := 0; i < len(b)/2; i++ {
					if !assert.EqualValues(b[i], b[i+len(b)/2]) {
						return
					}
				}
			}

			return true
		}
	}

	target := 0
	testutil.CheckRunResultsInfWithTarget(t, inputGenerator, target, CF1641B)
}
