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
	var n, v, s, ans int
	cnt := map[int]int{0: 1}
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		s += v * (n%2*2 - 1)
		ans += cnt[s]
		cnt[s]++
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
