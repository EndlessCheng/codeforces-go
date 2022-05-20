package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF747D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans int
	var t string
	Fscan(in, &n, &k)
	gap := []int{}
	cnt := int(-1e9)
	for pre := true; n > 0; n-- {
		Fscan(in, &t)
		s := t[0] != '-'
		if s != pre {
			ans++
			if s {
				cnt = 0
			} else if cnt > 0 {
				gap = append(gap, cnt)
			}
		}
		if s {
			cnt++
		} else if k--; k < 0 {
			Fprint(out, -1)
			return
		}
		pre = s
	}

	sort.Ints(gap)
	for _, v := range gap {
		if k < v {
			break
		}
		k -= v
		ans -= 2
	}
	if 0 < cnt && cnt <= k {
		ans--
	}
	Fprint(out, ans)
}

//func main() { CF747D(os.Stdin, os.Stdout) }
