package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1082/E
// https://codeforces.com/problemset/status/1082/problem/E
func TestCF1082E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 9
9 9 9 9 9 9
outputCopy
6
inputCopy
3 2
6 2 6
outputCopy
2
inputCopy
10 1
0 1 1 0 0 1 0 1 1 1
outputCopy
8
inputCopy
29 3
1 4 3 4 3 3 4 1 1 3 4 1 3 3 4 1 3 1 2 3 1 4 2 2 3 4 2 4 2
outputCopy
12
inputCopy
4 3
4 3 4 4
outputCopy
3
inputCopy
6 5
1 1 5 1 4 1
outputCopy
4
inputCopy
7 5
1 1 4 5 1 1 2
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, -1, CF1082E)
}

func TestCompareCF1082E(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 7)
		rg.Int(1, 5)
		rg.NewLine()
		rg.IntSlice(n, 1, 5)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, tar int
		Fscan(in, &n, &tar)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		mx := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j <= n; j++ {
				for add := -10; add < 10; add++ {
					cnt := 0
					for _, v := range a[:i] {
						if v == tar {
							cnt++
						}
					}
					for k := i; k < j; k++ {
						if a[k]+add == tar {
							cnt++
						}
					}
					for _, v := range a[j:] {
						if v == tar {
							cnt++
						}
					}
					if cnt > mx {
						mx = cnt
					}
				}
			}
		}
		Fprintln(out, mx)
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//6 9
//9 9 9 9 9 9
//outputCopy
//6
//inputCopy
//3 2
//6 2 6
//outputCopy
//2
//inputCopy
//10 1
//0 1 1 0 0 1 0 1 1 1
//outputCopy
//8
//inputCopy
//29 3
//1 4 3 4 3 3 4 1 1 3 4 1 3 3 4 1 3 1 2 3 1 4 2 2 3 4 2 4 2
//outputCopy
//12`
//	testutil.AssertEqualCase(t, rawText, 0, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1082E)
}
