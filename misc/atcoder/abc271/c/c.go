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
	var n, v, ans int
	Fscan(in, &n)
	s := map[int]int{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[v] = 1
	}
	for n > 0 {
		ans++
		n -= 2 - s[ans]
	}
	Fprint(out, ans+n)
}

func main() { run(os.Stdin, os.Stdout) }
