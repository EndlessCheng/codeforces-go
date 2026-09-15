package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1713F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range 19 {
		for j := range n {
			if j&(1<<i) != 0 {
				a[j] ^= a[j^(1<<i)]
			}
		}
	}
	for i := range 19 {
		for j := range n {
			if j&(1<<i) != 0 {
				a[j^1<<i] ^= a[j]
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		Fprint(out, a[i], " ")
	}
}

//func main() { cf1713F(bufio.NewReader(os.Stdin), os.Stdout) }
