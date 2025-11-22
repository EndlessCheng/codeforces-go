package main

import (
	"cmp"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf498A(in io.Reader, out io.Writer) {
	var x1, y1, x2, y2, n, a, b, c, ans int
	f := func(x, y int) int { return a*x + b*y + c }
	Fscan(in, &x1, &y1, &x2, &y2, &n)
	for range n {
		Fscan(in, &a, &b, &c)
		if cmp.Compare(f(x1, y1), 0)*cmp.Compare(f(x2, y2), 0) < 0 {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf498A(bufio.NewReader(os.Stdin), os.Stdout) }
