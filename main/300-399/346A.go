package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF346A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, v, mx, g int
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		if v > mx {
			mx = v
		}
		g = gcd(g, v)
	}
	if (mx/g-n)&1 > 0 {
		Fprint(out, "Alice")
	} else {
		Fprint(out, "Bob")
	}
}

//func main() { CF346A(os.Stdin, os.Stdout) }
