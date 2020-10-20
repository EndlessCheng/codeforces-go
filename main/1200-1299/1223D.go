package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1223D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		pos := make([][2]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &v)
			if pos[v][0] == 0 {
				pos[v] = [2]int{i, i}
			} else {
				pos[v][1] = i
			}
		}
		var tot, cnt, mxCnt, prev int
		for i := 1; i <= n; i++ {
			if pos[i][0] == 0 {
				continue
			}
			tot++
			if prev == 0 || pos[i][0] < pos[prev][1] {
				cnt = 1
			} else {
				cnt++
			}
			if cnt > mxCnt {
				mxCnt = cnt
			}
			prev = i
		}
		Fprintln(out, tot-mxCnt)
	}
}

//func main() { CF1223D(os.Stdin, os.Stdout) }
