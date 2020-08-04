package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func Test(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customTestCases := [][2]string{
		{
			`5 4
1 2 3 4 5
1 2 3 1
2 1 3
1 2 3 -1
2 1 5`,
			`19
55`,
		},
		{
			`8 9
        	            	2 9 -8 -3 -4 -10 9 4
        	            	2 2 7
        	            	1 5 7 1
        	            	1 5 7 -7
        	            	2 5 5
        	            	2 3 3
        	            	1 8 8 -8
        	            	1 1 8 6
        	            	2 2 7
        	            	2 6 7`,
			"-28\n-10\n-8\n8\n8",
		},
		{
			`8 4
        	            	0 1 1 1 0 1 1 1
        	            	2 2 7
        	            	1 5 7 0
        	            	1 5 7 1
        	            	2 5 5`,
			"17\n1",
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}

// 无尽对拍
func Test2(t *testing.T) {
	//return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		q := rg.Int(1, 20)
		rg.NewLine()
		rg.IntSlice(n, -100, 100)
		for ; q > 0; q-- {
			op := rg.Int(1, 2)
			l := rg.Int(1, n)
			rg.Int(l, n)
			if op == 1 {
				rg.Int(-100, 100)
			}
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n, q, op, l, r, v int
		Fscan(in, &n, &q)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		for ; q > 0; q-- {
			if Fscan(in, &op, &l, &r); op == 1 {
				Fscan(in, &v)
				for i := l; i <= r; i++ {
					a[i] += v
				}
			} else {
				sum := 0
				for i, v := range a[l : r+1] {
					sum += v * (i + 1)
				}
				Fprintln(out, sum)
			}
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
