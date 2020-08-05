package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customTestCases := [][2]string{
		{
			`10 3
1 3 4 91220
1 5 9 48623
2 3 5 39412`,
			`0
0
0
39412
39412
39412
48623
48623
48623
48623`,
		},
		{
			`10 6
1 1 8 4
2 4 9 1
2 3 6 5
1 0 5 3
1 2 2 5
2 6 7 0`,
			`3
4
5
4
3
3
0
0
1
0`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, 0, run)
}

// 无尽对拍
func Test2(t *testing.T) {
	return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		q := rg.Int(1, 10)
		rg.NewLine()
		for ; q > 0; q-- {
			if rand.Intn(2) == 0 {
				rg.Str(1, 1, 'W', 'W')
			} else {
				rg.Str(1, 1, 'B', 'B')
			}
			rg.Int(-10, 10)
			rg.Int(1, 10)
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var q, l, d int
		var s string
		Fscan(in, &q)
		a := make([]int, 50)
		for ; q > 0; q-- {
			Fscan(in, &s, &l, &d)
			l += 13
			r := l + d
			v := 0
			if s[0] == 'B' {
				v = 1
			}
			for i := l; i < r; i++ {
				a[i] = v
			}
			s := 0
			c := 0
			for i, v := range a {
				s += v
				if i > 0 && v == 1 && a[i-1] == 0 {
					c++
				}
			}
			Fprintln(out, c, s)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
