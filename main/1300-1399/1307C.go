package main

import (
	"bytes"
	. "fmt"
	"io"
	"io/ioutil"
)

// github.com/EndlessCheng/codeforces-go
func CF1307C(_r io.Reader, _w io.Writer) {
	s, _ := ioutil.ReadAll(_r)
	s = bytes.TrimSpace(s)
	cnt := [26]int64{}
	cntPair := [26][26]int64{}
	for _, b := range s {
		b -= 'a'
		for i, c := range cnt {
			cntPair[i][b] += c
		}
		cnt[b]++
	}
	ans := int64(0)
	for _, c := range cnt {
		if c > ans {
			ans = c
		}
	}
	for _, cp := range cntPair {
		for _, c := range cp {
			if c > ans {
				ans = c
			}
		}
	}
	Fprint(_w, ans)
}

//func main() { CF1307C(os.Stdin, os.Stdout) }
