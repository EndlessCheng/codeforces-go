package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k, s int
	Fscan(in, &n, &k)
	a := make([]int, n)
	id := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return a[id[i]] < a[id[j]] })

	for _, i := range id {
		d := a[i] - s
		if d > 0 {
			if k <= d*n {
				s += k / n // 剩下所有苹果，每个至少吃 k/n 个
				k %= n     // 有 k 个苹果要多吃一个
				break
			}
			s += d // 吃掉所有 >= s 的苹果
			k -= d * n
		}
		n--
	}

	for _, v := range a {
		if v <= s {
			v = 0
		} else {
			v -= s
			if k > 0 {
				v--
				k--
			}
		}
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
