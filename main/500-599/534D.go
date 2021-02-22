package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF534D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, v int
	Fscan(in, &n)
	id := make([][]int, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		id[v] = append(id[v], i)
	}
	ans := make([]interface{}, n)
	for i, c := 0, 0; i < n; i++ {
		for c >= 0 && len(id[c]) == 0 {
			c -= 3
		}
		if c < 0 {
			Fprint(out, "Impossible")
			return
		}
		ans[i] = id[c][0]
		id[c] = id[c][1:]
		c++
	}
	Fprintln(out, "Possible")
	Fprintln(out, ans...)
}

//func main() { CF534D(os.Stdin, os.Stdout) }
