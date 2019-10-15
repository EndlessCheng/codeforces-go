package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol486C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, p int
	var s string
	Fscan(in, &n, &p, &s)
	n2 := n / 2
	p--
	if p >= n2 {
		p = n - 1 - p
	}

	pos := []int{}
	ans := 0
	for i := 0; i < n2; i++ {
		if s[i] != s[n-1-i] {
			pos = append(pos, i)
			d := (s[i] - s[n-1-i] + 26) % 26
			if d > 13 {
				d = 26 - d
			}
			ans += int(d)
		}
	}
	if ans == 0 {
		Fprint(out, 0)
		return
	}
	dl, dr := p-pos[0], pos[len(pos)-1]-p
	if dl <= 0 {
		ans += dr
	} else if dr <= 0 {
		ans += dl
	} else {
		if dl <= dr {
			ans += 2*dl + dr
		} else {
			ans += 2*dr + dl
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol486C(os.Stdin, os.Stdout)
//}
