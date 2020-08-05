package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol913D(reader io.Reader, writer io.Writer) {
	type pair struct {
		a, t, idx int
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, t int
	Fscan(in, &n, &t)
	arr := make([]pair, n)
	for i := range arr {
		Fscan(in, &arr[i].a, &arr[i].t)
		arr[i].idx = i + 1
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].t < arr[j].t })
	var idxAns []interface{}
	f := func(target int) bool {
		if target == 0 {
			return !true
		}
		idxAns = []interface{}{}
		spent := 0
		for _, p := range arr {
			if p.a >= target {
				if spent+p.t > t {
					break
				}
				spent += p.t
				idxAns = append(idxAns, p.idx)
				if len(idxAns) == target {
					break
				}
			}
		}
		return !(len(idxAns) == target)
	}
	ans := sort.Search(n+1, f) - 1
	f(ans)
	Fprintln(out, ans)
	Fprintln(out, ans)
	Fprintln(out, idxAns...)
}

//func main() {
//	Sol913D(os.Stdin, os.Stdout)
//}
