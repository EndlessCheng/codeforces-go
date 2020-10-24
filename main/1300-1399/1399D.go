package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1399D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		ans := make([]interface{}, n)
		id := [2][]int{}
		for i, b := range s {
			b &= 1
			k := len(id[0]) + len(id[1]) + 1
			if len(id[b^1]) > 0 {
				k, id[b^1] = id[b^1][0], id[b^1][1:]
			}
			id[b] = append(id[b], k)
			ans[i] = k
		}
		Fprintln(out, len(id[0])+len(id[1]))
		Fprintln(out, ans...)
	}
}

//func main() { CF1399D(os.Stdin, os.Stdout) }
