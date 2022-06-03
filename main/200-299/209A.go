package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF209A(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var n int
	Fscan(in, &n)
	f := make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = (f[i-1] + f[i-2] + 2) % mod
	}
	Fprint(out, f[n])
}

//func main() { CF209A(os.Stdin, os.Stdout) }
