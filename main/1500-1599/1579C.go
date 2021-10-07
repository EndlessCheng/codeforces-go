package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1579C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, m, k int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		a := make([][]byte, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := n - 1; i >= k; i-- {
			for j, b := range a[i] {
				if b != '.' {
					l, r := 1, 1
					for ; i >= l && j >= l && a[i-l][j-l] != '.'; l++ { // 可以通过记录左上和右上最近的 . 的位置来优化循环
					}
					for ; i >= r && j+r < m && a[i-r][j+r] != '.'; r++ {
					}
					if r < l {
						l = r
					}
					if l > k {
						for l--; l >= 0; l-- {
							a[i-l][j-l] = 0
							a[i-l][j+l] = 0
						}
					}
				}
			}
		}
		for _, r := range a {
			for _, v := range r {
				if v == '*' {
					Fprintln(out, "NO")
					continue o
				}
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1579C(os.Stdin, os.Stdout) }
