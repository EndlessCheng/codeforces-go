package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3903(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	f := make([][2]int, n)
	for i := range a {
		Fscan(in, &a[i])
		for j, w := range a[:i] {
			if w > a[i] {
				f[i][0] = max(f[i][0], f[j][1]+1)
			} else if w < a[i] {
				f[i][1] = max(f[i][1], f[j][0])
			}
		}
		f[i][1]++
		ans = max(ans, f[i][0], f[i][1])
	}
	Fprint(out, ans)
}

//func main() { p3903(bufio.NewReader(os.Stdin), os.Stdout) }
