package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1874B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	pow5 := [8]int{1}
	for i := 1; i < 8; i++ {
		pow5[i] = pow5[i-1] * 5
	}

	st := 0
	for i, p5 := range pow5 {
		st += i & 3 * p5
	}

	f := make([]int, pow5[7]*5)
	for i := range f {
		f[i] = 1e9
	}
	f[st] = 0
	q := []int{st}
	for len(q) > 0 {
		mask := q[0]
		q = q[1:]
		for op := range 4 {
			newMask := 0
			for i, p5 := range pow5 {
				cd := mask / p5 % 5
				c, d := cd>>1, cd&1
				switch op {
				case 0:
					c &= d
				case 1:
					c |= d
				case 2:
					d ^= c
				default:
					d ^= i >> 2
				}
				newMask += (c<<1 | d) * p5
			}
			if f[newMask] == 1e9 {
				f[newMask] = f[mask] + 1
				q = append(q, newMask)
			}
		}
	}
	for mask := range f {
		for _, p5 := range pow5 {
			if mask/p5%5 == 4 {
				for i := 1; i < 5; i++ {
					f[mask] = min(f[mask], f[mask-i*p5])
				}
				break
			}
		}
	}

	var T, a, b, c, d, m int
	mp := [8]int{}
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &a, &b, &c, &d, &m)
		for i := range mp {
			mp[i] = 4
		}
		for i := range 30 {
			mab := m>>i&1<<2 | a>>i&1<<1 | b>>i&1
			cd := c>>i&1<<1 | d>>i&1
			if mp[mab] == 4 {
				mp[mab] = cd
			} else if mp[mab] != cd {
				Fprintln(out, -1)
				continue o
			}
		}
		mask := 0
		for i, cd := range mp {
			mask += cd * pow5[i]
		}
		if f[mask] == 1e9 {
			Fprintln(out, -1)
		} else {
			Fprintln(out, f[mask])
		}
	}
}

//func main() { cf1874B(bufio.NewReader(os.Stdin), os.Stdout) }
