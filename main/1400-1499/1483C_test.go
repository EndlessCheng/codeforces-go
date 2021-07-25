package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/bits"
	"testing"
)

// https://codeforces.com/problemset/problem/1483/C
// https://codeforces.com/problemset/status/1483/problem/C
func TestCF1483C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 5 4
1 5 3 2 4
outputCopy
15
inputCopy
5
1 4 3 2 5
-3 4 -10 2 7
outputCopy
10
inputCopy
2
2 1
-2 -3
outputCopy
-3
inputCopy
10
4 7 3 2 5 1 9 10 6 8
-4 40 -46 -8 -16 4 -10 41 12 3
outputCopy
96
inputCopy
3
1 3 2
-1 1 -1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, -1, CF1483C)
}

func TestCompareCF1483C(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		rg.NewLine()
		rg.Permutation(1, n)
		rg.IntSlice(n, -5, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		h := make([]int, n)
		for i := range h {
			Fscan(in, &h[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		if n == 1 {
			Fprintln(out, b[0])
			return
		}

		calc := func(sub uint) (res int) {
			pre := 0
			for ; sub > 0; sub &= sub - 1 {
				p := bits.TrailingZeros(sub)
				miI := pre
				for i := pre; i <= p; i++ {
					if h[i] < h[miI] {
						miI = i
					}
				}
				res+=b[miI]
				pre = p+1
			}
			miI := pre
			for i := pre; i < n; i++ {
				if h[i] < h[miI] {
					miI = i
				}
			}
			res+=b[miI]
			return
		}
		ans := int(-1e18)
		for sub := uint(0); sub < 1<<(n-1); sub++ {
			res := calc(sub)
			if res > ans {
				ans = res
			}
		}
		Fprintln(out, ans)
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//5
//1 2 3 5 4
//1 5 3 2 4
//outputCopy
//15
//inputCopy
//5
//1 4 3 2 5
//-3 4 -10 2 7
//outputCopy
//10
//inputCopy
//2
//2 1
//-2 -3
//outputCopy
//-3
//inputCopy
//10
//4 7 3 2 5 1 9 10 6 8
//-4 40 -46 -8 -16 4 -10 41 12 3
//outputCopy
//96`
//	testutil.AssertEqualCase(t, rawText, 0, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1483C)
}
