package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF407B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mod int = 1e9 + 7
	var n, p int
	Fscan(in, &n)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &p)
		// 前缀和递推式
		s[i+1] = (s[i]*2%mod + mod - s[p-1] + 2) % mod
	}
	Fprint(out, s[n])
}

//func main() { CF407B(os.Stdin, os.Stdout) }
