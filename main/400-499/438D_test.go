package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/438/D
// https://codeforces.com/problemset/status/438/problem/D
func TestCF438D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 5
1 2 3 4 5
2 3 5 4
3 3 5
1 2 5
2 1 3 3
1 1 3
outputCopy
8
5
inputCopy
10 10
6 9 6 7 6 1 10 10 9 5
1 3 9
2 7 10 9
2 5 10 8
1 4 7
3 3 7
2 7 9 9
1 2 4
1 6 6
1 5 9
3 1 10
outputCopy
49
15
23
1
9
inputCopy
5 2
3 2 3 5 3
2 5 5 3
1 1 5
outputCopy
13`
	testutil.AssertEqualCase(t, rawText, -1, CF438D)
}

func TestCompareCF438D(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n, q := rg.Int(1, 5), rg.Int(1, 5)
		rg.NewLine()
		rg.IntSlice(n, 1,5)
		for i := 0; i < q; i++ {
			op := rg.Int(1, 3)
			if op == 1 {
				rg.Int(rg.Int(1, n), n)
			} else if op == 2 {
				rg.Int(rg.Int(1, n), n)
				rg.Int(1,5)
			} else {
				rg.Int(1, n)
				rg.Int(1,5)
			}
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n,q int
		Fscan(in, &n,&q)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		for i := 0; i < q; i++ {
			var op,l,r, k,x int
			Fscan(in, &op)
			if op == 1 {
				Fscan(in, &l, &r)
				sum := 0
				for _, v := range a[l-1 : r] {
					sum += v
				}
				Fprintln(out, sum)
			} else if op == 2{
				Fscan(in, &l, &r, &x)
				for i := l - 1; i < r; i++ {
					a[i]%=x
				}
			} else {
				Fscan(in, &k, &x)
				a[k-1] = x
			}
		}

	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//5 5
//1 2 3 4 5
//2 3 5 4
//3 3 5
//1 2 5
//2 1 3 3
//1 1 3
//outputCopy
//8
//5
//inputCopy
//10 10
//6 9 6 7 6 1 10 10 9 5
//1 3 9
//2 7 10 9
//2 5 10 8
//1 4 7
//3 3 7
//2 7 9 9
//1 2 4
//1 6 6
//1 5 9
//3 1 10
//outputCopy
//49
//15
//23
//1
//9`
//	testutil.AssertEqualCase(t, rawText, 0, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF438D)
}
