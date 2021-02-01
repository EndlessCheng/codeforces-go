package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/940/F
// https://codeforces.com/problemset/status/940/problem/F
func TestCF940F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 4
1 2 3 1 1 2 2 2 9 9
1 1 1
1 2 8
2 7 1
1 2 8
outputCopy
2
3
2`
	testutil.AssertEqualCase(t, rawText, 0, CF940F)
}

func TestCompareCF940F(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 100)
		q := rg.Int(1, 100)
		rg.NewLine()
		rg.IntSlice(n, 1e7, 1e9)
		for i := 0; i < q; i++ {
			if rg.Int(1,2) == 1 {
				l := rg.Int(1, n)
				rg.Int(l, n)
			} else {
				rg.Int(1, n)
				rg.Int(1e7, 1e9)
			}
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	runBF := func(in io.Reader, out io.Writer) {
		//return
		//var n, q int
		//Fscan(in, &n, &q)
		//a := make([]int, n+1)
		//for i := 1; i <= n; i++ {
		//	Fscan(in, &a[i])
		//}
		//for i := 0; i < q; i++ {
		//	var op string
		//	Fscan(in, &op)
		//	if op == "Q" {
		//		var l, r int
		//		Fscan(in, &l, &r)
		//		c := map[int]bool{}
		//		for _, v := range a[l : r+1] {
		//			c[v] = true
		//		}
		//		Fprintln(out, len(c))
		//	} else {
		//		var pos, val int
		//		Fscan(in, &pos, &val)
		//		a[pos] = val
		//	}
		//}
	}

	//dir, _ := filepath.Abs(".")
	//testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, runBF)
	//return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF940F)
	return

	// for hacking, write wrong codes in runBF
	testutil.AssertEqualRunResultsInf(t, inputGenerator, CF940F, runBF)
}

