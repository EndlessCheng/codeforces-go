package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF1619E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := make([]int, n+1)
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			cnt[v]++
		}
		Fprint(out, cnt[0], " ")
		s, ex, fill := 0, []int{}, int64(0)
		for i, c := range cnt[:n] {
			if s += c; s <= i {
				Fprint(out, strings.Repeat("-1 ", n-i))
				break
			}
			if c == 0 {
				fill += int64(i - ex[len(ex)-1])
				ex = ex[:len(ex)-1]
			}
			Fprint(out, fill+int64(cnt[i+1]), " ")
			for ; c > 1; c-- {
				ex = append(ex, i)
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1619E(os.Stdin, os.Stdout) }
