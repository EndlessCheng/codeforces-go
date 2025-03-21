package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf30E(in io.Reader, out io.Writer) {
	calcPi := func(s string) []int {
		pi := make([]int, len(s))
		match := 0
		for i := 1; i < len(pi); i++ {
			v := s[i]
			for match > 0 && s[match] != v {
				match = pi[match-1]
			}
			if s[match] == v {
				match++
			}
			pi[i] = match
		}
		return pi
	}

	var s string
	Fscan(in, &s)
	n := len(s)

	rev := []byte(s)
	slices.Reverse(rev)
	pi := calcPi(string(rev) + "#" + s)[n:] // 注意 pi[0] 对应 '#'
	preMaxI := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preMaxI[i] = preMaxI[i-1]
		if pi[i] > pi[preMaxI[i]] {
			preMaxI[i] = i
		}
	}

	halfLen := make([]int, n)
	var boxM, boxR, maxSum, maxI int
	for i := range halfLen {
		hl := 1
		if i < boxR {
			hl = min(halfLen[boxM*2-i], boxR-i)
		}
		for i >= hl && i+hl < n && s[i-hl] == s[i+hl] {
			hl++
			boxM, boxR = i, i+hl
		}
		halfLen[i] = hl

		l, r := i-hl+1, i+hl-1
		k := min(pi[preMaxI[l]], n-1-r)
		sum := r - l + 1 + k*2
		if sum > maxSum {
			maxSum, maxI = sum, i
		}
	}

	l, r := maxI-halfLen[maxI]+1, maxI+halfLen[maxI]-1
	k := min(pi[preMaxI[l]], n-1-r)
	if k == 0 {
		Fprintln(out, 1)
		Fprintln(out, l+1, r-l+1)
	} else {
		Fprintln(out, 3)
		Fprintln(out, preMaxI[l]-k+1, k)
		Fprintln(out, l+1, r-l+1)
		Fprintln(out, n-k+1, k)
	}
}

//func main() { cf30E(bufio.NewReader(os.Stdin), os.Stdout) }
