package main

import (
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(in io.Reader, out io.Writer) {
	a := []int{}
	var build func(int)
	build = func(v int) {
		if v > 4444444444 {
			return
		}
		a = append(a, v)
		build(v*10 + 4)
		build(v*10 + 7)
	}
	build(0) // 1024 个数
	sort.Ints(a)

	f := func(n int) (ans int) {
		for i := 1; ; i++ {
			cur := a[i]
			if cur > n {
				cur = n
			}
			ans += (cur - a[i-1]) * a[i]
			if cur == n {
				return
			}
		}
	}

	var l, r int
	Fscan(in, &l, &r)
	Fprintln(out, f(r)-f(l-1))
}

func main() { run(os.Stdin, os.Stdout) }
