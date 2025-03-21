package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf30E_zfunc(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)

	rev := []byte(s)
	slices.Reverse(rev)
	t := string(rev) + "#" + s
	z := make([]int, n*2+1)
	type pair struct{ v, i int }
	maxZ := make([]pair, n+1)
	q := []int{}
	for i, boxL, boxR := 1, 0, 0; i <= n*2; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] <= n*2 && t[z[i]] == t[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
		if i > n {
			for len(q) > 0 && q[0]+z[q[0]]-1 < i {
				q = q[1:]
			}
			if z[i] > 0 {
				q = append(q, i)
			}
			maxZ[i-n] = maxZ[i-n-1]
			if len(q) > 0 && i-q[0]+1 > maxZ[i-n].v {
				maxZ[i-n] = pair{i - q[0] + 1, q[0] - n - 1}
			}
		}
	}

	halfLen := make([]int, n)
	var boxM, boxR, mx, mxI int
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
		k := min(maxZ[l].v, n-1-r)
		s := r - l + 1 + k*2
		if s > mx {
			mx, mxI = s, i
		}
	}

	l, r := mxI-halfLen[mxI]+1, mxI+halfLen[mxI]-1
	k := min(maxZ[l].v, n-1-r)
	if k == 0 {
		Fprintln(out, 1)
		Fprintln(out, l+1, r-l+1)
	} else {
		Fprintln(out, 3)
		Fprintln(out, maxZ[l].i+1, k)
		Fprintln(out, l+1, r-l+1)
		Fprintln(out, n-k+1, k)
	}
}

//func main() { cf30E(bufio.NewReader(os.Stdin), os.Stdout) }
