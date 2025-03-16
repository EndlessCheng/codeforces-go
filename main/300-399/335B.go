package main

import (
	. "fmt"
	"io"
	"slices"
	"strings"
)

// https://github.com/EndlessCheng
func cf335B(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
		if cnt[b-'a'] == 100 {
			Fprint(out, strings.Repeat(string(b), 100))
			return
		}
	}

	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	dp := func(odd int) {
		for i := n - 1; i >= 0; i-- {
			f[i][i] = odd
			for j := i + 1; j < n; j++ {
				if s[i] == s[j] {
					f[i][j] = f[i+1][j-1] + 2
				} else {
					f[i][j] = max(f[i+1][j], f[i][j-1])
				}
			}
		}
	}
	dp(0)
	ok := f[0][n-1] >= 100
	if !ok {
		dp(1)
	}

	t := []byte{}
	i, j := 0, n-1
	for i < j && (!ok || len(t) < 50) {
		if s[i] == s[j] {
			t = append(t, s[i])
			i++
			j--
		} else if f[i][j] == f[i+1][j] {
			i++
		} else {
			j--
		}
	}
	rev := slices.Clone(t)
	slices.Reverse(rev)
	if !ok && i == j {
		t = append(t, s[i])
	}
	t = append(t, rev...)
	Fprintf(out, "%s", t)
}

//func main() { cf335B(bufio.NewReader(os.Stdin), os.Stdout) }
