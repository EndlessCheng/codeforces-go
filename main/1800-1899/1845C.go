package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1845C1(in io.Reader, out io.Writer) {
	var T int
	var s, l, r string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &l, &l, &r)
		n := len(s)
		nxt := make([][10]int, n+1)
		for j := range nxt[n] {
			nxt[n][j] = n
		}
		for i := n - 1; i >= 0; i-- {
			nxt[i] = nxt[i+1]
			nxt[i][s[i]-'0'] = i
		}

		cur := -1
		for i, b := range l {
			cur = slices.Max(nxt[cur+1][b-'0' : r[i]-'0'+1])
			if cur >= n {
				Fprintln(out, "YES")
				continue o
			}
		}
		Fprintln(out, "NO")
	}
}

func cf1845C(in io.Reader, out io.Writer) {
	var T, m int
	var s, l, r string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &m, &l, &r)
		vis, j := 0, 0
		for _, b := range s {
			vis |= 1 << (b - '0')
			msk := 1<<(r[j]-l[j]+1) - 1
			if vis>>(l[j]-'0')&msk == msk {
				j++
				if j == m {
					Fprintln(out, "NO")
					continue o
				}
				vis = 0
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1845C(bufio.NewReader(os.Stdin), os.Stdout) }
