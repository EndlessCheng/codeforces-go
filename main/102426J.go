package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF102426J(_r io.Reader, _w io.Writer) {
	var n int
	Fscan(_r, &n)
	used := make([]bool, 2*n+1)
	ans := int64(0)
	for i := n; i > 0; i-- {
		if !used[2*i] {
			ans += int64(i)
			used[i] = true
		}
	}
	Fprint(_w, ans)
}

//func main() { CF102426J(os.Stdin, os.Stdout) }
