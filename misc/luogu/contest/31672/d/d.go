package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	buf := make([]byte, 4096)
	_i := len(buf)
	rc := func() byte {
		if _i == len(buf) {
			_r.Read(buf)
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int32) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int32(b&15)
		}
		return
	}
	min := func(a, b int32) int32 {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int32) int32 {
		if a > b {
			return a
		}
		return b
	}

	n, q := r(), r()
	if n <= 1000 && q <= 1000 {
		a := [1001]int{}
		for i := int32(1); i <= n; i++ {
			a[i] = int(r())
		}
	o:
		for ; q > 0; q-- {
			if r() == 0 {
				a[r()] = int(r())
			} else {
				b := a
				l, R := r(), r()+1
				sort.Ints(b[l:R])
				bb := b[l:R]
				c := a
				l, R = r(), r()+1
				sort.Ints(c[l:R])
				cc := c[l:R]
				for i, v := range bb {
					if cc[i]-v != cc[0]-bb[0] {
						Fprintln(out, "NO")
						continue o
					}
				}
				Fprintln(out, "YES")
			}
		}
		return
	}

	a := make([]int32, n)
	small := true
	for i := range a {
		a[i] = r()
		if a[i] > 100 {
			small = false
		}
	}
	if small {
		type block struct {
			l, r int32
			cnt  [101]int32
		}
		blockSize := int32(math.Sqrt(float64(n))) * 10
		blockNum := (n-1)/blockSize + 1
		blocks := make([]block, blockNum)
		for _i, v := range a {
			i := int32(_i)
			j := i / blockSize
			if i%blockSize == 0 {
				blocks[j] = block{l: i, r: i - 1}
			}
			blocks[j].cnt[v]++
			blocks[j].r++
		}
		query := func(l, r int32) (cnt [101]int32) {
			for i := range blocks {
				b := &blocks[i]
				if b.r < l {
					continue
				}
				if b.l > r {
					break
				}
				if l <= b.l && b.r <= r {
					for j, c := range b.cnt {
						cnt[j] += c
					}
				} else {
					bl := max(b.l, l)
					br := min(b.r, r)
					for _, v := range a[bl : br+1] {
						cnt[v]++
					}
				}
			}
			return
		}

		for ; q > 0; q-- {
			if r() == 0 {
				i, val := r()-1, r()
				j := i / blockSize
				blocks[j].cnt[a[i]]--
				blocks[j].cnt[val]++
				a[i] = val
			} else {
				c1 := query(r()-1, r()-1)
				c2 := query(r()-1, r()-1)
			o2:
				for i, v := range c1 {
					if v > 0 {
						for j, w := range c2 {
							if w > 0 {
								for k := 0; i+k <= 100 && j+k <= 100; k++ {
									if c1[i+k] != c2[j+k] {
										Fprintln(out, "NO")
										break o2
									}
								}
								Fprintln(out, "YES")
								break o2
							}
						}
					}
				}
			}
		}
		return
	}
}

func main() { run(os.Stdin, os.Stdout) }
