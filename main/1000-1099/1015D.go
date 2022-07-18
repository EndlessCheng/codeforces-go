package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1015D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, s, d int64
	Fscan(in, &n, &k, &s)
	if s < k || s > k*(n-1) {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	cur := int64(1)
	for k > 0 && s-(n-1) >= k-1 {
		s -= n - 1
		k--
		cur = n + 1 - cur
		Fprint(out, cur, " ")
	}
	if k == 0 {
		return
	}
	if cur == 1 {
		cur += s - k + 1
		d = 1
	} else {
		cur -= s - k + 1
		d = -1
	}
	Fprint(out, cur, " ")
	for ; k > 1; k-- {
		d = -d
		cur += d
		Fprint(out, cur, " ")
	}
}

//func main() { CF1015D(os.Stdin, os.Stdout) }
