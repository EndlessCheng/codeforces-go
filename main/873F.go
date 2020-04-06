package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// github.com/EndlessCheng/codeforces-go
func CF873F(_r io.Reader, _w io.Writer) {
	var n int
	var s, forbidden []byte
	Fscan(bufio.NewReader(_r), &n, &s, &forbidden)

	for i, j := 0, n-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		forbidden[i], forbidden[j] = forbidden[j], forbidden[i]
		j--
	}
	ans := int64(0)
	prefixSum := make([]int, n+1)
	for i, v := range forbidden {
		if ans == 0 && v == '0' {
			ans = int64(n - i)
		}
		prefixSum[i+1] = prefixSum[i] + int(v-'0')
	}

	sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
	rank := make([]int, n)
	for i := range rank {
		rank[sa[i]] = i
	}
	height := make([]int, n)
	h := 0
	for i, sai := range rank {
		if h > 0 {
			h--
		}
		if sai > 0 {
			for j := int(sa[sai-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
			}
		}
		height[sai] = h
	}

	type pair struct{ v, i int }
	posL := make([]int, n)
	stack := []pair{{-1, -1}}
	for i, v := range height {
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posL[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}
	posR := make([]int, n)
	stack = []pair{{-1, n}}
	for i := n - 1; i >= 0; i-- {
		v := height[i]
		for {
			if top := stack[len(stack)-1]; top.v < v {
				posR[i] = top.i
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{v, i})
	}

	for i, h := range height {
		if h > 0 {
			if v := int64(h) * int64(posR[i]-posL[i]-(prefixSum[posR[i]]-prefixSum[posL[i]])); v > ans {
				ans = v
			}
		}
	}
	Fprint(_w, ans)
}

func main() { CF873F(os.Stdin, os.Stdout) }
