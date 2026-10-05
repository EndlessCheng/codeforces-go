package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1582G(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	mx := slices.Max(a)
	Fscan(in, &s)
	b := make([]int, n)
	e := make([]int, n)
	pd := make([][]int, mx+1)
	v := make([][]int, mx+1)
	for i := 2; i <= mx; i++ {
		if len(pd[i]) == 0 {
			for j := 1; j*i <= mx; j++ {
				k := i * j
				for k%i == 0 {
					pd[i*j] = append(pd[i*j], i)
					k /= i
				}
			}
		}
	}

	for i := range n {
		if s[i] == '*' {
			b[i] = 1
		}
		if a[i] == 1 {
			b[i] = 1
		}
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		if b[i] == 0 {
			e[i] = i
			for _, x := range pd[a[i]] {
				v[x] = append(v[x], i)
			}
			continue
		}

		for _, x := range pd[a[i]] {
			if len(v[x]) > 0 {
				t := v[x][len(v[x])-1]
				v[x] = v[x][:len(v[x])-1]
				a[t] /= x
			}
		}

		s1 := i + 1
		for s1 < n {
			if e[s1] == n {
				s1 = n
				break
			}
			if a[e[s1]] == 1 {
				s1 = e[s1] + 1
			} else {
				break
			}
		}

		if s1 == n {
			e[i] = n
		} else {
			e[i] = e[s1]
		}
		ans += e[i] - i
	}
	Fprint(out, ans)
}

//func main() { cf1582G(bufio.NewReader(os.Stdin), os.Stdout) }
