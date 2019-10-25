package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1249B2(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var q int
	for Fscan(in, &q); q > 0; q-- {
		var n int
		Fscan(in, &n)
		p := make([]int, n)
		for i := range p {
			Fscan(in, &p[i])
			p[i]--
		}
		ans := make([]interface{}, n)
		vis := make([]bool, n)
		for _, pi := range p {
			if vis[pi] {
				continue
			}
			cycle := 0
			vals := []int{}
			for j := pi; !vis[j]; j = p[j] {
				vis[j] = true
				cycle++
				vals = append(vals, j)
			}
			for _, val := range vals {
				ans[val] = cycle
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() {
//	Sol1249B2(os.Stdin, os.Stdout)
//}
