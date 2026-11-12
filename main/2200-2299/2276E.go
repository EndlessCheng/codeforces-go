package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2276E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		diff := make([]int, n+1)
		diff2 := make([]int, n+1)
		last2, last20, last202, last2025, last2026 := -1, -1, -1, -1, -1
		for i, b := range s {
			if b == '2' {
				last202 = last20
				last2 = i
			} else if b == '0' {
				last20 = last2
			} else if b == '5' {
				last2025 = last202
			} else if b == '6' {
				last2026 = last202
			}
			l, r := last2026+1, last2025
			if l <= r {
				diff[l]++
				diff[i+1] -= r - l + 1
				diff2[l+1]++
				diff2[r+1]--
			}
		}

		sd2, sd := 0, 0
		for i, d := range diff[:n] {
			sd2 += diff2[i]
			sd += d + sd2
			Fprint(out, (i+1)*(n-i)-sd, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2276E(bufio.NewReader(os.Stdin), os.Stdout) }
