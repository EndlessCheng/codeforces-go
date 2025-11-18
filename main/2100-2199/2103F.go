package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2103F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := slices.Clone(a)
		s := make([][2]int, n)
		for i := range s {
			s[i] = [2]int{-1, -1}
		}
		maxS := make([][2]int, n+1)
		for i := n - 2; i >= 0; i-- {
			mx := 0
			nor := a[i]
			r := i + 1
			for ; r < n; r++ {
				nor = 1<<k - 1 ^ (nor | a[r])
				if s[r][i&1] == nor {
					mx = maxS[r][i&1]
					break
				}
				s[r][i&1] = nor
			}
			for j := min(r, n-1); j >= i; j-- {
				mx = max(mx, s[j][i&1])
				ans[j] = max(ans[j], mx)
				maxS[j][0] = max(maxS[j+1][0], s[j][0])
				maxS[j][1] = max(maxS[j+1][1], s[j][1])
			}
		}

		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2103F(bufio.NewReader(os.Stdin), os.Stdout) }
