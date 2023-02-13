package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1213D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	const mx int = 2e5
	a := [mx + 1][]int{}
	var n, k, v int
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &v)
		if a[v] == nil {
			a[v] = []int{1}
		} else {
			a[v][0]++
		}
	}

	ans := int(1e9)
	for i := mx; i > 0; i-- {
		b := a[i]
		if b == nil {
			continue
		}
		s, left := 0, k
		for j, c := range b {
			if left <= c {
				ans = min(ans, s+left*j)
				break
			}
			s += c * j
			left -= c
		}
		i2 := i >> 1
		if a[i2] == nil {
			a[i2] = append([]int{0}, b...)
		} else {
			for j, c := range b {
				if j+1 == len(a[i2]) {
					a[i2] = append(a[i2], b[j:]...)
					break
				}
				a[i2][j+1] += c
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1213D2(os.Stdin, os.Stdout) }
