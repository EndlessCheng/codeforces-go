package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf444C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	var n, m, op, l, r, v int
	Fscan(in, &n, &m)
	a := make([]int, n)
	c := make([]int, n)
	for i := range c {
		c[i] = i + 1
	}

	B := int(math.Sqrt(float64(n)))
	type block struct{ l, r, sum, c, todo int }
	bs := make([]block, (n-1)/B+1)
	for i := 0; i < n; i += B {
		bs[i/B] = block{l: i, r: min(i+B, n)}
	}
	spread := func(b *block) {
		if b.todo > 0 {
			for j := b.l; j < b.r; j++ {
				a[j] += b.todo
				c[j] = b.c
			}
		}
		b.todo = 0
		b.c = 0
	}
	update := func(l, r, v int) (s int) {
		for j := l; j < r; j++ {
			d := abs(v - c[j])
			c[j] = v
			a[j] += d
			s += d
		}
		return
	}

	for range m {
		Fscan(in, &op, &l, &r)
		l--
		if op == 1 {
			Fscan(in, &v)
			for i := range bs {
				b := &bs[i]
				if b.r <= l {
					continue
				}
				if b.l >= r {
					break
				}
				if l <= b.l && b.r <= r {
					if b.c > 0 {
						d := abs(v - b.c)
						b.sum += d * (b.r - b.l)
						b.todo += d
					} else {
						b.sum += update(b.l, b.r, v)
					}
					b.c = v
				} else {
					spread(b)
					b.sum += update(max(b.l, l), min(b.r, r), v)
				}
			}
		} else {
			ans := 0
			for i := range bs {
				b := &bs[i]
				if b.r <= l {
					continue
				}
				if b.l >= r {
					break
				}
				if l <= b.l && b.r <= r {
					ans += b.sum
				} else {
					spread(b)
					for j := max(b.l, l); j < min(b.r, r); j++ {
						ans += a[j]
					}
				}
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { cf444C(bufio.NewReader(os.Stdin), os.Stdout) }
