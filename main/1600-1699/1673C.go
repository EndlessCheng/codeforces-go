package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1673C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	pal := []int{}
	for i := 1; i < 400; i++ {
		p := i
		for x := i / 10; x > 0; x /= 10 {
			p = p*10 + x%10
		}
		pal = append(pal, p)
		if i < 100 {
			p := i
			for x := i; x > 0; x /= 10 {
				p = p*10 + x%10
			}
			pal = append(pal, p)
		}
	}

	f := [40001]int{}
	f[0] = 1
	for _, v := range pal {
		for j := v; j < len(f); j++ {
			f[j] = (f[j] + f[j-v]) % 1_000_000_007
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		Fprintln(out, f[n])
	}
}

//func main() { CF1673C(os.Stdin, os.Stdout) }
