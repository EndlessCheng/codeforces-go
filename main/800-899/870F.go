package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf870F(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n)
	p := make([]int, n+1)
	s := make([]int, n+1)
	phi := make([]int, n+1)
	used := make([]bool, n+1)
	phi[1] = 1
	for i := 2; i <= n; i++ {
		if !used[i] {
			m++
			p[m] = i
			phi[i] = i - 1
			s[i]++
		}
		ans += phi[i] - 1

		for j, x := 1, i<<1; j <= m && x <= n; {
			used[x] = true
			s[p[j]]++
			if i%p[j] != 0 {
				phi[x] = phi[i] * (p[j] - 1)
			} else {
				phi[x] = phi[i] * p[j]
				break
			}
			j++
			x = i * p[j]
		}
	}

	ans += (n - 1) * (n - 2) / 2
	for i := 2; i <= n; i++ {
		s[i] += s[i-1]
	}
	for i := 1; i <= m; i++ {
		x := max(p[i]+1, n>>1+1)
		y := max(p[i]+1, n/p[i]+1)
		ans += (s[p[i]] - s[p[i-1]]) * (s[x-1] - s[y-1] - (s[n]-s[x-1])<<1)
	}
	Fprint(out, ans)
}

//func main() { cf870F(bufio.NewReader(os.Stdin), os.Stdout) }
