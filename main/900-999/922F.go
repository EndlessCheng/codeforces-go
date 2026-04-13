package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf922F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, k int
	Fscan(in, &n, &k)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 2 * i; j <= n; j += i {
			d[j]++
		}
	}

	has := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		k -= d[i]
		has[i] = true
		if k <= 0 {
			n = i
			break
		}
	}
	if k > 0 {
		Fprint(out, "No")
		return
	}

	k = -k
	for i := 1; i <= n; i++ {
		if d[i] == 1 && k >= n/i {
			k -= n / i
			has[i] = false
		}
	}

	cnt := 0
	for i := 1; i <= n; i++ {
		if has[i] {
			cnt++
		}
	}

	Fprintln(out, "Yes")
	Fprintln(out, cnt)
	for i := 1; i <= n; i++ {
		if has[i] {
			Fprint(out, i, " ")
		}
	}
}

//func main() { cf922F(bufio.NewReader(os.Stdin), os.Stdout) }
