package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol1200E(reader io.Reader, writer io.Writer) {
	maxMergeLen := func(s string) int {
		n := len(s)
		f := make([]int, n)
		cnt := 0
		for i := 1; i < n; i++ {
			c := s[i]
			for cnt > 0 && s[cnt] != c {
				cnt = f[cnt-1]
			}
			if s[cnt] == c {
				cnt++
			}
			f[i] = cnt
		}
		return f[n-1]
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var w, tmpMerge string
	Fscan(in, &n, &w)
	ans := append(make([]byte, 0, 1e6), w...)
	for n--; n > 0; n-- {
		Fscan(in, &w)
		if len(ans) >= len(w) {
			tmpMerge = w + "$" + string(ans[len(ans)-len(w):])
		} else {
			tmpMerge = w[:len(ans)] + "$" + string(ans)
		}
		if l := maxMergeLen(tmpMerge); l < len(w) {
			ans = append(ans, w[l:]...)
		}
	}
	Fprintln(out, string(ans))
}

//func main() {
//	Sol1200E(os.Stdin, os.Stdout)
//}
