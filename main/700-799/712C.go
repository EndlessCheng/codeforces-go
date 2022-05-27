package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF712C(in io.Reader, out io.Writer) {
	var t, x, ans int
	Fscan(in, &t, &x)
	for y, z := x, x; x < t; ans++ {
		x, y, z = y, z, y+z-1
	}
	Fprint(out, ans)
}

//func main() { CF712C(os.Stdin, os.Stdout) }
