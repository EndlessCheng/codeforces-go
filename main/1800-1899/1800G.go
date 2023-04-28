package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"reflect"
	"runtime/debug"
	"sort"
	"unsafe"
)

// https://space.bilibili.com/206214
func init() { debug.SetMemoryLimit(200 << 20) }

func CF1800G(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		tid := map[string]int{}
		sym := make([]bool, n)
		var f func(int, int) int
		f = func(v, fa int) int {
			ids := make([]int, 0, len(g[v]))
			for _, w := range g[v] {
				if w != fa {
					ids = append(ids, f(w, v))
				}
			}
			sort.Ints(ids)
			sh := (*reflect.SliceHeader)(unsafe.Pointer(&ids))
			sh.Len *= bits.UintSize / 8
			s := *(*string)(unsafe.Pointer(sh))
			id, ok := tid[s]
			if !ok {
				ok := true
				cnt := map[int]int{}
				sh.Len /= bits.UintSize / 8
				for _, i := range ids {
					cnt[i]++
				}
				left := len(ids) % 2
				for i, c := range cnt {
					if c%2 > 0 {
						if left == 0 || !sym[i] {
							ok = false
							break
						}
						left--
					}
				}
				sh.Len *= bits.UintSize / 8
				id = len(tid)
				tid[s] = id
				sym[id] = ok
			}
			return id
		}
		if sym[f(0, -1)] {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1800G(os.Stdin, os.Stdout) }
