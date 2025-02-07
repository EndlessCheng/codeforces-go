package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	s := ""
	Fscan(in, &s)
	n := len(s)
	ans0, ans1 := n, n
	for i, b := range s {
		if b == '0' {
			ans0 = min(ans0, max(n-1-i, i))
		} else {
			ans1 = min(ans1, max(n-1-i, i))
		}
	}
	Fprint(out, max(ans0, ans1))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
