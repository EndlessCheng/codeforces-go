package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1458A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, m, a1, v int
	Fscan(in, &n, &m, &a1)
	g := 0
	for range n - 1 {
		Fscan(in, &v)
		g = gcd(g, v-a1)
	}
	for range m {
		Fscan(in, &v)
		Fprint(out, abs(gcd(a1+v, g)), " ")
	}
}

//func main() { cf1458A(bufio.NewReader(os.Stdin), os.Stdout) }
