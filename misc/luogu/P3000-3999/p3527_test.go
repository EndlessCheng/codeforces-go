package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://www.luogu.com.cn/problem/P3527
func Test_p3527(t *testing.T) {
	customTestCases := [][2]string{
		{
			`3 5
1 3 2 1 3
10 5 7
3
4 2 4
1 3 1
3 5 2`,
			`3
NIE
1`,
		},
	}

	tarCase := 0 // -1
	testutil.AssertEqualStringCase(t, customTestCases, tarCase, p3527)
}

func TestCompare_p3527(t *testing.T) {
	return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		m := rg.Int(1, 9)
		rg.NewLine()
		rg.IntSlice(m, 1, n)
		rg.IntSlice(n, 1, 9)
		q := rg.Int(1, 9)
		rg.NewLine()
		for i := 0; i < q; i++ {
			rg.Int(1, m)
			rg.Int(1, m)
			rg.Int(1, 9)
			rg.NewLine()
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, m, q int
		Fscan(in, &n, &m)
		c := make([]int, m)
		for i := range c {
			Fscan(in, &c[i])
			c[i]--
		}
		need := make([]int, n)
		for i := range need {
			Fscan(in, &need[i])
		}
		ans := make([]int, n)
		Fscan(in, &q)
		for qi := 1; qi <= q; qi++ {
			var l, r, v int
			Fscan(in, &l, &r, &v)
			l--
			if l < r {
				for _, c := range c[l:r] {
					need[c] -= v
				}
			} else {
				for _, c := range c[:r] {
					need[c] -= v
				}
				for _, c := range c[l:] {
					need[c] -= v
				}
			}
			for i, v := range need {
				if v <= 0 && ans[i] == 0 {
					ans[i] = qi
				}
			}
		}
		for _, v := range ans {
			if v == 0 {
				Fprintln(out, "NIE")
			} else {
				Fprintln(out, v)
			}
		}
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	customTestCases := [][2]string{
//		{
//			`3 5
//1 3 2 1 3
//10 5 7
//3
//4 2 4
//1 3 1
//3 5 2`,
//			`3
//NIE
//1`,
//		},
//	}
//	testutil.AssertEqualStringCase(t, customTestCases, 0, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, p3527)
}
