package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://space.bilibili.com/206214
func CF1861C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		curSize := 0
		sortedSize := 1
		unsortedSize := math.MaxInt
		for _, b := range s {
			if b == '+' {
				curSize++
			} else if b == '-' {
				curSize--
				if curSize < unsortedSize {
					unsortedSize = math.MaxInt // 后面 s[i]='1' 是可以的
				}
				if curSize < sortedSize {
					sortedSize = max(curSize, 1)
				}
			} else if b == '0' {
				if curSize <= sortedSize {
					Fprintln(out, "NO")
					continue o
				}
				unsortedSize = min(unsortedSize, curSize)
			} else {
				if curSize >= unsortedSize { // 长度 >= unsortedSize 的都是无序的
					Fprintln(out, "NO")
					continue o
				}
				sortedSize = max(curSize, 1)
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1861C(os.Stdin, os.Stdout) }
