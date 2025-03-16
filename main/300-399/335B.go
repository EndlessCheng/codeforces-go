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
	for i := n - 1; i >= 0; i-- {
		f[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				f[i][j] = f[i+1][j-1] + 2
			} else {
				f[i][j] = max(f[i+1][j], f[i][j-1])
			}
		}
	}

	t := []byte{}
	i, j := 0, n-1
	for i < j && len(t) < 50 {
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
	if i == j && len(t) < 50 {
		t = append(t, s[i])
	}
	Fprintf(out, "%s%s", t, rev)
}

//func main() { cf335B(bufio.NewReader(os.Stdin), os.Stdout) }
