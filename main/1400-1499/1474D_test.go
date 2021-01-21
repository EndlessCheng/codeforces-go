package main

import (
	."fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/contest/1474/problem/D
// https://codeforces.com/problemset/status/1474/problem/D
func TestCF1474D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
1 2 1
3
1 1 2
5
2 2 2 1 3
5
2100 1900 1600 3000 1600
2
2443 2445
outputCopy
YES
YES
YES
YES
NO
inputCopy
1
4
3 2 1 2
outputCopy
NO
inputCopy
1 
4
2 3 4 1
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, -1, CF1474D)
}

func TestCompareCF1474D(t *testing.T) {
	inputGenerator := func() string {
		rg := testutil.NewRandGenerator()
		rg.Int(1, 1)
		n := rg.Int(2, 7)
		rg.NewLine()
		rg.IntSlice(n, 1, n)
		return rg.String()
	}

	runBF := func(in io.Reader, out io.Writer) {
		var T, n int
		o: for Fscan(in, &T); T > 0; T-- {
			Fscan(in, &n)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}
			check := func(a []int) bool {
				b := append([]int(nil), a...)
				for i := 1; i < n; i++ {
					if b[i-1] > b[i] {
						return false
					}
					b[i] -= b[i-1]
				}
				return b[n-1] == 0
			}
			if check(a) {
				Fprintln(out, "YES")
				continue
			}
			for i := 1; i < n; i++ {
				b := append([]int(nil), a...)
				b[i-1], b[i] = b[i], b[i-1]
				if check(b) {
					Fprintln(out, "YES")
					continue o
				}
			}
			Fprintln(out, "NO")
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1474D)
}
