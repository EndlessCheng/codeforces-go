package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF375B(_r io.Reader, out io.Writer) {
	var n, m, ans int
	Fscanln(_r, &n, &m)
	in := bufio.NewScanner(_r)
	r1 := make([][5001]int16, n)
	for i := range r1 {
		in.Scan()
		s := in.Bytes()
		for j := m - 1; j >= 0; j-- {
			if s[j] == '1' {
				r1[i][j] = r1[i][j+1] + 1
			}
		}
	}
	col := make([]int, n)
	for j := 0; j < m; j++ {
		for i := range col {
			col[i] = int(r1[i][j])
		}
		sort.Ints(col)
		for i, c := range col {
			if r := c * (n - i); r > ans {
				ans = r
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF375B(os.Stdin, os.Stdout) }
