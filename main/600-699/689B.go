package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf689B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, w int
	Fscan(in, &n)
	from := make([][]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &w)
		from[w-1] = append(from[w-1], i)
	}

	f := make([]int, n)
	Fprint(out, 0)
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + 1
		for _, v := range from[i] {
			f[i] = min(f[i], f[v]+1)
		}
		Fprint(out, " ", f[i])
	}
}

//func main() { cf689B(bufio.NewReader(os.Stdin), os.Stdout) }
