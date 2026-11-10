package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf524F(in io.Reader, out io.Writer) {
	var str string
	Fscan(in, &str)
	n := len(str)
	c := make([]byte, n*2+1)
	copy(c[1:], str)

	t := 0
	for i := 1; i <= n; i++ {
		if c[i] == '(' {
			t++
		} else {
			t--
		}
	}

	for i := n + 1; i <= n*2; i++ {
		c[i] = c[i-n]
	}

	if t < 0 {
		Fprint(out, strings.Repeat("(", -t))
	}

	h := make([]int, n*2+1)
	for i := 1; i <= n*2; i++ {
		if c[i] == '(' {
			h[i] = h[i-1] + 1
		} else {
			h[i] = h[i-1] - 1
		}
	}

	q := make([]int, n*2+1)
	H, T := 1, 0
	for i := 1; i <= n; i++ {
		for T > 0 && h[i] < h[q[T]] {
			T--
		}
		T++
		q[T] = i
	}

	b := make([]bool, n*2+1)
	for i := 1; i <= n; i++ {
		if h[q[H]]-h[i-1] < min(t, 0) {
			b[i] = true
		}
		for T >= H && h[n+i] < h[q[T]] {
			T--
		}
		T++
		q[T] = n + i
		if q[H] == i {
			H++
		}
	}

	i, j := 1, 2
	for {
		for b[i] {
			i++
		}
		for b[j] {
			j++
		}
		if i == j {
			j++
			for b[j] {
				j++
			}
		}
		if i > n || j > n {
			break
		}
		k := 0
		for k < n && c[i+k] == c[j+k] {
			k++
		}
		if c[i+k] < c[j+k] {
			j = j + k + 1
		} else {
			i += k + 1
		}
	}

	i = min(i, j)
	Fprintf(out, "%s", c[i:n+i])
	if t > 0 {
		Fprint(out, strings.Repeat(")", t))
	}
}

//func main() { cf524F(bufio.NewReader(os.Stdin), os.Stdout) }
