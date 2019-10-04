package main

import (
	"bufio"
	. "container/heap"
	. "fmt"
	"io"
	"sort"
)

type intHeap1140C struct {
	sort.IntSlice
}

func (h *intHeap1140C) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *intHeap1140C) Pop() (v interface{}) {
	n := len(h.IntSlice)
	h.IntSlice, v = h.IntSlice[:n-1], h.IntSlice[n-1]
	return
}

// github.com/EndlessCheng/codeforces-go
func Sol1140C(reader io.Reader, writer io.Writer) {
	type pair struct {
		x, y int
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	arr := make([]pair, n)
	for i := range arr {
		Fscan(in, &arr[i].y, &arr[i].x)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].x > arr[j].x || arr[i].x == arr[j].x && arr[i].y > arr[j].y })

	var ans, sum int64
	h := &intHeap1140C{}
	for _, p := range arr {
		Push(h, p.y)
		sum += int64(p.y)
		for h.Len() > k {
			sum -= int64(Pop(h).(int))
		}
		if newAns := sum * int64(p.x); newAns > ans {
			ans = newAns
		}
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1140C(os.Stdin, os.Stdout)
//}
