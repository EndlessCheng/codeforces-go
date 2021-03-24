package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1239A(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	f := [1e5 + 1]int{1, 1}
	for i := 2; i <= 1e5; i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}
	var n, m int
	Fscan(in, &n, &m)
	Fprint(out, ((f[n]+f[m])%mod+mod-1)%mod*2%mod)
}

//func main() { CF1239A(os.Stdin, os.Stdout) }
