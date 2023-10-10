package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
	"strings"
)

// https://space.bilibili.com/206214
func CF1680C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		ans := len(s)
		in0 := 0
		out1 := strings.Count(s, "1")
		left := 0
		for _, b := range s {
			v := int(b & 1)
			in0 += v ^ 1
			out1 -= v
			for in0 > out1 {
				v = int(s[left] & 1)
				in0 -= v ^ 1
				out1 += v
				left++
			}
			ans = min(ans, out1)
		}
		Fprintln(out, ans)
	}
}

func CF1680C_binarySearch(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		tot1 := strings.Count(s, "1")
		Fprintln(out, sort.Search(n, func(mx int) bool {
			in0 := 0     // 窗口内的 0 的个数
			out1 := tot1 // 窗口外的 1 的个数
			left := 0
			for _, b := range s {
				v := int(b & 1)
				in0 += v ^ 1
				out1 -= v
				for in0 > mx {
					v = int(s[left] & 1)
					in0 -= v ^ 1
					out1 += v
					left++
				}
				if out1 <= mx {
					return true
				}
			}
			return false
		}))
	}
}

//func main() { CF1680C(os.Stdin, os.Stdout) }
