package main

import (
	"bufio"
	. "fmt"
	"io"
	"maps"
	"slices"
)

// https://github.com/EndlessCheng
func cf666A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var s string
	Fscan(in, &s)
	ans := map[string]bool{}

	n := len(s)
	vis := make([][4]bool, n)
	var dfs func(int, string)
	dfs = func(i int, pre string) {
		if i < 6 || vis[i][len(pre)] {
			return
		}
		vis[i][len(pre)] = true

		suf := s[i-1 : i+1]
		if suf != pre {
			ans[suf] = true
			dfs(i-2, suf)
		}

		suf = s[i-2 : i+1]
		if i > 6 && suf != pre {
			ans[suf] = true
			dfs(i-3, suf)
		}
	}
	dfs(n-1, "")

	suf := slices.Sorted(maps.Keys(ans))
	Fprintln(out, len(suf))
	for _, v := range suf {
		Fprintln(out, v)
	}
}

//func main() { cf666A(bufio.NewReader(os.Stdin), os.Stdout) }
