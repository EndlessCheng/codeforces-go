package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1283A(in io.Reader, out io.Writer) {
	var T, h, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &h, &m)
		Fprintln(out, 1440-h*60-m)
	}
}

//func main() { CF1283A(os.Stdin, os.Stdout) }
