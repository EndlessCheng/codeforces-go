package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1415F(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n int
	Fscan(in, &n)
	t := make([]int, n+1)
	x := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &t[i], &x[i])
	}
	d := func(a, b int) int { return abs(x[a] - x[b]) }

	f := make([]int, n+1)
	for i := range f {
		f[i] = -1e18
	}
	f[0] = 0
	g := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if d(i-1, i) <= f[i-1]+t[i]-t[i-1] {
			f[i] = min(t[i]-t[i-1], f[i-1]+t[i]-t[i-1]-d(i-1, i))
		}
		if i >= 2 && g[i-1] && d(i-2, i) <= t[i]-t[i-2] {
			f[i] = max(f[i], min(t[i]-t[i-1], t[i]-t[i-2]-d(i-2, i)))
		}
		for j := i - 2; j >= 0 && !g[i]; j-- {
			g[i] = f[j] >= 0 && max(f[j], d(j, i))+d(i, j+1) <= f[j]+t[j+1]-t[j] ||
				g[j] && max(t[j]-t[j-1], d(j-1, i))+d(i, j+1) <= t[j+1]-t[j-1]
			if t[j+1]-t[j] < d(j, j+1) {
				break
			}
		}
	}

	if f[n] >= 0 || g[n] {
		Fprint(out, "YES")
	} else {
		Fprint(out, "NO")
	}
}

//func main() { cf1415F(bufio.NewReader(os.Stdin), os.Stdout) }
