package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type ST91 [][]int

func newST91(a []int) ST91 {
	n := len(a)
	sz := bits.Len(uint(n))
	st := make(ST91, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
	}
	for j := 1; 1<<j <= n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = gcd91(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

func (st ST91) query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return gcd91(st[l][k], st[r-1<<k][k])
}

func CF891A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, c1 int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		if a[i] == 1 {
			c1++
		}
	}
	if c1 > 0 {
		Fprint(out, n-c1)
		return
	}

	st := newST91(a)
	ans := int(1e9)
	for l := range a {
		sz := sort.Search(n-l, func(sz int) bool { return st.query(l, l+sz+1) == 1 })
		// 注意这里的 sz 实际上要 +1，但是由于只需要取一个相对最小值，+1 可以不写
		if sz < n-l && sz < ans {
			ans = sz
		}
	}
	if ans == 1e9 {
		Fprint(out, -1)
	} else {
		Fprint(out, ans+n-1)
	}
}

//func main() { CF891A(os.Stdin, os.Stdout) }

func gcd91(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
