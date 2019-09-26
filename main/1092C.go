package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

func Sol1092C(reader io.Reader, writer io.Writer) {
	type pair struct {
		s   string
		idx int
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	n = 2*n - 2
	arr := make([]pair, n)
	for i := range arr {
		var s string
		Fscan(in, &s)
		arr[i] = pair{s, i}
	}
	sort.Slice(arr, func(i, j int) bool { return len(arr[i].s) < len(arr[j].s) })

	p, s := arr[n-2].s, arr[n-1].s
outer:
	for _, comb := range []string{p + s[len(s)-1:], s + p[len(p)-1:]} {
		ans := make([]byte, n)
		for i := 0; i < n; i += 2 {
			p0, p1 := arr[i], arr[i+1]
			if strings.HasPrefix(comb, p0.s) && strings.HasSuffix(comb, p1.s) {
				ans[p0.idx], ans[p1.idx] = 'P', 'S'
			} else if strings.HasSuffix(comb, p0.s) && strings.HasPrefix(comb, p1.s) {
				ans[p0.idx], ans[p1.idx] = 'S', 'P'
			} else {
				continue outer
			}
		}
		Fprint(out, string(ans))
		return
	}
}

//func main() {
//	Sol1092C(os.Stdin, os.Stdout)
//}
