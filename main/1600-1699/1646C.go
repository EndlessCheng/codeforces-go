package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF1646C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	f := []uint64{2}
	for i := uint64(3); f[len(f)-1]*i <= 1e12; i++ {
		f = append(f, f[len(f)-1]*i)
	}

	var T, n uint64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		ans := 99
		for i := 0; i < 1<<len(f); i++ {
			s, c := n, 0
			for j, v := range f {
				if i>>j&1 > 0 {
					s -= v
					c++
				}
			}
			if s >= 0 {
				c += bits.OnesCount64(s)
				if c < ans {
					ans = c
				}
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1646C(os.Stdin, os.Stdout) }
