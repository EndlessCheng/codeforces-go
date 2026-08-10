package main

import (
	. "fmt"
	"io"
	"strconv"
)

// https://github.com/EndlessCheng
func cf676E(in io.Reader, out io.Writer) {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	b := make([]bool, n+1)
	cnt := 0
	for i := 0; i <= n; i++ {
		var s string
		Fscan(in, &s)
		if s == "?" {
			b[i] = true
			cnt++
		} else {
			a[i], _ = strconv.Atoi(s)
		}
	}

	if k == 0 {
		if b[0] && (n+1-cnt)&1 > 0 || !b[0] && a[0] == 0 {
			Fprint(out, "Yes")
		} else {
			Fprint(out, "No")
		}
		return
	}
	if cnt > 0 {
		if n&1 > 0 {
			Fprint(out, "Yes")
		} else {
			Fprint(out, "No")
		}
		return
	}

	c := 0
	for i := n; i >= 0 && abs(c) <= 20000; i-- {
		a[i] += c
		c = a[i] * k
	}
	if c != 0 {
		Fprint(out, "No")
	} else {
		Fprint(out, "Yes")
	}
}

//func main() { cf676E(bufio.NewReader(os.Stdin), os.Stdout) }
