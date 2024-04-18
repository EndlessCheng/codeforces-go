package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf940B(in io.Reader, out io.Writer) {
	var n, k, a, b, ans int
	Fscan(in, &n, &k, &a, &b)
	for n > 1 {
		ans += n % k * a
		n -= n % k
		if (n-n/k)*a <= b {
			ans += (n - 1) * a
			break
		}
		ans += b
		n /= k
	}
	Fprint(out, ans)
}

//func main() { cf940B(os.Stdin, os.Stdout) }
