package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1859E(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]struct{ x, y int }, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i].x)
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i].y)
		}
		f := make([]int, k+1)
		mx := make([][4]int, n+1)
		for i := range mx {
			for j := range mx[i] {
				mx[i][j] = -1e18
			}
		}
		for i := 1; i <= n; i++ {
			x, y := a[i].x, a[i].y
			for j := min(i, k); j > 0; j-- {
				mx[i-j][0] = max(mx[i-j][0], f[j-1]-y-x)
				mx[i-j][1] = max(mx[i-j][1], f[j-1]+y-x)
				mx[i-j][2] = max(mx[i-j][2], f[j-1]-y+x)
				mx[i-j][3] = max(mx[i-j][3], f[j-1]+y+x)
				f[j] = max(f[j],
					mx[i-j][0]+x+y,
					mx[i-j][1]-x+y,
					mx[i-j][2]+x-y,
					mx[i-j][3]-x-y)
			}
		}
		Fprintln(out, f[k])
	}
}

//func main() { cf1859E(bufio.NewReader(os.Stdin), os.Stdout) }
