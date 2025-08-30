package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf628E(in io.Reader, out io.Writer) {
	var n, m, ans int
	var s string
	Fscan(in, &n, &m)
	emptyL := make([][]int, n)
	emptyR := make([][]int, n)
	for i := range emptyL {
		Fscan(in, &s)
		emptyL[i] = make([]int, m)
		p := -1
		for j, b := range s {
			if b == '.' {
				p = j
			}
			emptyL[i][j] = p
		}

		emptyR[i] = make([]int, m)
		p = m
		for j := m - 1; j >= 0; j-- {
			if s[j] == '.' {
				p = j
			}
			emptyR[i][j] = p
		}
	}

	t := make([]int, m+1)
	time := make([]int, m+1)
	now := 0
	add := func(i int) {
		for ; i <= m; i += i & -i {
			if time[i] < now {
				time[i] = now
				t[i] = 0
			}
			t[i]++
		}
	}
	del := func(i int) {
		for ; i <= m; i += i & -i {
			t[i]--
		}
	}
	pre := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			if time[i] == now {
				res += t[i]
			}
		}
		return
	}

	delTodo := make([][]int, m)
	keys := []int{}
	clear := func() {
		now++
		for _, l := range keys {
			delTodo[l] = delTodo[l][:0]
		}
		keys = keys[:0]
	}

	for K := range n + m - 1 {
		clear()
		for j := min(K, m-1); j >= max(K-n+1, 0); j-- {
			i := K - j
			r := emptyR[i][j]
			if r == j {
				clear()
				continue
			}
			add(j + 1)
			ans += pre(r)
			l := emptyL[i][j] + 1
			delTodo[l] = append(delTodo[l], j+1)
			keys = append(keys, l)
			for _, j := range delTodo[j] {
				del(j)
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf628E(bufio.NewReader(os.Stdin), os.Stdout) }
