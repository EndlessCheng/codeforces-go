package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1333C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, l int
	var ans, s int64
	sumP := map[int64]int{0: 0}
	Fscan(in, &n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		s += int64(v)
		if p, ok := sumP[s]; ok && p+1 > l { // 子区间和为 0 => 出现了两个同样的前缀和
			l = p + 1
		}
		ans += int64(i - l)
		sumP[s] = i
	}
	Fprint(_w, ans)
}

// 分治做法
func CF1333C2(_r io.Reader, _w io.Writer) {
	ans := int64(0)
	var f func([]int64)
	f = func(a []int64) {
		n := len(a)
		if n == 1 {
			if a[0] != 0 {
				ans++
			}
			return
		}
		m := n >> 1
		f(a[:m])
		f(a[m:])
		sumPos := map[int64]int{}
		s := int64(0)
		maxPos := m
		for i := m - 1; i >= 0 && a[i] != 0; i-- {
			s += a[i]
			if s == 0 || sumPos[s] != 0 {
				break
			}
			sumPos[s] = i
			maxPos = i
		}
		s = 0
		vis := map[int64]bool{}
		for i := m; maxPos < m && i < n && a[i] != 0; i++ {
			s += a[i]
			if s == 0 || vis[s] {
				break
			}
			vis[s] = true
			if p, ok := sumPos[-s]; ok && p+1 > maxPos {
				maxPos = p + 1
			}
			ans += int64(m - maxPos)
		}
	}

	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	f(a)
	Fprint(_w, ans)
}

func main() { CF1333C(os.Stdin, os.Stdout) }
