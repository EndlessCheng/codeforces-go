package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf44C(in io.Reader, out io.Writer) {
	var n, m, l, r int
	Fscan(in, &n, &m)
	cnt := make([]int, n+1)
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		for i := l; i <= r; i++ {
			cnt[i]++
		}
	}
	for i := 1; i <= n; i++ {
		if cnt[i] != 1 {
			Fprint(out, i, cnt[i])
			return
		}
	}
	Fprint(out, "OK")
}

//func main() { cf44C(os.Stdin, os.Stdout) }
