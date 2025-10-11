package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1924C(in io.Reader, out io.Writer) {
	const M = 999999893
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % M
			}
			x = x * x % M
		}
		return res
	}
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := 1 - pow(2, n/2)
		b := pow(2, (n-1)/2*2+1)
		inv := pow((a*a-b)%M, M-2)
		Fprintln(out, (a*inv%M+M)%M)
	}
}

//func main() { cf1924C(bufio.NewReader(os.Stdin), os.Stdout) }
