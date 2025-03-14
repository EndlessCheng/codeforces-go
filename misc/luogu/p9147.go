package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p9147(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := 1
	f := make([][2]int, n)
	f[0][0] = 1
	f[0][1] = 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			f[i][0] = f[i-1][0] + 1
			f[i][1] = f[i-1][1] + 1
		} else {
			f[i][0] = 1
			f[i][1] = 2
		}
		if i > 1 && a[i] > a[i-2]+1 {
			f[i][1] = max(f[i][1], f[i-2][0]+2)
		}
		ans = max(ans, f[i][0], f[i][1], f[i-1][0]+1)
	}
	Fprint(out, ans)
}

//func main() { p9147(bufio.NewReader(os.Stdin), os.Stdout) }
