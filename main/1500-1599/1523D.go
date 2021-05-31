package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"math/rand"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func CF1523D(_r io.Reader, out io.Writer) {
	t0 := time.Now()
	rand.Seed(t0.UnixNano())
	in := bufio.NewReader(_r)
	var n, m, p, mx int
	var ans uint64
	Fscanln(in, &n, &m, &p)
	low := (n + 1) / 2
	a := make([]uint64, n)
	for i := range a {
		Fscanf(in, "%b\n", &a[i])
	}
	for time.Since(t0) < 2*time.Second {
		// 由于答案是某个元素的子集，且该子集至少出现在一半的元素中，则可以任取一元素 mask，取不到某个 a[i] 的概率是 0.5^loop，非常小
		// 然后统计 mask 的所有子集在 a 中的出现次数
		// 这里的技巧是，统计 a[i]&mask 的出现次数，然后再据此统计 a[i]&mask 的子集的出现次数
		mask := a[rand.Intn(n)]
		if bits.OnesCount64(mask) <= mx {
			continue
		}
		cnt := map[uint64]int{}
		for _, v := range a {
			cnt[v&mask]++
		}
		cnt2 := map[uint64]int{}
		for v, c := range cnt {
			for s := v; s > 0; s = (s - 1) & v {
				cnt2[s] += c
			}
		}
		for v, c := range cnt2 {
			if c >= low {
				if o := bits.OnesCount64(v); o > mx {
					ans, mx = v, o
				}
			}
		}
	}
	Fprintf(out, "%0*b", m, ans)
}

//func main() { CF1523D(os.Stdin, os.Stdout) }
