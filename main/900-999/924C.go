package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF924C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var n int
	Fscan(in, &n)
	m := make([]int, n)
	for i := range m {
		Fscan(in, &m[i])
	}

	sufMax := make([]int, n+1)
	for i := n - 1; i > 0; i-- {
		sufMax[i] = max(sufMax[i+1]-1, m[i]+1)
	}

	ans := int64(0)
	for i, mx := 1, 1; i < n; i++ {
		mx = max(mx, sufMax[i])
		ans += int64(mx - m[i] - 1)
	}
	Fprint(out, ans)
}

//func main() { CF924C(os.Stdin, os.Stdout) }
