package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1511G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q, x, y int
	Fscan(in, &n, &m)
	f := [21][]int{}
	for i := range f {
		f[i] = make([]int, m+1)
	}

	a := make([]int, m+1)
	for range n {
		Fscan(in, &x)
		a[x]++
	}
	for i := range m {
		a[i+1] += a[i]
	}

	mx := bits.Len(uint(m)) - 1
	for j, k := 1, 1; j <= mx; j++ {
		for i := 1; i+k*2-1 <= m; i++ {
			f[j][i] = f[j-1][i] ^ f[j-1][i+k] ^ ((a[i+k*2-1] - a[i+k-1]) & 1 * k)
		}
		k *= 2
	}

	Fscan(in, &q)
	for range q {
		Fscan(in, &x, &y)
		y -= x - 1
		sum := 0
		res := 0
		for i := mx; i >= 0; i-- {
			if y>>i&1 > 0 {
				res ^= f[i][x] ^ ((a[x+1<<i-1] - a[x-1]) & 1 * sum)
				sum += 1 << i
				x += 1 << i
			}
		}
		if res > 0 {
			Fprint(out, "A")
		} else {
			Fprint(out, "B")
		}
	}
}

//func main() { cf1511G(bufio.NewReader(os.Stdin), os.Stdout) }
