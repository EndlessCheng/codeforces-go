package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1328F(_r io.Reader, out io.Writer) {
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n, k int64
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	s := make([]int64, n+1)
	for i, v := range a {
		s[i+1] = s[i] + int64(v)
	}
	ans := int64(1e18)
	for i := int64(0); i < n; {
		st := i
		for i++; i < n && a[i] == a[st]; i++ {
		}
		d := k - i + st
		if d <= 0 {
			ans = 0
			break
		}
		v := int64(a[st])
		s1, s2 := (v-1)*st-s[st], s[n]-s[i]-(v+1)*(n-i)
		if i >= k {
			ans = min(ans, s1+d)
		}
		if n-st >= k {
			ans = min(ans, s2+d)
		}
		ans = min(ans, s1+s2+d)
	}
	Fprint(out, ans)
}

//func main() { CF1328F(os.Stdin, os.Stdout) }
