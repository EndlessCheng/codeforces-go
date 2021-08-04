package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1335E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := make([][]int, 200)
		for i := range a {
			Fscan(in, &a[i])
			a[i]--
			pos[a[i]] = append(pos[a[i]], i)
		}
		ans := 0
		s := make([]int, n+1)
		for i, ps := range pos {
			ans = max(ans, len(ps))
			for j, v := range a {
				s[j+1] = s[j]
				if v == i {
					s[j+1]++
				}
			}
			for _, ps := range pos {
				for k, m := 0, len(ps); k < m/2; k++ {
					ans = max(ans, (k+1)*2+s[ps[m-1-k]]-s[ps[k]+1])
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1335E2(os.Stdin, os.Stdout) }
