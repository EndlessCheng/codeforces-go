package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1739E(in io.Reader, out io.Writer) {
	var n int
	a := [2]string{}
	Fscan(in, &n, &a[0], &a[1])
	f := make([][2]int, n+1)
	for j := n - 2; j >= 0; j-- {
		for i := range 2 {
			f[j][i] = f[j+1][i] + int(a[i^1][j]&1)
			if a[i^1][j] == '1' {
				f[j][i] = min(f[j][i], f[j+2][i^1]+int(a[i][j+1]&1))
			}
		}
	}
	Fprint(out, strings.Count(a[0], "1")+strings.Count(a[1], "1")-f[0][0])
}

//func main() { cf1739E(bufio.NewReader(os.Stdin), os.Stdout) }
