package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, v, s, p, q, r int
	Fscan(in, &n, &p, &q, &r)
	has := map[int]bool{0: true}
	for range n {
		Fscan(in, &v)
		s += v
		has[s] = true
	}

	for s := range has {
		if has[s+p] && has[s+p+q] && has[s+p+q+r] {
			Fprint(out, "Yes")
			return
		}
	}
	Fprint(out, "No")
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
