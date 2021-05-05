package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/266/C
// https://codeforces.com/problemset/status/266/problem/C
func TestCF266C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
2
2 1 2
1 1 2
inputCopy
3
3 1
1 3
outputCopy
3
2 2 3
1 1 3
1 1 2
inputCopy
3
2 1
3 2
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF266C)
}

func TestCheckCF266C(t *testing.T) {
	//return
	assert := assert.New(t)

	inputGenerator := func() (string, testutil.OutputChecker) {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2,50)
		rg.NewLine()
		ps := rg.UniquePoints(n-1, 1, n, 1, n)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
		}
		for _, p := range ps {
			x, y := p[0]-1, p[1]-1
			a[x][y] = 1
		}
		return rg.String(), func(myOutput string) (_b bool) {
			in := strings.NewReader(myOutput)

			var m, tp, x, y int
			for Fscan(in, &m); m > 0; m-- {
				Fscan(in, &tp, &x, &y)
				x--
				y--
				if tp == 1 {
					a[x], a[y] = a[y], a[x]
				} else {
					for _, r := range a {
						r[x], r[y] = r[y], r[x]
					}
				}
			}

			for i, r := range a {
				for j, v := range r {
					if v == 1 && i <= j {
						assert.Fail("bad")
						return
					}
				}
			}

			return true
		}
	}

	testutil.CheckRunResultsInf(t, inputGenerator, CF266C)
}
