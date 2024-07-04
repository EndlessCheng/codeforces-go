package main

import (
	"bufio"
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strconv"
	"testing"
)

// https://codeforces.com/problemset/problem/652/C
// https://codeforces.com/problemset/status/652/problem/C
func TestCF652C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 3 2 4
3 2
2 4
outputCopy
5
inputCopy
9 5
9 7 2 3 1 4 6 5 8
1 6
4 5
2 7
7 2
2 7
outputCopy
20`
	testutil.AssertEqualCase(t, rawText, 0, CF652C)
}

func TestCompareCF652C(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(2, 4)
		m := rg.Int(1,2)
		rg.NewLine()
		rg.Permutation( 1, n)
		for i := 0; i < m; i++ {
			rg.Int(rg.Int(1,n-1)+1,n)
			rg.NewLine()
		}
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF_CF652C, CF652C)
}

var scanner *bufio.Scanner
var writer *bufio.Writer

func getI64() int64 {
	scanner.Scan()
	x, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return x
}
func getI() int {
	return int(getI64())
}
func getF() float64 {
	scanner.Scan()
	x, _ := strconv.ParseFloat(scanner.Text(), 64)
	return x
}
func getS() string {
	scanner.Scan()
	return scanner.Text()
}

func runBF_CF652C(_r io.Reader, out io.Writer) {
	scanner = bufio.NewScanner(_r)
	scanner.Split(bufio.ScanWords)
	writer = bufio.NewWriter(out)
	defer writer.Flush()

	n, m := getI(), getI()
	values := make([]int, n)
	positions := make([]int, n + 1)
	for i := range values {
		values[i] = getI()
		positions[values[i]] = i
	}
	rightToLeft := make([]int, n)
	for i := 0; i < n; i++ {
		rightToLeft[i] = -1
	}
	for i := 0; i < m; i++ {
		a, b := positions[getI()], positions[getI()]
		if b < a {
			a, b = b, a
		}
		if a > rightToLeft[b] {
			rightToLeft[b] = a
		}
	}

	result := int64(0)
	left, right, seek := 0, 0, -1
	for right < n {
		for right + 1 < n {
			seek = rightToLeft[right + 1]
			if seek < left {
				right++
			} else {
				break
			}
		}
		for left <= seek {
			result += int64(right - left + 1)
			left++
		}
		right++
	}
	for left < n {
		result += int64(n - left)
		left++
	}

	writer.WriteString(fmt.Sprintf("%d\n", result))
}
