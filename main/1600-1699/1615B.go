package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1615B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e5
	s := [mx + 1][18]int{}
	for i := 1; i <= mx; i++ {
		s[i] = s[i-1]
		for x, j := i, 0; x > 0; x >>= 1 {
			s[i][j] += x & 1
			j++
		}
	}

	var T, l, r int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &l, &r)
		l--
		max := 0
		for i := 0; i < 18; i++ {
			if c := s[r][i] - s[l][i]; c > max {
				max = c
			}
		}
		Fprintln(out, r-l-max)
	}
}

//func main() { CF1615B(os.Stdin, os.Stdout) }
