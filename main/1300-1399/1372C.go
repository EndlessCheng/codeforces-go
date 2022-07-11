package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1372C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := 0
		for i, pre := 1, true; i <= n; i++ {
			Fscan(in, &v)
			if v == i != pre {
				c++
				pre = !pre
			}
		}
		if c == 2 {
			c = 1
		} else if c > 2 {
			c = 2
		}
		Fprintln(out, c)
	}
}

//func main() { CF1372C(os.Stdin, os.Stdout) }
