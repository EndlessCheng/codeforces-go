package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF486A(in io.Reader, out io.Writer) {
	var n int64
	Fscan(in, &n)
	ans := n / 2
	if n&1 > 0 {
		ans -= n
	}
	Fprint(out, ans)
}

//func main() { CF486A(os.Stdin, os.Stdout) }
