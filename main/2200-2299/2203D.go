package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2203D(in io.Reader, out io.Writer) {
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		mx := n + m + 1
		cnt := make([]int, mx)
		for range n {
			Fscan(in, &v)
			cnt[v]++
		}

		div := make([]int, mx)
		for i, c := range cnt {
			if c > 0 {
				for j := i; j < mx; j += i {
					div[j] += c
				}
			}
		}

		alice, bob := 0, 0
		for range m {
			Fscan(in, &v)
			if div[v] == n {
				alice++
			} else if div[v] == 0 {
				bob++
			}
		}

		common := m - alice - bob
		alice += (common + 1) / 2
		if alice*2 > m {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

//func main() { cf2203D(bufio.NewReader(os.Stdin), os.Stdout) }
