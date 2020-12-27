package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }

	var n int
	Fscan(in, &n)
	l := 1
	for i := 2; i <= n; i++ {
		l = lcm(l, i)
	}
	Fprint(out, l+1)
}

func main() { run(os.Stdin, os.Stdout) }
