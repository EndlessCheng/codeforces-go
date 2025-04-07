package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf862A(in io.Reader, out io.Writer) {
	var n, x, v, ans int
	Fscan(in, &n, &x)
	has := make([]bool, x)
	for range n {
		Fscan(in, &v)
		if v < x {
			has[v] = true
		} else if v == x {
			ans++
		}
	}
	for _, b := range has {
		if !b {
			ans++
		}
	}
	Fprint(out, ans)
}

//func main() { cf862A(os.Stdin, os.Stdout) }
