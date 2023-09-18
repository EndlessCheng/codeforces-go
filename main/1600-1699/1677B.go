package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1677B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &s)
		col := make([]bool, m)
		col1 := 0
		window1 := 0
		f := make([]int, m)
		for i, c := range s {
			if c == '1' && !col[i%m] {
				col[i%m] = true
				col1++ // 多了一个有 1 的列
			}
			window1 += int(c & 1)
			if i >= m {
				window1 -= int(s[i-m] & 1)
			}
			if window1 > 0 { // 最新进来的 m 个人中有 1
				f[i%m]++ // 本质是 f[i] = f[i-m] + 1，但可以用长为 m 的数组滚动
			}
			Fprint(out, col1+f[i%m], " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1677B(os.Stdin, os.Stdout) }
