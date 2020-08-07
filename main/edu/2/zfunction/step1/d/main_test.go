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
			`4
zaza
az
abacaba
ab
aaaaa
a
aaa
b`,
			`6
14
0
6`,
		},
		{
			`8
abbbbabbbabb
b
ababbaa
ab
az
az
z
a
a
aa
ababaaba
aba
abcdefgh
fgh
abacabadaba
ba`,
			`3
14
2
1
1
19
30
24`,
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
