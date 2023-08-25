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
	var n, k, s, v, ans int
	cnt := map[int]int{0: 1}
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &v)
		s += v
		ans += cnt[s-k]
		cnt[s]++
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
