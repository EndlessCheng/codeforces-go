package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1196B(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		odds := []int{}
		for i := range n {
			Fscan(in, &v)
			if v&1 > 0 {
				odds = append(odds, i)
			}
		}
		if len(odds) < k || len(odds)&1 != k&1 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
			for _, i := range odds[:k-1] {
				Fprint(out, i+1, " ")
			}
			Fprintln(out, n)
		}
	}
}

//func main() { cf1196B(bufio.NewReader(os.Stdin), os.Stdout) }
