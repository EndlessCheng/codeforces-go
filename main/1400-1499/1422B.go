package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1422B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}
		s := int64(0)
		for i := 0; 2*i < n; i++ {
			if 2*i+1 == n {
				for j, v := range a[i][:m/2] {
					s += int64(abs(v - a[i][m-1-j]))
				}
				continue
			}
			for j := 0; 2*j < m; j++ {
				b := []int{a[i][j], a[n-1-i][j]}
				if 2*j+1 == m {
					s += int64(abs(b[0] - b[1]))
					continue
				}
				b = append(b, a[i][m-1-j], a[n-1-i][m-1-j])
				sort.Ints(b)
				s += int64(b[3]-b[1]) + int64(b[2]-b[0])
			}
		}
		Fprintln(out, s)
	}
}

//func main() { CF1422B(os.Stdin, os.Stdout) }
