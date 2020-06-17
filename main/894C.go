package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF894C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, g int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		g = gcd(g, a[i])
	}
	if g != a[0] {
		Fprint(out, -1)
		return
	}
	Fprintln(out, 2*n)
	for _, v := range a {
		Fprint(out, v, a[0], " ")
	}
}

//func main() { CF894C(os.Stdin, os.Stdout) }
