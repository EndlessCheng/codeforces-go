package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf794E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}

	ans := make([]int, n+1)
	f := make([]int, n+1)
	g := make([]int, n+1)
	mx := 0
	for i := 1; i <= n; i++ {
		f[i] = max(a[i], a[i+1])
		g[i] = max(min(a[i], a[i-1]), min(a[i], a[i+1]))
		mx = max(mx, a[i])
	}

	if n&1 > 0 {
		ans[1] = g[n/2+1]
	} else {
		ans[1] = f[n/2]
	}
	for i := 2; i < n; i++ {
		if (n+i)&1 > 0 {
			ans[i] = max(ans[i-2], f[(n+1-i)/2], f[(n+i)/2])
		} else {
			ans[i] = max(ans[i-2], g[(n+2-i)/2], g[(n+1+i)/2])
		}
	}

	for i := 1; i < n; i++ {
		Fprint(out, ans[i], " ")
	}
	Fprint(out, mx)
}

//func main() { cf794E(bufio.NewReader(os.Stdin), os.Stdout) }
