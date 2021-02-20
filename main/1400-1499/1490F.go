package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1490F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		cc := make([]int, n+1)
		for _, c := range cnt {
			cc[c]++
		}
		save, s := 0, 0
		for i, c := range cc {
			if i*(len(cnt)-s) > save {
				save = i * (len(cnt) - s)
			}
			s += c
		}
		Fprintln(out, n-save)
	}
}

//func main() { CF1490F(os.Stdin, os.Stdout) }
