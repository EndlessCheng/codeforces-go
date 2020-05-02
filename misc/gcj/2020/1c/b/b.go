package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

type pair struct{ cnt, i int }
type pairs []pair

// GCJ Golang 版本过低没有 sort.Slice
func (p pairs) Len() int           { return len(p) }
func (p pairs) Less(i, j int) bool { return p[i].cnt > p[j].cnt }
func (p pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) [10]byte {
		var ignore, s []byte
		Fscan(in, &ignore)
		heads := make(pairs, 26)
		for i := range heads {
			heads[i].i = i
		}
		vis := [26]bool{}
		for i := 0; i < 1e4; i++ {
			Fscan(in, &ignore, &s)
			heads[s[0]-'A'].cnt++
			for _, b := range s {
				vis[b-'A'] = true
			}
		}
		sort.Sort(heads)
		ans := [10]byte{}
		for i, p := range heads[:9] {
			ans[i+1] = byte('A' + p.i)
			vis[p.i] = false
		}
		for i, v := range vis {
			if v {
				ans[0] = byte('A' + i)
				break
			}
		}
		return ans
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: %s\n", _case, solve(_case))
	}
}

func main() { run(os.Stdin, os.Stdout) }
