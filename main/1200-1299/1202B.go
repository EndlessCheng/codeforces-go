package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"io/ioutil"
)

// github.com/EndlessCheng/codeforces-go
func CF1202B(_r io.Reader, _w io.Writer) {
	s, _ := ioutil.ReadAll(_r)
	s = bytes.TrimSpace(s)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	cnts := [10]int{}
	for i, c := range s[1:] {
		cnts[(c-s[i]+10)%10]++
	}
	ans := [10][10]int{}
	for x := range ans {
	loop:
		for y := range ans {
			ops := [10]int{}
			for i := range ops {
				ops[i] = 99
			}
			for i := 0; i < 10; i++ {
				for j := 0; j < 10; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if v := (i*x + j*y) % 10; i+j < ops[v] {
						ops[v] = i + j
					}
				}
			}
			for i, op := range ops {
				if op == 99 && cnts[i] > 0 {
					ans[x][y] = -1
					continue loop
				}
				ans[x][y] += (op - 1) * cnts[i]
			}
		}
	}
	for _, ai := range ans {
		for _, v := range ai {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() {
//	CF1202B(os.Stdin, os.Stdout)
//}
