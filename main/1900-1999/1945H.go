package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
func cf1945H(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n, x int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &x)
		a := make([]int, n)
		C := [19]int{}
		mx, and := 0, -1
		for i := range a {
			Fscan(in, &a[i])
			mx = max(mx, a[i])
			and &= a[i]
			for j := uint(a[i]); j > 0; j &= j - 1 {
				C[bits.TrailingZeros(j)]++
			}
		}

		sp := make([]bool, n)
		bitW := bits.Len(uint(mx))
		u := 1<<bitW - 1
		for i, v := range a {
			for j := uint(v ^ u); j > 0; j &= j - 1 {
				p := bits.TrailingZeros(j)
				if C[p] == n-1 || C[p] == n-2 {
					sp[i] = true // 含 1 或 2 个 0
					break
				}
			}
		}

		pr := func(i, j int) {
			Fprintln(out, "YES")
			Fprintln(out, 2, a[i], a[j])
			Fprint(out, n-2)
			for k, v := range a {
				if k != i && k != j {
					Fprint(out, " ", v)
				}
			}
			Fprintln(out)
		}

		pos := make([][]int, mx+1)
		cntBit := C[:bitW]
		for i, v := range a {
			if !sp[i] {
				pos[v] = append(pos[v], i)
				continue
			}
			// 枚举特殊数和 a 中另一个数
			for j, w := range a {
				if j == i {
					continue
				}
				g := gcd(v, w)
				// 特殊数会影响 and 值，单独计算
				and := 0
				for k, c := range cntBit {
					if c-v>>k&1-w>>k&1 == n-2 {
						and |= 1 << k
					}
				}
				if g > and+x {
					pr(i, j)
					continue o
				}
			}
		}

		// 枚举非特殊数：选两个不影响 and 值的数
		ans := []int{}
		for i := and + x + 1; i <= mx; i++ {
			ans = ans[:0]
			for j := i; j <= mx; j += i {
				ans = append(ans, pos[j]...)
				if len(ans) > 1 {
					pr(ans[0], ans[1])
					continue o
				}
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1945H(os.Stdin, os.Stdout) }

