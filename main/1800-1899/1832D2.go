package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1832D2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, q, k, ans int
	Fscan(in, &n, &q)
	a := make([]int, n)
	s := -int64(n-1) * int64(n-2) / 2
	for i := range a {
		Fscan(in, &a[i])
		s += int64(a[i])
	}
	sort.Ints(a)
	pre := make([]int, n)
	pre[0] = a[0]
	for i := 1; i < n; i++ {
		pre[i] = min(pre[i-1], a[i]-i)
	}
	for ; q > 0; q-- {
		Fscan(in, &k)
		if k < n {
			ans = min(pre[k-1]+k, a[k])
		} else if n == 1 {
			ans = a[0] - k/2
			if k%2 > 0 {
				ans += k
			}
		} else {
			s := s + int64(k)*int64(n-1)
			ans = pre[n-2] + k
			k -= n - 1
			if k%2 > 0 {
				s += int64(k)
				ans = min(ans, a[n-1]+k)
				k--
			} else {
				ans = min(ans, a[n-1])
			}
			s -= int64(ans) * int64(n)
			k /= 2
			if int64(k) > s {
				ans -= (k-int(s)-1)/n + 1
			}
		}
		Fprint(out, ans, " ")
	}
}

//func main() { CF1832D2(os.Stdin, os.Stdout) }
