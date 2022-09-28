package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF597C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	f := make([]int64, n)
	for i := range a {
		Fscan(in, &a[i])
		f[i] = 1
	}
	for ; k > 0; k-- {
		tree := make([]int64, n+1)
		add := func(i int, val int64) {
			for ; i < len(tree); i += i & -i {
				tree[i] += val
			}
		}
		sum := func(i int) (res int64) {
			for ; i > 0; i &= i - 1 {
				res += tree[i]
			}
			return
		}
		for i, v := range a {
			add(v, f[i])
			f[i] = sum(v - 1)
		}
	}
	ans := int64(0)
	for _, v := range f {
		ans += v
	}
	Fprintln(out, ans)
}

//func main() { CF597C(os.Stdin, os.Stdout) }
