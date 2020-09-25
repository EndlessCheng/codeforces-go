package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1034B(in io.Reader, out io.Writer) {
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	var n, m, ans int64
	Fscan(in, &n, &m)
	if n > m {
		n, m = m, n
	}
	if n == 1 {
		ans = m - m%6 + max(m%6-3, 0)*2
	} else if n == 2 {
		switch m {
		case 2:
			ans = 0
		case 3:
			ans = 4
		case 7:
			ans = 12
		default:
			ans = m * 2
		}
	} else {
		ans = n * m &^ 1
	}
	Fprint(out, ans)
}

//func main() { CF1034B(os.Stdin, os.Stdout) }
