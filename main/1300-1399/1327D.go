package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1327D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 2e5
	factors := [mx + 1][]int{}
	for i := 1; i <= mx; i++ {
		for j := i; j <= mx; j += i {
			factors[j] = append(factors[j], i)
		}
	}

	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		p := make([]int, n)
		for i := range p {
			Fscan(in, &p[i])
			p[i]--
		}
		colors := make([]int, n)
		for i := range colors {
			Fscan(in, &colors[i])
		}

		ans := n
		vis := make([]bool, n)
		for _, v := range p {
			if vis[v] {
				continue
			}
			cs := []int{}
			for ; !vis[v]; v = p[v] {
				vis[v] = true
				cs = append(cs, colors[v])
			}
		o:
			for _, f := range factors[len(cs)] {
				for i, c := range cs[:f] {
					same := true
					for j := i + f; j < len(cs); j += f {
						if cs[j] != c {
							same = false
							break
						}
					}
					if same {
						if f < ans {
							ans = f
						}
						break o
					}
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1327D(os.Stdin, os.Stdout) }
