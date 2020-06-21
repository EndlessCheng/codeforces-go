package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1370C(in io.Reader, out io.Writer) {
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		m := n
		c2, c := 0, 0
		for i := 2; i*i <= n; i++ {
			for ; n%i == 0; n /= i {
				if i == 2 {
					c2++
				} else {
					c++
				}
			}
		}
		if n > 1 {
			if n == 2 {
				c2++
			} else {
				c++
			}
		}
		if m == 2 || c2 == 0 && c > 0 || c2 == 1 && c > 1 || c2 > 1 && c > 0 {
			Fprintln(out, "Ashishgup")
		} else {
			Fprintln(out, "FastestFinger")
		}
	}
}

//func main() { CF1370C(os.Stdin, os.Stdout) }
