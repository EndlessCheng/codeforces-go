package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, x int
	Fscan(in, &n, &k)
	a := make([]int, k)
	miss := make([]int, 0, n-k)
	for i := range a {
		Fscan(in, &a[i])
		for x++; x < a[i]; x++ {
			miss = append(miss, x)
		}
	}
	for x++; x <= n; x++ {
		miss = append(miss, x)
	}

	j := 0
	for _, v := range a[:k-1] {
		Fprint(out, v, " ")
		if j < len(miss) && miss[j] < v {
			Fprint(out, miss[j], " ")
			j++
		}
	}
	i := len(miss) - 1
	for ; i >= j && miss[i] > a[k-1]; i-- {
		Fprint(out, miss[i], " ")
	}
	Fprint(out, a[k-1], " ")
	for ; i >= j; i-- {
		Fprint(out, miss[i], " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
