package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf159D(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	n := len(s)
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range s {
		t = append(t, '#', c)
	}
	t = append(t, '#', '$')
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for t[i-hl] == t[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl
	}

	diff := make([]int, n+1)
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if r < l {
			continue
		}
		diff[(l+r+1)/2]++
		diff[r+1]--
	}
	cntEnd := diff[:n]
	for i := 1; i < n; i++ {
		cntEnd[i] += cntEnd[i-1]
	}
	for i := 1; i < n; i++ {
		cntEnd[i] += cntEnd[i-1]
	}

	sum := make([]int, n+1)
	for i, v := range cntEnd {
		sum[i+1] = sum[i] + v
	}
	ans := 0
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if r < l {
			continue
		}
		ans += sum[(l+r)/2] - sum[max(l-1, 0)]
	}
	Fprint(out, ans)
}

//func main() { cf159D(os.Stdin, os.Stdout) }
