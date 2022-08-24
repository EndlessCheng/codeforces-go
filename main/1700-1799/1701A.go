package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1701A(in io.Reader, out io.Writer) {
	var T, v int
	for Fscan(in, &T); T > 0; T-- {
		c := 0
		for i := 0; i < 4; i++ {
			Fscan(in, &v)
			c += v
		}
		if c == 0 {
			Fprintln(out, 0)
		} else if c < 4 {
			Fprintln(out, 1)
		} else {
			Fprintln(out, 2)
		}
	}
}

//func main() { CF1701A(os.Stdin, os.Stdout) }
