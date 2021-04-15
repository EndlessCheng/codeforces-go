package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const sz = 5
	const mask = 1<<sz - 1

	T, a := 0, [sz + 2]int{}
	for Fscanln(in, &T); T > 0; T-- {
		for i := 1; i <= sz; i++ {
			Fscanf(in, "%b\n", &a[i])
		}
		ans := 7
	o:
		for a[0] = 0; a[0] < 1<<sz; a[0]++ {
			a, c := a, 0
			for i := 1; i <= sz; i++ {
				for s := mask &^ uint(a[i-1]); s > 0; s &= s - 1 {
					p := bits.TrailingZeros(s)
					a[i] ^= 3 << p
					if p > 0 {
						a[i] ^= 1 << (p - 1)
					}
					a[i+1] ^= 1 << p
					c++
				}
			}
			if a[sz]&mask != mask {
				continue o
			}
			ans = min(ans, c)
		}
		if ans == 7 {
			ans = -1
		}
		Fprintln(out, ans)
		Fscanf(in, "\n")
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
