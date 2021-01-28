package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, v int
	Fscan(in, &n, &m)
	c := map[int]int{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		c[v]++
	}
	s := int64(0)
	for ; m > 0; m-- {
		Fscan(in, &v)
		s += int64(c[v])
	}
	Fprint(out, s)
}

func main() { run(os.Stdin, os.Stdout) }
