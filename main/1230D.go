package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1230D(reader io.Reader, writer io.Writer) {
	type pair struct {
		a, b int64
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprint(out, 0)
		return
	}
	arr := make([]pair, n)
	for i := range arr {
		Fscan(in, &arr[i].a)
	}
	for i := range arr {
		Fscan(in, &arr[i].b)
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i].a < arr[j].a })
	multi := map[int64]int{}
	for i := 0; i < n-1; i++ {
		if arr[i].a == arr[i+1].a {
			multi[arr[i].a] = 1
		}
	}
	multiList := []int64{}
	for k := range multi {
		multiList = append(multiList, k)
	}
	ans := int64(0)
	for _, p := range arr {
		a, b := p.a, p.b
		if _, ok := multi[a]; ok {
			ans += b
		} else {
			for _, k := range multiList {
				if a&^k == 0 {
					ans += b
					break
				}
			}
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1230D(os.Stdin, os.Stdout)
//}
