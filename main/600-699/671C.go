package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf671C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	next := make([]int, n+1)
	mx := 0
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		next[i] = i
		mx = max(mx, a[i])
	}

	s := make([][]int, mx+1)
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= a[i]; j++ {
			if a[i]%j == 0 {
				s[a[i]/j] = append(s[a[i]/j], i)
				if j*j != a[i] {
					s[j] = append(s[j], i)
				}
			}
		}
	}

	ans := 0
	f := func(l, r, g int) {
		for i := l; i <= min(n, r); i++ {
			if next[i] < r {
				ans += g * (r - next[i])
				next[i] = r
			} else {
				break
			}
		}
	}

	for g := mx; g > 0; g-- {
		if len(s[g]) < 2 {
			continue
		}
		f(s[g][1]+1, n+1, g)
		f(s[g][0]+1, s[g][len(s[g])-1], g)
		f(1, s[g][len(s[g])-2], g)
	}
	Fprint(out, ans)
}

//func main() { cf671C(bufio.NewReader(os.Stdin), os.Stdout) }
