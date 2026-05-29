package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	f := func() []int {
		s := make([]int, n+1)
		for i := range n {
			Fscan(in, &s[i+1])
			s[i+1] += s[i]
		}
		return s
	}
	sa, sb, sc := f(), f(), f()

	mx2, mx := int(-1e18), int(-1e18)
	for i := 1; i < n; i++ {
		mx2 = max(mx2, mx+sb[i]-sc[i])
		mx = max(mx, sa[i]-sb[i])
	}
	Fprint(out, sc[n]+mx2)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
