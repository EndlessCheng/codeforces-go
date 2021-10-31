package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1005C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, ans int
	Fscan(in, &n)
	cnt := make(map[int]int, n)
	for ; n > 0; n-- {
		Fscan(in, &v)
		cnt[v]++
	}
o:
	for v, c := range cnt {
		for i := 1; i < 31; i++ {
			if s := 1 << i; v*2 != s && cnt[s-v] > 0 || v*2 == s && c > 1 {
				continue o
			}
		}
		ans += c
	}
	Fprint(out, ans)
}

//func main() { CF1005C(os.Stdin, os.Stdout) }
