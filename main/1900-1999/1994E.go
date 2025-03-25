package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1994E(in io.Reader, out io.Writer) {
	var T, k, n int
	Fscan(in, &T)
	for range T {
		ans := 0
		Fscan(in, &k)
		for range k {
			Fscan(in, &n)
			c := ans | n
			for i := 19; i >= 0; i-- {
				if ans&n>>i&1 > 0 {
					c |= 1<<i - 1
					break
				}
			}
			ans = c
			for range n - 1 {
				Fscan(in, &n)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1994E(bufio.NewReader(os.Stdin), os.Stdout) }
