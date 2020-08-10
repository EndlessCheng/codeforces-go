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
			`8
1 0
2 0
2 1
2 2
3 0
3 3
3 6
3 4`,
			`0
0
0
1
0
0
1
3`,
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
