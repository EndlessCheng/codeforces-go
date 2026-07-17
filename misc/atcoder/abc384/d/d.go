package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, tar, s int
	Fscan(in, &n, &tar)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		s += a[i]
	}
	tar %= s

	s = 0
	l := 0
	for i := range n * 2 {
		s += a[i%n]
		for s > tar {
			s -= a[l%n]
			l++
		}
		if s == tar {
			Fprint(out, "Yes")
			return
		}
	}
	Fprint(out, "No")
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
