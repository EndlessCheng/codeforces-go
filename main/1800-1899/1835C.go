package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1835C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		half := 1 << k // 值域的一半
		s := make([]int, half*2+1)
		pos := make([]int, half)
		for i := range pos {
			pos[i] = -1
		}
		pos[0] = 0
		pos2 := make([][2]int, half)
		for i := 1; ; i++ {
			Fscan(in, &s[i])
			s[i] ^= s[i-1]
			si := s[i]
			j := pos[si&(half-1)]
			if j >= 0 {
				s2 := (si ^ s[j]) >> k
				if pos2[s2][1] > 0 {
					a := append(pos2[s2][:], j)
					slices.Sort(a)
					Fprintln(out, a[0]+1, a[1], a[2]+1, i)
					for range half*2 - i {
						Fscan(in, &k)
					}
					break
				}
				pos2[s2] = [2]int{j, i}
			}
			pos[si&(half-1)] = i
		}
	}
}

//func main() { cf1835C(bufio.NewReader(os.Stdin), os.Stdout) }
