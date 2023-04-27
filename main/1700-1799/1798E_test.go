package main

import (
	"bufio"
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1798/E
// https://codeforces.com/problemset/status/1798/problem/E
func TestCF1798E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 2 1 7
7
3 1 3 1 2 1 1
4
2 7 1 1
outputCopy
0 1 1 
0 1 1 0 1 1 
1 1 1 
inputCopy
1
19
3 4 1 2 1 7 7 3 1 3 1 2 1 1 4 2 7 1 1
outputCopy
0 0 1 1 1 1 1 1 1 0 1 0 1 0 2 1 1 1 
inputCopy
1
4
2 1 2 1
outputCopy
1 1 2
inputCopy
1
5
2 5 1 2 4
outputCopy
2 2 1 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1798E)
}

func TestCompare(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One()
		n := rg.Int(2, 9)
		rg.NewLine()
		rg.IntSlice(n, 1, 9)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, CF1798EAC, CF1798E)
}

func CF1798EAC(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var scan = fmt.Fscan
	var println = func(iface interface{}) {
		fmt.Fprintln(io.Writer(out), iface)
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	var t int
	scan(in, &t)
	for ; t > 0; t-- {
		var n int
		scan(in, &n)
		a := make([]int, n)
		for i := range a {
			scan(in, &a[i])
		}
		reachEnds := make([]int, n)
		steps := make([]int, n)
		steps[n-1] = 1
		ans := []int{}
		maxRE := 0
		for i := n - 2; i >= 0; i-- {
			if reachEnds[i+1] == a[i] {
				ans = append(ans, 0)
			} else if reachEnds[i+1] > 0 {
				ans = append(ans, 1)
			} else if steps[i+1] >= a[i] {
				ans = append(ans, 1)
			} else {
				ans = append(ans, 2)
			}
			steps[i] = maxRE + 1
			if next := i + a[i] + 1; next < n {
				steps[i] = max(steps[i], steps[next]+1)
				if reachEnds[next] > 0 {
					reachEnds[i] = reachEnds[next] + 1
					maxRE = max(maxRE, reachEnds[i])
				}
			} else if next == n {
				reachEnds[i] = 1
				maxRE = max(maxRE, reachEnds[i])
			}
		}
		for i := len(ans) - 1; i >= 0; i-- {
			fmt.Fprintf(out, "%d ", ans[i])
		}
		println("")
	}

}
