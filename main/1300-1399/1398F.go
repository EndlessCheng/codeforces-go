package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1398F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var s string
	Fscan(in, &n, &s)
	pre0 := make([]int, n)
	pre1 := make([]int, n)
	p0, p1 := -1, -1
	for i, b := range s {
		if b == '0' {
			p0 = i
		} else if b == '1' {
			p1 = i
		}
		pre0[i] = p0
		pre1[i] = p1
	}

	for x := 1; x <= n; x++ {
		ans := 0
		for i := 0; i+x <= n; {
			j := min(pre0[i+x-1], pre1[i+x-1])
			if j < i {
				ans++
				i += x
			} else {
				i = j + 1
			}
		}
		Fprint(out, ans, " ")
	}
}

//func main() { cf1398F(bufio.NewReader(os.Stdin), os.Stdout) }
