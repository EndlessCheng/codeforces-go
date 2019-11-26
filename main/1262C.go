package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func Sol1262C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	reverse := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	type pair struct{ l, r int }

	var t, n, k int
	var s []byte
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n, &k, &s)
		ops := []pair{}
		targetStr := strings.Repeat("()", k-1) + strings.Repeat("(", n/2-k+1) + strings.Repeat(")", n/2-k+1)
		for i, c := range s {
			if c != targetStr[i] {
				j := i + 1
				for ; s[j] != targetStr[i]; j++ {
				}
				reverse(s[i : j+1])
				ops = append(ops, pair{i + 1, j + 1})
			}
		}
		Fprintln(out, len(ops))
		for _, op := range ops {
			Fprintln(out, op.l, op.r)
		}
	}
}

//func main() {
//	Sol1262C(os.Stdin, os.Stdout)
//}
