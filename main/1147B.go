package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1147B(reader io.Reader, writer io.Writer) {
	calcMinPeriod := func(pattern []int) int {
		n := len(pattern)
		maxMatchLengths := make([]int, n)
		maxLength := 0
		for i := 1; i < n; i++ {
			c := pattern[i]
			for maxLength > 0 && pattern[maxLength] != c {
				maxLength = maxMatchLengths[maxLength-1]
			}
			if pattern[maxLength] == c {
				maxLength++
			}
			maxMatchLengths[i] = maxLength
		}
		if val := maxMatchLengths[n-1]; val > 0 {
			if val%(n-val) == 0 {
				return val/(n-val) + 1
			}
		}
		return 0
	}
	calcGCD := func(a, b int64) int64 {
		for b > 0 {
			a, b = b, a%b
		}
		return a
	}
	calcLCM := func(a, b int64) int64 {
		return a / calcGCD(a, b) * b
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	lenPosMat := make([][]int, n/2+1)
	for ; m > 0; m-- {
		var a, b int
		Fscan(in, &a, &b)
		if a > b {
			a, b = b, a
		}
		segLen := b - a
		pos := a
		if segLen > n-segLen {
			segLen = n - segLen
			pos = b
		}
		lenPosMat[segLen] = append(lenPosMat[segLen], pos)
	}

	k := int64(-1)
	for segLen, posList := range lenPosMat {
		if len(posList) == 0 {
			continue
		}
		var p int64
		if segLen*2 < n {
			sort.Ints(posList)
			posList = append(posList, posList[0]+n)
			diffList := make([]int, len(posList)-1)
			for i := range diffList {
				diffList[i] = posList[i+1] - posList[i] - segLen
			}
			p = int64(calcMinPeriod(diffList))
			if p == 0 {
				Fprint(out, "No")
				return
			}
		} else if segLen*2 == n {
			p = int64(segLen)
		} else {
			break
		}
		if k == -1 {
			k = p
		} else {
			k = calcLCM(k, p)
		}
		if k >= int64(n) {
			Fprint(out, "No")
			return
		}
	}
	Fprint(out, "Yes")
}

func main() {
	Sol1147B(os.Stdin, os.Stdout)
}
