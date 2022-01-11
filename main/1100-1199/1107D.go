package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1107D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const w = bits.UintSize
	var n, ans int
	var s string
	Fscan(in, &n)
	a := make([][]uint, n)
	for i := range a {
		a[i] = make([]uint, (n+w-1)/w)
		Fscan(in, &s)
		for j, b := range s {
			if b < 'A' {
				b -= '0'
			} else {
				b -= 'A' - 10
			}
			a[i][j*4/w] |= uint(bits.Reverse8(uint8(b))>>4) << (j % (w / 4) * 4)
		}
	}

	f := func(d int) bool {
		for i, r := range a {
			for j := 0; j < n; j += d {
				st, end := j, j+d-1
				if a[i-i%d][j/w]>>(j%w)&1 == 0 {
					i := st / w
					if i == end/w {
						if r[i]>>(st%w)&(1<<(end-st+1)-1) != 0 {
							return false
						}
						continue
					}
					if r[i]>>(st%w) != 0 {
						return false
					}
					for i++; i < end/w; i++ {
						if r[i] != 0 {
							return false
						}
					}
					if r[end/w]&(1<<(end%w+1)-1) != 0 {
						return false
					}
				} else {
					i := st / w
					if i == end/w {
						if r[i]>>(st%w)&(1<<(end-st+1)-1) != 1<<(end-st+1)-1 {
							return false
						}
						continue
					}
					if ^(r[i] | (1<<(st%w) - 1)) != 0 {
						return false
					}
					for i++; i < end/w; i++ {
						if ^r[i] != 0 {
							return false
						}
					}
					if r[end/w]&(1<<(end%w+1)-1) != 1<<(end%w+1)-1 {
						return false
					}
				}
			}
		}
		return true
	}
	for d := 1; d*d <= n; d++ {
		if n%d == 0 {
			if d*d < n && f(n/d) {
				ans = n / d
				break
			}
			if f(d) {
				ans = d
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1107D(os.Stdin, os.Stdout) }
