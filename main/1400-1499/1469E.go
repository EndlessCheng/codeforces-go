package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func CF1469E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s string
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		if n == 1 {
			Fprintln(out, "YES")
			Fprintln(out, s[0]&1)
			continue
		}
		k2 := bits.Len(uint(n)) - 1
		left := k - k2
		if left <= 0 {
			has := make([]bool, 1<<k)
			mask := 1<<(k-1) - 1
			v, _ := strconv.ParseUint(s[:k-1], 2, 64)
			x := int(v)
			for _, c := range s[k-1:] {
				x = x<<1 | int(c&1)
				has[x] = true
				x &= mask
			}
			for i := 1<<k - 1; i >= 0; i-- {
				if !has[i] {
					Fprintf(out, "YES\n%0*b\n", k, 1<<k-1^i)
					continue o
				}
			}
			Fprintln(out, "NO")
		} else {
			has := make([]bool, 1<<k2)
			mask := 1<<(k2-1) - 1
			c1 := strings.Count(s[:left-1], "1")
			v, _ := strconv.ParseUint(s[left:left+k2-1], 2, 64)
			x := int(v)
			for i := left - 1; i+k2 < n; i++ {
				c1 += int(s[i] & 1)
				x = x<<1 | int(s[i+k2]&1)
				if c1 == left {
					has[x] = true
				}
				x &= mask
				c1 -= int(s[i-left+1] & 1)
			}
			for i := 1<<k2 - 1; ; i-- {
				if !has[i] {
					Fprintf(out, "YES\n%s%0*b\n", strings.Repeat("0", left), k2, 1<<k2-1^i)
					break
				}
			}
		}
	}
}

//func main() { CF1469E(os.Stdin, os.Stdout) }
