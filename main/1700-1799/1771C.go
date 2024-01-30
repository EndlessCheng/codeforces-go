package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1771C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx = 31623
	primes := []int32{}
	np := [mx]bool{}
	for i := int32(2); i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int32, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		vis := map[int32]bool{}
		for _, x := range a {
			for _, p := range primes {
				if x%p > 0 {
					continue
				}
				if vis[p] {
					Fprintln(out, "YES")
					continue o
				}
				vis[p] = true
				for x /= p; x%p == 0; x /= p {
				}
			}
			if x > 1 {
				if vis[x] {
					Fprintln(out, "YES")
					continue o
				}
				vis[x] = true
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf1771C(os.Stdin, os.Stdout) }
