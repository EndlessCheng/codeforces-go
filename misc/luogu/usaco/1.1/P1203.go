package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// O(n) DP 做法 https://tracyqwerty.blog.luogu.org/solution-p1203

// github.com/EndlessCheng/codeforces-go
func SolP1203(reader io.Reader, writer io.Writer) {
	maxs := func(vals ...int) int {
		ans := vals[0]
		for _, val := range vals[1:] {
			if val > ans {
				ans = val
			}
		}
		return ans
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)
	s += s

	calc := func(c1, c2 byte, i int) (cnt int) {
		for j := i - 1; j >= 0; j-- {
			if s[j] != c1 && s[j] != 'w' {
				break
			}
			cnt++
		}
		for j := i; j < 2*n; j++ {
			if s[j] != c2 && s[j] != 'w' {
				break
			}
			cnt++
		}
		return
	}

	ans := -1
	for i := 1; i < 2*n; i++ {
		c1, c2 := s[i-1], s[i]
		if c1 == c2 {
			continue
		}
		if c1 == 'w' {
			ans = maxs(ans, calc('r', c2, i), calc('b', c2, i))
		} else if c2 == 'w' {
			ans = maxs(ans, calc(c1, 'r', i), calc(c1, 'w', i))
		} else {
			ans = maxs(ans, calc(c1, c2, i))
		}
	}
	if ans == -1 || ans > n {
		ans = n
	}
	Fprintln(out, ans)
}

func main() {
	SolP1203(os.Stdin, os.Stdout)
}
