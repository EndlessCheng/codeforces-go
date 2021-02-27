package main

import (
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF758D(in io.Reader, out io.Writer) {
	var base, ans int64
	var k string
	Fscan(in, &base, &k)
	for pb := int64(1); k != ""; pb *= base {
		x := int64(0)
		// 判断 x+p10 < base 是为了防止 k=1000...0 导致乘法爆掉
		for i, p10 := len(k)-1, int64(1); i >= 0 && x+p10 < base && x+int64(k[i]&15)*p10 < base; i-- {
			if k[i] > '0' {
				x += int64(k[i]&15) * p10
				k = k[:i]
			}
			p10 *= 10
		}
		if x > 0 {
			ans += x * pb
		} else {
			k = k[:len(k)-1]
		}
	}
	Fprint(out, ans)
}

//func main() { CF758D(os.Stdin, os.Stdout) }
