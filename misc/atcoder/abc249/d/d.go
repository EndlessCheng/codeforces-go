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
	const mx int = 2e5
	var n, v, ans int
	cnt := [mx + 1]int{}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
	for i := 1; i <= mx; i++ {
		for j := 1; i*j <= mx; j++ {
			ans += cnt[i] * cnt[j] * cnt[i*j]
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
