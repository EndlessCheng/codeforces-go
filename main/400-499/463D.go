package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF463D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n)
	index := make([][]int, k)
	for i := range index {
		index[i] = make([]int, n+1)
		for j := range a {
			Fscan(in, &a[j])
			index[i][a[j]] = j
		}
	}

	f := make([]int, n)
	for i, x := range a { // 以最后一个排列 a 为基准
	next:
		for j, y := range a[:i] { // 枚举在 x 左边的数 y
			for _, idx := range index {
				if idx[y] > idx[x] { // 对于其余排列，y 的位置必须在 x 的左边
					continue next
				}
			}
			f[i] = max(f[i], f[j])
		}
		f[i]++
		ans = max(ans, f[i])
	}
	Fprint(out, ans)
}

//func main() { CF463D(os.Stdin, os.Stdout) }
