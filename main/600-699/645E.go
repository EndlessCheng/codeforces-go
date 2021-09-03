package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF645E(in io.Reader, out io.Writer) {
	const mod int = 1e9 + 7
	var n, k int
	var s []byte
	Fscan(bufio.NewReader(in), &n, &k, &s)
	for i := range s {
		s[i] -= 'a'
	}

	perm := make([]byte, 0, k)
	vis := make([]bool, k)
	for i := len(s) - 1; i >= 0; i-- {
		b := s[i]
		if !vis[b] {
			vis[b] = true
			perm = append(perm, b)
		}
	}
	for i, b := range vis {
		if !b {
			perm = append(perm, byte(i))
		}
	}
	for i := 0; i < k/2; i++ {
		perm[i], perm[k-1-i] = perm[k-1-i], perm[i]
	}
	s = append(append(s, bytes.Repeat(perm, n/k)...), perm[:n%k]...)

	f := [26]int{}
	sumF := 0
	for _, b := range s {
		tmp := (sumF + mod - f[b]) % mod
		f[b] = (sumF + 1) % mod
		sumF = (tmp + f[b]) % mod
	}
	Fprint(out, (sumF+1)%mod)
}

//func main() { CF645E(os.Stdin, os.Stdout) }
