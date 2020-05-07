package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run1(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) int {
		var n int
		Fscan(in, &n)
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, n)
			for j := range mat[i] {
				Fscan(in, &mat[i][j])
			}
		}

		check := func(s int) bool {
			for i := 0; i < n; i++ {
				mp := map[int]bool{}
				for j := 0; j < n; j++ {
					if s>>uint(n*i+j)&1 == 1 {
						if mp[mat[i][j]] {
							return false
						}
						mp[mat[i][j]] = true
					}
				}
			}
			for j := 0; j < n; j++ {
				mp := map[int]bool{}
				for i := 0; i < n; i++ {
					if s>>uint(n*i+j)&1 == 1 {
						if mp[mat[i][j]] {
							return false
						}
						mp[mat[i][j]] = true
					}
				}
			}
			return true
		}
		n2 := n * n
		for k := n2; k > 0; k-- {
			for sub := 1<<uint(k) - 1; sub < 1<<uint(n2); {
				if check(sub) {
					return n2 - k
				}
				x := sub & -sub
				y := sub + x
				sub = sub&^y/x>>1 | y
			}
		}
		panic(1)
	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: %d\n", _case, solve(_case))
	}
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(_case int) {
		var n int
		Fscan(in, &n)

	}

	var t int
	Fscan(in, &t)
	for _case := 1; _case <= t; _case++ {
		Fprintf(out, "Case #%d: ", _case)
		solve(_case)
	}
}

func main() { run(os.Stdin, os.Stdout) }
