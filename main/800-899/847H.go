package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol847H(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	arr := make([]int64, n)
	for i := range arr {
		Fscan(in, &arr[i])
	}

	l, r := 0, n-1
	for ; l < n-1 && arr[l] < arr[l+1]; l++ {
	}
	for ; r > 0 && arr[r-1] > arr[r]; r-- {
	}
	incSum := int64(0)
	for l < r {
		if l+1 == r {
			if arr[l] == arr[r] {
				incSum++
			}
			break
		}
		if arr[l] <= arr[r] {
			if delta := arr[l] + 1 - arr[l+1]; delta >= 0 {
				incSum += delta
				arr[l+1] += delta
			}
			l++
		} else {
			if delta := arr[r] + 1 - arr[r-1]; delta >= 0 {
				incSum += delta
				arr[r-1] += delta
			}
			r--
		}
	}
	Fprintln(out, incSum)
}

//func main() {
//	Sol847H(os.Stdin, os.Stdout)
//}
