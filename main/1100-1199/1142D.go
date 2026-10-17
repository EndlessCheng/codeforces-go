package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1142D(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	f := make([][11]int, n+1)
	ans := 0
	for i := n - 1; i >= 0; i-- {
		d := int(s[i] - '0')
		for j := d + 1; j <= 10; j++ {
			f[i][j] = f[i+1][(j*(j-1)/2+d+10)%11] + 1
		}
		if d > 0 {
			ans += f[i+1][d] + 1
		}
	}
	Fprint(out, ans)
}

//func main() { cf1142D(bufio.NewReader(os.Stdin), os.Stdout) }
