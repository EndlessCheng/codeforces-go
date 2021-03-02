package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1494C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		f := func(a, b []int) (res int) {
			has := map[int]bool{}
			for _, v := range a {
				has[v] = true
			}
			same := 0
			for _, v := range b {
				if has[v] {
					same++
				}
			}
			for i, v := range b {
				if has[v] {
					same--
				}
				cnt := sort.SearchInts(a, v+1)
				j := sort.SearchInts(b, v-cnt+1)
				if i-j+1+same > res {
					res = i - j + 1 + same
				}
			}
			return
		}
		x, y := sort.SearchInts(a, 0), sort.SearchInts(b, 0)
		ans := f(a[x:], b[y:])
		for i := 0; i < (x+1)/2; i++ {
			a[i], a[x-1-i] = -a[x-1-i], -a[i]
		}
		for i := 0; i < (y+1)/2; i++ {
			b[i], b[y-1-i] = -b[y-1-i], -b[i]
		}
		ans += f(a[:x], b[:y])
		Fprintln(out, ans)
	}
}

//func main() { CF1494C(os.Stdin, os.Stdout) }
