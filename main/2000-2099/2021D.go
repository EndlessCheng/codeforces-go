package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2021D(in io.Reader, out io.Writer) {
	const inf int = 1e18
	var t int
	for Fscan(in, &t); t > 0; t-- {
		var n, m int
		Fscan(in, &n, &m)
		f1 := make([]int, m+2)
		f2 := make([]int, m+2)
		f3 := make([]int, m+2)
		f4 := make([]int, m+2)
		a := make([]int, m+2)
		ans := -inf
		fk := 0
		for n > 0 {
			n--
			for i := 1; i <= m; i++ {
				Fscan(in, &a[i])
				a[i] += a[i-1]
				f3[i] = -inf
				f4[i] = -inf
			}
			f := 0
			g := -inf

			for i := 2; i <= m; i++ {
				f = min(f, a[i-2])
				f4[i] = max(f4[i-1], f2[i-1]-f, f1[i]-f)
				j := m + 1 - i
				g = max(g, a[j+1])
				f3[j] = max(f3[j+1], f1[j+1]+g, f2[j]+g)
			}

			for i := 1; i <= m; i++ {
				f1[i] = f3[i] - a[i-1]
				f2[i] = f4[i] + a[i]
				if fk == 0 {
					f1[i] = max(f1[i], a[i]-a[i-1])
					f2[i] = max(f2[i], a[i]-a[i-1])
				}
				if n == 0 {
					ans = max(ans, f1[i], f2[i])
				}
			}
			fk = 1
		}

		Fprintln(out, ans)
	}
}

//func main() { cf2021D(bufio.NewReader(os.Stdin), os.Stdout) }
