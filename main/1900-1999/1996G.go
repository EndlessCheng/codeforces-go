package main

import (
	. "fmt"
	"io"
	"math/rand"
)

func cf1996G(in io.Reader, out io.Writer) {
	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		d := make([]int, n)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			k := rand.Int()
			d[v-1] ^= k
			d[w-1] ^= k
		}
		c := map[int]int{}
		ans, s := 0, 0
		for _, v := range d {
			s ^= v
			c[s]++
			ans = max(ans, c[s])
		}
		Fprintln(out, n-ans)
	}
}

//func main() { cf1996G(bufio.NewReader(os.Stdin), os.Stdout) }
