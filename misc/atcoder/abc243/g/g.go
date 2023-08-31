package main

import (
	. "fmt"
	"io"
	"math"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	f := [54773]int{1: 1}
	rt := 2
	for i := 2; i < len(f); i++ {
		f[i] = f[i-1]
		if rt*rt == i {
			f[i] += f[rt]
			rt++
		}
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		m := int(math.Sqrt(float64(n)))
		if m*m > n {
			m--
		}
		ans := 0
		for i := 1; i*i <= m; i++ {
			ans += f[i] * (m - i*i + 1)
		}
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
