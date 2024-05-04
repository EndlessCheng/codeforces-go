package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf27E(in io.Reader, out io.Writer) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43}
	var tarD int
	Fscan(in, &tarD)
	ans := int(1e18)
	var f func(int, int, int, int)
	f = func(i, res, leftD, preE int) {
		if leftD == 1 {
			ans = min(ans, res)
			return
		}
		for e := 1; e <= preE && res <= ans/primes[i]; e++ {
			res *= primes[i]
			if leftD%(e+1) == 0 {
				f(i+1, res, leftD/(e+1), e)
			}
		}
	}
	f(0, 1, tarD, 99)
	Fprint(out, ans)
}

//func main() { cf27E(os.Stdin, os.Stdout) }
