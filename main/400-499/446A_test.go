package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/446/A
// https://codeforces.com/problemset/status/446/problem/A
func TestCF446A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
7 2 3 1 5 6
outputCopy
5
inputCopy
10
424238336 649760493 681692778 714636916 719885387 804289384 846930887 957747794 596516650 189641422
outputCopy
9
inputCopy
50
804289384 846930887 681692778 714636916 957747794 424238336 719885387 649760493 596516650 189641422 25202363 350490028 783368691 102520060 44897764 967513927 365180541 540383427 304089173 303455737 35005212 521595369 294702568 726956430 336465783 861021531 59961394 89018457 101513930 125898168 131176230 145174068 233665124 278722863 315634023 369133070 468703136 628175012 635723059 653377374 656478043 801979803 859484422 914544920 608413785 756898538 734575199 973594325 149798316 38664371
outputCopy
19
inputCopy
5
7 4 6 2 3
outputCopy
3
inputCopy
4
3 4 1 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF446A)
}

func TestCompareCF446A(t *testing.T) {
	//return
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 9)
		rg.NewLine()
		rg.IntSlice(n, 1, 9)
		return rg.String()
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n <= 2 {
			Fprintln(out, n)
			return
		}
		ans := 0
		f := func(b []int) {
			c := 1
			mx := 1
			for i := 1; i < len(b); i++ {
				if b[i] > b[i-1] {
					c++
					mx = max(mx, c)
				} else {
					c = 1
				}
			}
			ans = max(ans, mx)
		}

		for i := range a {
			if i > 0 {
				b := append([]int(nil), a...)
				b[i] = b[i-1]+1
				f(b)
			}
			if i < n-1 {
				b := append([]int(nil), a...)
				b[i] = b[i+1]-1
				f(b)
			}

		}
		Fprintln(out, ans)
	}

	// 先用 runBF 跑下样例，检查 runBF 是否正确
//	rawText := `
//inputCopy
//6
//7 2 3 1 5 6
//outputCopy
//5
//inputCopy
//10
//424238336 649760493 681692778 714636916 719885387 804289384 846930887 957747794 596516650 189641422
//outputCopy
//9
//inputCopy
//50
//804289384 846930887 681692778 714636916 957747794 424238336 719885387 649760493 596516650 189641422 25202363 350490028 783368691 102520060 44897764 967513927 365180541 540383427 304089173 303455737 35005212 521595369 294702568 726956430 336465783 861021531 59961394 89018457 101513930 125898168 131176230 145174068 233665124 278722863 315634023 369133070 468703136 628175012 635723059 653377374 656478043 801979803 859484422 914544920 608413785 756898538 734575199 973594325 149798316 38664371
//outputCopy
//19`
//	testutil.AssertEqualCase(t, rawText, 2, runBF)
//	return

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF446A)
}
