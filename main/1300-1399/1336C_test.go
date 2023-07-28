package main

import (
	"bufio"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1336/C
// https://codeforces.com/problemset/status/1336/problem/C
func TestCF1336C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
abab
ba
outputCopy
12
inputCopy
defineintlonglong
signedmain
outputCopy
0
inputCopy
rotator
rotator
outputCopy
4
inputCopy
cacdcdbbbb
bdcaccdbbb
outputCopy
24
inputCopy
baabbaaaaa
ababa
outputCopy
120
inputCopy
cbaabccadacadbaacbadddcdabcacdbbabccbbcbbcbbaadcabadcbdcadddccbbbbdacdbcbaddacaadbcddadbabbdbbacabdd
cccdacdabbacbbcacdca
outputCopy
773806867
inputCopy
acacdddadaddcbbbbdacb
babbbcadddacacdaddbdc
outputCopy
64`
	testutil.AssertEqualCase(t, rawText, 0, CF1336C)
}

func CF1336C_save(in io.Reader, out io.Writer) {
	const mod = 998244353
	var s, t []byte
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)

	// s 前缀对应 t 的后缀
	d0 := make([][]int, m)
	for i := range d0 {
		d0[i] = make([]int, m)
		for j := range d0[i] {
			d0[i][j] = -1
		}
	}
	var f1 func(int, int) int
	f1 = func(l, r int) (res int) {
		if l > r {
			return 1
		}
		p := &d0[l][r]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		i := r - l
		if s[i] == t[l] {
			res = f1(l+1, r)
		}
		if s[i] == t[r] {
			res = (res + f1(l, r-1)) % mod
		}
		return
	}

	d1 := make([][]int, n)
	for i := range d1 {
		d1[i] = make([]int, m)
		for j := range d1[i] {
			d1[i][j] = -1
		}
	}
	var f2 func(int, int) int
	f2 = func(i, j int) (res int) {
		if j < 0 {
			return n + 1 - i
		}
		if i == n {
			return
		}
		p := &d1[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		res = f2(i+1, j)
		if s[i] == t[j] {
			res = (res + f2(i+1, j-1)) % mod
		}
		return
	}

	ans := int64(0)
	for tl := 0; tl < m; tl++ {
		res := f1(tl, m-1)
		res2 := -1
		//for j := tl - 1; j >= 0 ; j-- {
		//	if s[m-1-j] != t[j] {
		//		res2 = f2(m-1-j, j)
		//		break
		//	}
		//}
		//if res2 < 0 {
		//	res2 = f2(m, -1)
		//}
		if tl > 0 && s[m-tl] == t[tl-1] {
			// skip one
			res2 = f2(m-tl+1, tl-1)
		} else {
			res2 = f2(m-tl, tl-1)
		}
		Println(tl, res, res2)
		//j := tl - 1
		//for ; j >= 0 && s[m-1-j] == t[j]; j-- {
		//}
		//if tl > 0 && j < 0 {
		//	res2--
		//}
		//Println(tl, res, res2)
		//if res2 > 0 {
		ans = (ans + int64(res)*int64(res2)) % mod
		//}
	}
	// 整个 rev(t) 作为 s 的子序列
	p2 := 1
	for i, j := 0, m-1; ; {
		if j >= 0 && s[i] == t[j] {
			p2 = p2 * 2 % mod
			i++
			j--
		} else {
			res2 := f2(i, m-1)
			Println(p2, res2)
			ans = (ans + int64(p2)*int64(res2)) % mod
			break
		}
	}
	Fprint(out, ans)
}
