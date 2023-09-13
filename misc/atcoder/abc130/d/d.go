package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, l, s, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
		for s >= k {
			s -= a[l]
			l++
		}
		ans += l
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
