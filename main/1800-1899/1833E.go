package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1833E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		to := make([]int, n)
		for i := range to {
			Fscan(in, &to[i])
			to[i]--
		}
		c2, c3 := 0, 0
		time := make([]int, len(to))
		clock := 1
		for x, t := range time {
			if t > 0 {
				continue
			}
			for t0 := clock; x >= 0; x = to[x] {
				if time[x] > 0 {
					if time[x] >= t0 {
						sz := clock - time[x]
						if sz == 2 {
							c2++
						} else {
							c3++
						}
					}
					break
				}
				time[x] = clock
				clock++
			}
		}
		ans := c3
		if c2 > 0 {
			ans++
		}
		Fprintln(out, ans, c2+c3)
	}
}

//func main() { CF1833E(os.Stdin, os.Stdout) }
