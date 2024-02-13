package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf897B(in io.Reader, out io.Writer) {
	var k, p, ans int
	Fscan(in, &k, &p)
	for base := 1; k > 0; base *= 10 {
		for i := base; i < base*10 && k > 0; i++ {
			x := i
			for t := i; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			ans += x
			k--
		}
	}
	Fprint(out, ans%p)
}

//func main() { cf897B(os.Stdin, os.Stdout) }
