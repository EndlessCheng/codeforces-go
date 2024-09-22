package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf1983C(in io.Reader, out io.Writer) {
	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := [3][]int{}
		for i := range s {
			s[i] = make([]int, n+1)
			for j := 1; j <= n; j++ {
				Fscan(in, &s[i][j])
				s[i][j] += s[i][j-1]
			}
		}
		tar := (s[0][n] + 2) / 3
		for _, p := range perm3 {
			i, j, k := p[0], p[1], p[2]
			l := sort.SearchInts(s[i], tar)
			r := sort.SearchInts(s[k], s[0][n]-tar+1) - 1
			if s[j][r]-s[j][l] >= tar {
				ans := [6]any{}
				ans[i*2] = 1
				ans[i*2+1] = l
				ans[j*2] = l + 1
				ans[j*2+1] = r
				ans[k*2] = r + 1
				ans[k*2+1] = n
				Fprintln(out, ans[:]...)
				continue o
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { cf1983C(bufio.NewReader(os.Stdin), os.Stdout) }
