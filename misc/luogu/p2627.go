package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p2627(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
	}

	f := make([]int, n+2)
	q := []int{0}
	for i := 1; i <= n; i++ {
		if q[0] < i-k {
			q = q[1:]
		}
		f[i+1] = max(f[i], s[i]+f[q[0]]-s[q[0]])
		for len(q) > 0 && f[q[len(q)-1]]-s[q[len(q)-1]] <= f[i]-s[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	Fprint(out, f[n+1])
}

//func main() { p2627(bufio.NewReader(os.Stdin), os.Stdout) }
