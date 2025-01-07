package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1948D(in io.Reader, out io.Writer) {
	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		for m := n / 2; m > 0; m-- {
			cnt := 0
			for i := m; i < n; i++ {
				if s[i] != '?' && s[i-m] != '?' && s[i] != s[i-m] {
					cnt = 0
					continue
				}
				cnt++
				if cnt == m {
					Fprintln(out, m*2)
					continue o
				}
			}
		}
		Fprintln(out, 0)
	}
}

//func main() { cf1948D(bufio.NewReader(os.Stdin), os.Stdout) }

func cf1948D2(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		lcp := make([][]int16, n+1)
		for i := range lcp {
			lcp[i] = make([]int16, n+1)
		}
		for i := n - 1; i >= 0; i-- {
			for j := n - 1; j >= i; j-- {
				if s[i] == '?' || s[j] == '?' || s[i] == s[j] {
					lcp[i][j] = lcp[i+1][j+1] + 1
				}
			}
		}
	o:
		for m := n / 2; m >= 0; m-- {
			for i := m; i <= n-m; i++ {
				if lcp[i-m][i] >= int16(m) {
					Fprintln(out, m*2)
					break o
				}
			}
		}
	}
}
