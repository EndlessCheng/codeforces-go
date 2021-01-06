package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math"
	"testing"
)

func TestCF1470B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
6 8 4 2
1
0
6
12 3 20 5 80 1
1
1
outputCopy
2
3
inputCopy
1
40
1 3 5 6 2 10 7 10 4 2 10 10 6 1 1 8 1 3 7 2 9 10 4 2 4 1 9 1 1 4 1 7 8 1 7 9 3 10 7 1
40
1
7
0
5
1
9
9
2
0
5
7
7
3
10
4
6
0
9
0
8
5
5
3
3
6
5
1
10
8
8
3
1
1
4
7
1
5
8
4
1
outputCopy
31
31
17
31
31
31
31
31
17
31
31
31
31
31
31
31
17
31
17
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
31
inputCopy
1
3
956629 511546 328256
3
0
0
348237548065718627
outputCopy
1
1
1`

	testutil.AssertEqualCase(t, rawText, -1, CF1470B)
}

func TestCompare(t *testing.T) {
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.Int(1, 1)
		rg.NewLine()
		n := rg.Int(1, 6)
		rg.NewLine()
		rg.IntSlice(n, 1, 16)
		q := rg.Int(3, 3)
		rg.NewLine()
		for i := 0; i < q; i++ {
			rg.Int(i, i)
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	sqCheck := func(a int) bool { r := int(math.Round(math.Sqrt(float64(a)))); return r*r == a }
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		solve := func(Case int) {
			var n int
			Fscan(in, &n)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}
			f := func(a []int) int {
				mx := 0
				for _, v := range a {
					c := 0
					for _, w := range a {
						g := gcd(v, w)
						if sqCheck(v * w / g / g) {
							c++
						}
					}
					mx = max(mx, c)
				}
				return mx
			}
			var q int
			Fscan(in, &q)
			ans := make([]int, q)
			for i := 0; i < q; i++ {
				ans[i] = f(a)
				b := make([]int, n)
				for i, v := range a {
					mul := 1
					for _, w := range a {
						g := gcd(v, w)
						if sqCheck(v * w / g / g) {
							mul *= w
							if mul == 0 {
								panic(-1)
							}
						}
					}
					b[i] = mul
				}
				a = b
			}
			for i := 0; i < q; i++ {
				var w int
				Fscan(in, &w)
				Fprintln(out, ans[w])
			}
		}

		var t int
		Fscan(in, &t)
		for Case := 1; Case <= t; Case++ {
			solve(Case)
		}
	}

	//testutil.AssertEqualCase(t, rawText, 0, runBF)
	//return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1470B)
}
