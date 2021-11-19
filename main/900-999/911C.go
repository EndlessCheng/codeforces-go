package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF911C(in io.Reader, out io.Writer) {
	c := map[int]int{}
	for i, v := 0, 0; i < 3; i++ {
		Fscan(in, &v)
		c[v]++
	}
	if c[1] > 0 || c[2] > 1 || c[3] > 2 || c[2] > 0 && c[4] > 1 {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { CF911C(os.Stdin, os.Stdout) }
