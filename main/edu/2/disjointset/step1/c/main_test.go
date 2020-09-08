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
			`3 6
add 1 100
join 1 3
add 1 50
get 1
get 2
get 3`,
			`150
0
50`,
		},
		{
			`3 4
a 1 110
j 3 1
j 1 3
g 3`,
			`0`,
		},
		{
			`8 4
j 5 7
a 7 107
g 5
g 5`,
			`107
107`,
		},
		{
			`2 4
j 2 1
a 2 103
a 2 108
g 2`,
			`211`,
		},
	}
	testutil.AssertEqualStringCase(t, customTestCases, -1, run)
}

// 无尽对拍
func Test2(t *testing.T) {
	//return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		q := rg.Int(10, 10)
		rg.NewLine()
		for ; q > 0; q-- {
			op := rand.Intn(3)
			if op == 0 {
				rg.Byte('j')
				rg.Space()
				rg.Int(1, n)
				rg.Int(1, n)
			} else if op == 1 {
				rg.Byte('a')
				rg.Space()
				rg.Int(1, n)
				rg.Int(100, 110)
			} else {
				rg.Byte('g')
				rg.Space()
				rg.Int(1, n)
			}
			rg.NewLine()
		}
		//Println(rg.String())
		return rg.String()
	}

	var fa, exp []int
	initFa := func(n int) {
		fa = make([]int, n)
		for i := range fa {
			fa[i] = i
		}
		exp = make([]int, n)
	}
	var f func(int) int
	f = func(x int) int {
		if fa[x] != x {
			fa[x] = f(fa[x])
		}
		return fa[x]
	}
	merge := func(from, to int) { fa[f(from)] = f(to) }

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n, q, x, y int
		var s []byte
		Fscan(in, &n, &q)
		initFa(n + 1)
		for ; q > 0; q-- {
			Fscan(in, &s, &x)
			if s[0] == 'j' {
				Fscan(in, &y)
				merge(x, y)
			} else if s[0] == 'a' {
				Fscan(in, &y)
				for i := 1; i <= n; i++ {
					if f(i) == f(x) {
						exp[i] += y
					}
				}
			} else {
				Fprintln(out, exp[x])
			}
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
