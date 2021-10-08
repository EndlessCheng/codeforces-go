package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1579F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, d, v int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &d)
		q := []int{}
		vis := make([]bool, n)
		for i := 0; i < n; i++ {
			if Fscan(in, &v); v == 0 { // 注：这里若改用 bool 读入会快不少
				q = append(q, i)
				vis[i] = true
			}
		}
		ans := -1
		for ; len(q) > 0; ans++ {
			tmp := q
			q = nil
			for _, v := range tmp {
				if w := (v + d) % n; !vis[w] {
					q = append(q, w)
					vis[w] = true
				}
			}
		}
		for _, v := range vis {
			if !v {
				Fprintln(out, -1)
				continue o
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1579F(os.Stdin, os.Stdout) }
