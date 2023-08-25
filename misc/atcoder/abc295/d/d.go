package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(in io.Reader, out io.Writer) {
	var s string
	var ans, mask int
	Fscan(bufio.NewReader(in), &s)
	cnt := [1024]int{1}
	for _, c := range s {
		mask ^= 1 << (c - '0')
		ans += cnt[mask]
		cnt[mask]++
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
