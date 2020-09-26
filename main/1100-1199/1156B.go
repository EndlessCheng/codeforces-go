package main

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

func Sol1156B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	const na = "No answer"

	var n int
	for Fscan(in, &n); n > 0; n-- {
		var s string
		Fscan(in, &s)

		cnt := make([]int, 26)
		for _, b := range s {
			cnt[b-'a']++
		}
		indexes := []int{}
		for i, c := range cnt {
			if c > 0 {
				indexes = append(indexes, i)
			}
		}

		ans := ""
		switch len(indexes) {
		case 1:
			ans = s
		case 2:
			if indexes[0]+1 == indexes[1] {
				ans = na
			} else {
				ans = s
			}
		case 3:
			i0, i1, i2 := indexes[0], indexes[1], indexes[2]
			if i0+1 == i1 && i1+1 == i2 {
				ans = na
			} else if i0+1 == i1 {
				ans = strings.Repeat(string(byte('a'+i0)), cnt[i0]) +
					strings.Repeat(string(byte('a'+i2)), cnt[i2]) +
					strings.Repeat(string(byte('a'+i1)), cnt[i1])
			} else if i1+1 == i2 {
				ans = strings.Repeat(string(byte('a'+i1)), cnt[i1]) +
					strings.Repeat(string(byte('a'+i0)), cnt[i0]) +
					strings.Repeat(string(byte('a'+i0)), cnt[i2])
			} else {
				ans = s
			}
		default:
			for i := 1; i < len(indexes); i += 2 {
				idx := indexes[i]
				ans += strings.Repeat(string(byte('a'+idx)), cnt[idx])
			}
			for i := 0; i < len(indexes); i += 2 {
				idx := indexes[i]
				ans += strings.Repeat(string(byte('a'+idx)), cnt[idx])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() {
//	Sol1156B(os.Stdin, os.Stdout)
//}
