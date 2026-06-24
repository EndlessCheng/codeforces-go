package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1770F(in io.Reader, out io.Writer) {
	var n, x, y, ans int
	Fscan(in, &n, &x, &y)
	for i := 1; i <= x; i <<= 1 {
		for j := y; j > 0; j = (j - 1) & y {
			if j&i > 0 && (n*j-i)&(x-i) == x-i {
				ans ^= i
			}
		}
	}
	Fprint(out, n&1*ans)
}

//func main() { cf1770F(os.Stdin, os.Stdout) }
