package _00_299

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF232A(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	c := func(n, k int) int {
		if k == 2 {
			return n * (n - 1) / 2
		}
		return n * (n - 1) * (n - 2) / 6
	}

	var n int
	Fscan(_r, &n)
	m := 3
	for ; c(m, 3) <= n; m++ {
	}
	m--
	n -= c(m, 3)
	ans := [101][101]int8{}
	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			ans[i][j] = 1
			ans[j][i] = 1
		}
	}
	for i := m - 1; i > 1; i-- {
		for ; c(i, 2) <= n; n -= c(i, 2) {
			for j := 0; j < i; j++ {
				ans[m][j] = 1
				ans[j][m] = 1
			}
			m++
		}
	}
	Fprintln(out, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			Fprint(out, ans[i][j])
		}
		Fprintln(out)
	}
}

//func main() { CF232A(os.Stdin, os.Stdout) }
