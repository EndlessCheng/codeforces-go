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
	var n, m, t, l, r, pre int
	Fscan(in, &n, &m, &t)
	cur := n
	for ; m > 0; m-- {
		Fscan(in, &l, &r)
		cur -= l - pre
		if cur < 1 {
			Fprint(out, "No")
			return
		}
		cur = min(n, cur+r-l)
		pre = r
	}
	if cur-t+pre < 1 {
		Fprint(out, "No")
	} else {
		Fprint(out, "Yes")
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
