package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf482A(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	if k == 1 {
		for i := 1; i <= n; i++ {
			Fprint(out, i, " ")
		}
		return
	}
	k--
	Fprint(out, "1 ", n)
	i := 2
	for k > 2 {
		Fprint(out, " ", i, " ", n+1-i)
		k -= 2
		i++
	}
	if k == 1 {
		for j := n + 1 - i; j >= i; j-- {
			Fprint(out, " ", j)
		}
	} else {
		for j := i; j <= n+1-i; j++ {
			Fprint(out, " ", j)
		}
	}
}

//func main() { cf482A(bufio.NewReader(os.Stdin), os.Stdout) }
