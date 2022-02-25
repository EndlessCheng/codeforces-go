package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1641B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cnt := map[int]int{}
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
			cnt[a[i]]++
		}
		for _, c := range cnt {
			if c&1 > 0 {
				Fprintln(out, -1)
				continue o
			}
		}

		op := [][2]int{}
		sz := []int{}
		for len(a) > 0 {
			pos := map[int]int{}
			for i, v := range a {
				p := pos[v]
				if p == 0 {
					pos[v] = i + 1
					continue
				}
				p--

				for j := p - 1; j >= 0; j-- {
					op = append(op, [2]int{i + (len(op)+len(sz))*2 - (p - 1 - j), a[j]})
				}
				for j := p + 1; j < i; j++ {
					op = append(op, [2]int{i + 1 + (len(op)+len(sz))*2 - (j - 1 - p), a[j]})
				}
				for j := p - 1; j >= 0; j-- {
					op = append(op, [2]int{i + 1 + (len(op)+len(sz))*2 - (i - 2 - j), a[j]})
				}
				sz = append(sz, (i+p)*2)

				for i, j := p+1, i-1; i < j; i++ {
					a[i], a[j] = a[j], a[i]
					j--
				}
				copy(a[p:], a[p+1:i])
				copy(a[i-1:], a[i+1:])
				a = a[:len(a)-2]
				break
			}
		}

		Fprintln(out, len(op))
		for _, p := range op {
			Fprintln(out, p[0], p[1])
		}
		Fprintln(out, len(sz))
		for _, v := range sz {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1641B(os.Stdin, os.Stdout) }
