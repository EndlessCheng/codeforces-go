package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func TestCF1439C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 6
10 10 10 6 6 5 5 5 3 1
2 3 50
2 4 10
1 3 10
2 2 36
1 4 7
2 2 17
outputCopy
8
3
6
2
inputCopy
4 2
8 7 7 2
1 4 4
2 1 15
outputCopy
2
inputCopy
2 2
3 2
1 2 9
2 1 17
outputCopy
1
inputCopy
2 2
5 1
1 2 2
2 1 6
outputCopy
1
inputCopy
3 2
4 2 1
1 3 3
2 1 9
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, -1, CF1439C)
}

func TestCompareCF1439C(t *testing.T) {
	return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1,10)
		q := rg.Int(5,5)
		rg.NewLine()
		rg.IntSliceOrdered(n, 1, 5, false, false)
		for i := 1; i <= q; i++ {
			rg.Int(1,2)
			rg.Int(1, n)
			rg.Int(1, 20)
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	runBF := func(in io.Reader, out io.Writer) {
		var n, q int
		Fscan(in, &n, &q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for ; q > 0; q-- {
			var tp, p, v int
			Fscan(in, &tp, &p, &v)
			if tp == 1 {
				for i := 0; i < p; i++ {
					a[i] = max(a[i], v)
				}
			} else {
				cnt := 0
				for _, w := range a[p-1:] {
					if v >= w {
						cnt++
						v -= w
					}
				}
				Fprintln(out, cnt)
			}
		}
	}

	// 对拍
	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1439C)
}
