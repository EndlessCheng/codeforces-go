package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF600C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(bufio.NewReader(in), &s)
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	ps := []int{}
	for i, c := range cnt {
		if c&1 > 0 {
			ps = append(ps, i)
		}
	}
	for i, j := 0, len(ps)-1; i < j; i++ {
		cnt[ps[i]]++
		cnt[ps[j]]--
		j--
	}
	for i, c := range cnt {
		Fprint(out, strings.Repeat(string(byte('a'+i)), c/2))
	}
	for i, c := range cnt {
		if c&1 > 0 {
			Fprintf(out, "%c", 'a'+i)
		}
	}
	for i := len(cnt) - 1; i >= 0; i-- {
		Fprint(out, strings.Repeat(string(byte('a'+i)), cnt[i]/2))
	}
}

//func main() { CF600C(os.Stdin, os.Stdout) }
