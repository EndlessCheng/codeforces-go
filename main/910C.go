package main

import (
	"bufio"
	. "fmt"
	"io"
)

func perm910C(arr []int, i int, do func([]int)) {
	if i == len(arr) {
		do(arr)
		return
	}
	perm910C(arr, i+1, do)
	for j := i + 1; j < len(arr); j++ {
		arr[i], arr[j] = arr[j], arr[i]
		perm910C(arr, i+1, do)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func Sol910C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	pow10 := [6]int{1, 1e1, 1e2, 1e3, 1e4, 1e5}

	var n int
	var coefs [10]int
	var mustNotZero [10]bool
	for Fscan(in, &n); n > 0; n-- {
		var s string
		Fscan(in, &s)
		mustNotZero[s[0]-'a'] = true
		for i, c := range s {
			coefs[c-'a'] += pow10[len(s)-1-i]
		}
	}

	minVal := 1 << 30
	perm910C([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, func(digits []int) {
		val := 0
		for i, d := range digits {
			if d == 0 {
				if mustNotZero[i] {
					return
				}
				continue
			}
			val += coefs[i] * d
		}
		if val < minVal {
			minVal = val
		}
	})
	Fprintln(out, minVal)
}

//func main() {
//	Sol910C(os.Stdin, os.Stdout)
//}
