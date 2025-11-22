package main

import (
	"cmp"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf498A(in io.Reader, out io.Writer) {
	var x1, y1, x2, y2, n, a, b, c, ans int
	Fscan(in, &x1, &y1, &x2, &y2, &n)
	for range n {
		Fscan(in, &a, &b, &c)
		if cmp.Compare(a*x1+b*y1+c, 0)*cmp.Compare(a*x2+b*y2+c, 0) < 0 {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf498A(bufio.NewReader(os.Stdin), os.Stdout) }
