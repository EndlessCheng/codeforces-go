package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1228D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	powP := make([]uint64, n+1)
	powP[0] = 1
	const prime uint64 = 1e8 + 7
	for i := 1; i <= n; i++ {
		powP[i] = powP[i-1] * prime
	}
	hash := make([]uint64, n+1)
	for ; m > 0; m-- {
		var x, y int
		Fscan(in, &x, &y)
		hash[x] += powP[y]
		hash[y] += powP[x]
	}

	mp := map[uint64]int{}
	idx := 0
	for i := 1; i <= n; i++ {
		if hash[i] == 0 {
			Fprint(out, -1)
			return
		}
		if mp[hash[i]] == 0 {
			idx++
			mp[hash[i]] = idx
		}
	}
	if idx != 3 {
		Fprint(out, -1)
		return
	}
	for _, h := range hash[1:] {
		Fprint(out, mp[h], " ")
	}
}

//func main() {
//	Sol1228D(os.Stdin, os.Stdout)
//}
