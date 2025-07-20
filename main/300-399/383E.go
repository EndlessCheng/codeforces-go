package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf383E(in io.Reader, out io.Writer) {
	const mx = 24
	f := [1 << mx]uint16{}
	n, s := uint16(0), ""
	Fscan(in, &n)
	for range n {
		Fscan(in, &s)
		f[1<<(s[0]-'a')|1<<(s[1]-'a')|1<<(s[2]-'a')]++
	}

	for i := range mx {
		for j := 0; j < 1<<mx; j++ {
			j |= 1 << i
			f[j] += f[j^1<<i]
		}
	}

	ans := 0
	for _, v := range f {
		x := int(n - v)
		ans ^= x * x
	}
	Fprint(out, ans)
}

//func main() { cf383E(bufio.NewReader(os.Stdin), os.Stdout) }
