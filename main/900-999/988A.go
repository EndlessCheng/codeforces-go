package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF988A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, v int
	Fscan(in, &n, &k)
	vis := [101]bool{}
	var ans []interface{}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if !vis[v] {
			ans = append(ans, i)
			if k--; k == 0 {
				break
			}
			vis[v] = true
		}
	}
	if k > 0 {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	Fprint(out, ans...)
}

//func main() { CF988A(os.Stdin, os.Stdout) }
