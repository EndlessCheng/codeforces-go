package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
var cnt70 [450 * 2e5]int32

func cf1270F(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n := len(s)
	B := int(math.Ceil(math.Sqrt(float64(n))))

	ans := 0
	for k := 1; k < B; k++ {
		pre := 0
		for _, b := range s {
			cnt70[pre+n]++
			pre += k*int(b&1) - 1
			ans += int(cnt70[pre+n])
		}
		pre = 0
		for _, b := range s {
			cnt70[pre+n]--
			pre += k*int(b&1) - 1
		}
	}

	pos1 := []int{-1}
	for i, b := range s {
		if b == '1' {
			pos1 = append(pos1, i)
		}
	}

	tot1 := 0
	for i, b := range s {
		tot1 += int(b & 1)
		mn := min(tot1, n/B)
		for cnt1 := 1; cnt1 <= mn; cnt1++ {
			j := tot1 - cnt1
			maxK := (i - pos1[j]) / cnt1
			minK := max((i-pos1[j+1])/cnt1+1, B)
			ans += max(maxK-minK+1, 0)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1270F(bufio.NewReader(os.Stdin), os.Stdout) }
