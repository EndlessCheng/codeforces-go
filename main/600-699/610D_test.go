package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strconv"
	"testing"
)

// https://codeforces.com/problemset/problem/610/D
// https://codeforces.com/problemset/status/610/problem/D
func TestCF610D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
0 1 2 1
1 4 1 2
0 3 2 3
outputCopy
8
inputCopy
4
-2 -1 2 -1
2 1 -2 1
-1 -2 -1 2
1 2 1 -2
outputCopy
16
inputCopy
2
4 4 4 0
4 3 4 0
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, -1, CF610D)
}

func TestCompareCF610D(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 7)
		rg.NewLine()
		h := rg.IntOnly(0, n)
		v := n - h
		for i := 0; i < h; i++ {
			rg.Int(0, 4)
			y := rg.Int(0, 4)
			rg.Int(0, 4)
			rg.Bytes(strconv.Itoa(y))
			rg.NewLine()
		}
		for i := 0; i < v; i++ {
			x := rg.Int(0, 4)
			rg.Int(0, 4)
			rg.Bytes(strconv.Itoa(x) + " ")
			rg.Int(0, 4)
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		mp := map[[2]int]bool{}
		for i := 0; i < n; i++ {
			var x1,y1,x2,y2 int
			Fscan(in, &x1, &y1, &x2, &y2)
			if y1 == y2 {
				if x1 > x2 {
					x1, x2 = x2, x1
				}
				for x := x1; x <= x2; x++ {
					mp[[2]int{x,y1}]=true
				}
			} else {
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				for y := y1; y <= y2; y++ {
					mp[[2]int{x1,y}]=true
				}
			}
		}
		Fprintln(out, len(mp))
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF610D)
}
