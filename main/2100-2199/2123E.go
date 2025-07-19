package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2123E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		cc := make([]int, n+1)
		mex := 0
		for cnt[mex] > 0 {
			cc[cnt[mex]]++
			mex++
		}

		s := 1
		for _, c := range cc[:n-mex+1] {
			s += c
			Fprint(out, s, " ")
		}
		for i := mex; i > 0; i-- {
			Fprint(out, i, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2123E(bufio.NewReader(os.Stdin), os.Stdout) }
