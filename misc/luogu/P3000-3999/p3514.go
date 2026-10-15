package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3514(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, x, sum int
	var s string
	Fscan(in, &n, &m, &s)

	minS := [2]int{0, 1e9}
	pos := [2]int{0, -1}
	mx := [2]int{}
	var l, r [2e6 + 1]int
	for i, b := range s {
		sum += 2 - int(b&1)
		p := sum & 1
		for k := range 2 {
			t := sum - minS[p^k]
			if t > mx[k] {
				mx[k] = t
				l[t] = pos[p^k]
				r[t] = i + 1
			}
		}
		if pos[p] < 0 {
			pos[p] = i + 1
			minS[p] = sum
		}
	}

	f := func(v int) {
		i, j := l[v], r[v]
		for v > 2 {
			if s[i] == 'T' {
				i++
			} else if s[j-1] == 'T' {
				j--
			} else {
				i++
				j--
			}
			v -= 2
			l[v] = i
			r[v] = j
		}
	}
	f(mx[0])
	f(mx[1])

	for range m {
		Fscan(in, &x)
		if r[x] > 0 {
			Fprintln(out, l[x]+1, r[x])
		} else {
			Fprintln(out, "NIE")
		}
	}
}

//func main() { p3514(bufio.NewReader(os.Stdin), os.Stdout) }
