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
			`5 7
1 0 3 3
2 2 4 2
3 1 3
2 1 5 1
1 0 2 2
3 0 3
3 3 5`,
			`8
10
4`,
		},
		{
			`2 3
        	            	2 0 2 2 
        	            	1 0 2 2 
        	            	3 0 1 `,
			`2`,
		},
		{
			`2 3 
        	            	1 0 2 1 
        	            	2 0 2 2 
        	            	3 0 2 `,
			`6`,
		},
		{
			`4 4 
        	            	2 3 4 2 
        	            	2 0 2 2 
        	            	1 0 4 2 
        	            	3 1 3 `,
			`4`,
		},
		{
			`2 2
        	            	2 0 2 3
        	            	3 1 2`,
			`3`,
		},
		{
			`3 3
        	            	2 1 3 3
        	            	2 0 3 2
        	            	3 2 3`,
			`5`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 4, run)
}

// 无尽对拍
func Test2(t *testing.T) {
	//return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		q := rg.Int(5, 20)
		rg.NewLine()
		for ; q > 0; q-- {
			op := rg.Int(1, 3)
			l := rg.Int(0, n-1)
			rg.Int(l+1, n)
			if op < 3 {
				rg.Int(0, 3)
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
		a := make([]int, n)
		for ; q > 0; q-- {
			if Fscan(in, &op, &l, &r); op == 1 {
				Fscan(in, &v)
				for i := l; i < r; i++ {
					a[i] = v
				}
			} else if op == 2 {
				Fscan(in, &v)
				for i := l; i < r; i++ {
					a[i] += v
				}
			} else {
				sum := 0
				for _, v := range a[l:r] {
					sum += v
				}
				Fprintln(out, sum)
			}
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
