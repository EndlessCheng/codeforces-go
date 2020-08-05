package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var b, a []byte
	Fscan(in, &b, &a)
	nb := len(b)
	id := make([]int, nb)
	for i := range id {
		Fscan(in, &id[i])
		id[i]--
	}
	Fprint(out, sort.Search(nb, func(p int) bool {
		used := make([]bool, nb)
		for _, i := range id[:p] {
			used[i] = true
		}
		c := []byte{}
		for i, u := range used {
			if !u {
				c = append(c, b[i])
			}
		}
		i, n := 0, len(a)
		j, m := 0, len(c)
		for {
			if i == n {
				return false
			}
			if j == m {
				return true
			}
			if a[i] == c[j] {
				i++
				j++
			} else {
				j++
			}
		}
	})-1)
}

func main() { run(os.Stdin, os.Stdout) }
