package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1513B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 1_000_000_007

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		and := -1
		for i := range a {
			Fscan(in, &a[i])
			and &= a[i]
		}
		cnt := 0
		for _, v := range a {
			if v == and {
				cnt++
			}
		}
		ans := cnt * (cnt - 1)
		for i := 2; i <= n-2; i++ {
			ans = ans * i % mod
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1513B(os.Stdin, os.Stdout) }
