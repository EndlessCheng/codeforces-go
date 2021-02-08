package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1481/C
// https://codeforces.com/problemset/status/1481/problem/C
func TestCF1481C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 1
1
1
1
5 2
1 2 2 1 1
1 2 2 1 1
1 2
3 3
2 2 2
2 2 2
2 3 2
10 5
7 3 2 1 7 9 4 2 7 9
9 9 2 1 4 9 4 2 3 9
9 9 7 4 3
5 2
1 2 2 1 1
1 2 2 1 1
3 3
6 4
3 4 2 4 1 2
2 3 1 3 1 1
2 2 3 4
outputCopy
YES
1
YES
2 2
YES
1 1 1
YES
2 1 9 5 9
NO
NO
inputCopy
1
4 3
4 2 3 2
4 4 2 4
3 3 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1481C)
}

func TestCheckCF1481C(t *testing.T) {
	//return
	assert := assert.New(t)

	inputGenerator := func() (string, testutil.OutputChecker) {
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(1, 10)
		m := rg.Int(1, 10)
		rg.NewLine()
		a := rg.IntSlice(n, 1, n)
		b := rg.IntSlice(n, 1, n)
		c := rg.IntSlice(m, 1, n)
		return rg.String(), func(output string) (_b bool) {
			in := strings.NewReader(output)
			var s string
			Fscan(in, &s)
			if s == "NO" {
				//Println(a)
				//Println(b)
				//Println(cc)
				//Println()
				return true
			}
			for i := 0; i < m; i++ {
				var id int
				Fscan(in, &id)
				a[id-1] = c[i]
			}
			return assert.EqualValues(b, a)
		}
	}

	testutil.CheckRunResultsInf(t, inputGenerator, CF1481C)
}
