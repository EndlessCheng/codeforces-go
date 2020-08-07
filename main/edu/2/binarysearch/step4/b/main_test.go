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
			`4 3
1 2 1
2 3 0
2 4 1`,
			`2
1 2 4 `,
		},
		{
			`3 3
1 2 1
2 3 2
1 3 1`,
			`1
1 3 `,
		},
		{
			`10 20
1 3 37
3 4 6
4 5 62
5 8 46
8 9 88
9 10 56
1 2 6
2 6 50
4 7 89
6 10 90
7 9 45
8 10 73
2 10 87
3 8 78
3 9 28
2 4 69
2 9 80
5 9 36
1 5 47
6 8 76`,
			`5
1 3 4 5 9 10 `,
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
		n := rg.Int(1, 10)
		rg.NewLine()
		rg.IntSlice(n, 1, n)
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0

		Fprint(out, ans)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
