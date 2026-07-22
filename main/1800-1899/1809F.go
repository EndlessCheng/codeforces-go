package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1809F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	f := func(x, k int) int {
		if x <= k {
			return x
		}
		return k + (x-k)*2
	}

	var T int
	for Fscan(in, &T); T > 0; T-- {
		var n, k int
		Fscan(in, &n, &k)
		a := make([]int, n+1)
		b := make([]int, n+1)
		c := make([]int, n+1)
		c2 := make([]int, n+1)

		all := 0
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			all += a[i]
		}
		for i := 1; i <= n; i++ {
			Fscan(in, &b[i])
		}

		ans := 0
		fl := false
		for i := 1; i <= n; i++ {
			if b[i] == 1 {
				fl = true
				u := i
				z := a[i]
				for b[u%n+1] == 2 {
					u = u%n + 1
					z += a[u]
				}
				ans += f(z, k)

				u = i
				z2 := a[i]
				for b[u%n+1] == 2 {
					u = u%n + 1
					c[u] = z2
					c2[u] = z - z2
					z2 += a[u]
				}
			}
		}

		for i := 1; i <= n; i++ {
			if !fl {
				Fprint(out, all*2, " ")
			} else if b[i] == 1 {
				Fprint(out, ans, " ")
			} else {
				Fprint(out, ans-f(c[i]+c2[i], k)+f(c[i], k)+c2[i]*2, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1809F(bufio.NewReader(os.Stdin), os.Stdout) }
