package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1080E(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sz := n*2 + 3
	halfLen := make([]int, sz-2)
	halfLen[1] = 1
	for r := range m {
		t := make([][26]uint8, sz)
		odd := make([]int8, sz)
		odd[0] = 2
		odd[sz-1] = 2
		for l := r; l >= 0; l-- {
			for i, row := range a {
				t[i*2+2][row[l]-'a']++
				odd[i*2+2] += int8(t[i*2+2][row[l]-'a']%2*2) - 1
			}
			boxM, boxR := 0, 0
			for i := 2; i < sz-2; i++ {
				hl := 0
				if i < boxR {
					hl = min(halfLen[boxM*2-i], boxR-i)
				}
				for odd[i-hl] <= 1 && odd[i+hl] <= 1 && t[i-hl] == t[i+hl] {
					hl++
					boxM, boxR = i, i+hl
				}
				halfLen[i] = hl
				ans += hl / 2
			}
		}
	}
	Fprint(out, ans)
}

//func main() { cf1080E(bufio.NewReader(os.Stdin), os.Stdout) }
