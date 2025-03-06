package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf17E(in io.Reader, out io.Writer) {
	const mod = 51123987
	var n, tot int
	var s []byte
	Fscan(in, &n, &s)
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
		tot += hl / 2
	}
	tot %= mod
	ans := tot * (tot - 1) / 2

	diff := make([]int, n+1)
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if l <= r {
			diff[(l+r+1)/2]++
			diff[r+1]--
		}
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
		sum[i+1] = (sum[i] + v) % mod
	}
	for i, hl := range halfLen {
		l, r := (i-hl)/2, (i+hl)/2-2
		if l <= r {
			ans -= sum[(l+r)/2] - sum[max(l-1, 0)]
		}
	}
	Fprint(out, (ans%mod+mod)%mod)
}

//func main() { cf17E(bufio.NewReader(os.Stdin), os.Stdout) }
