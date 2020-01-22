package main

import (
	"bytes"
	. "fmt"
	"io/ioutil"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF1005D() {
	s, _ := ioutil.ReadAll(os.Stdin)
	s = bytes.TrimSpace(s)
	n := len(s)
	for i, c := range s {
		s[i] = (c - '0') % 3
	}
	ans := 0
	for i := 0; i < n; i++ {
		if c := s[i]; c == 0 {
			ans++
		} else if i+1 < n {
			if c+s[i+1] == 3 {
				ans++
				i++
			} else if c == s[i+1] && i+2 < n {
				ans++
				i += 2
			}
		}
	}
	Print(ans)
}
