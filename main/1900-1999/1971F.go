package main

import (
	. "fmt"
	"io"
)

func cf1971F(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 1
		x, x2 := n, n
		for y := 1; y <= n; y++ {
			for x*x+y*y >= (n+1)*(n+1) {
				x--
			}
			for x2 > 0 && x2*x2+y*y >= n*n {
				x2--
			}
			ans += x - x2
		}
		Fprintln(out, ans*4)
	}
}

//func main() { cf1971F(bufio.NewReader(os.Stdin), os.Stdout) }
