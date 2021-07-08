package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF494B(in io.Reader, out io.Writer) {
	const mod int64 = 1e9 + 7
	var s, t []byte
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	match := make([]int, m)
	for i, c := 1, 0; i < m; i++ {
		v := t[i]
		for c > 0 && t[c] != v {
			c = match[c-1]
		}
		if t[c] == v {
			c++
		}
		match[i] = c
	}
	c := 0
	sum := make([]int64, n+1)
	iSum := make([]int64, n+1)
	ans := int64(0)
	lastSt := -1
	for i, v := range s {
		for c > 0 && t[c] != v {
			c = match[c-1]
		}
		if t[c] == v {
			c++
		}
		if c == m {
			lastSt = i - m + 1
			c = match[c-1]
		}
		if lastSt < 0 {
			continue
		}
		res := (int64(lastSt+1) + iSum[lastSt] - sum[lastSt]*int64(n-lastSt)) % mod
		ans += res
		sum[i+1] = (res + sum[i]) % mod
		iSum[i+1] = (res*int64(n-i) + iSum[i]) % mod
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { CF494B(os.Stdin, os.Stdout) }
