package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strconv"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, ans int
	var s string
	Fscan(in, &n, &m)
	type pair struct{ len, x int }
	cnt := map[pair]int{}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &s)
		a[i], _ = strconv.Atoi(s)
		cnt[pair{len(s), a[i] % m}]++
	}

	for _, v := range a {
		l := 1
		for p10 := uint(10); p10 <= 1e10; p10 *= 10 {
			r := int(uint(v) * p10 % uint(m))
			ans += cnt[pair{l, (m - r) % m}]
			l++
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
