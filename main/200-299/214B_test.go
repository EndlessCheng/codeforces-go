package main

import (
	"bufio"
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"strconv"
	"testing"
)

// https://codeforces.com/problemset/problem/214/B
// https://codeforces.com/problemset/status/214/problem/B
func TestCF214B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
0
outputCopy
0
inputCopy
11
3 4 5 4 5 3 5 3 4 4 0
outputCopy
5554443330
inputCopy
8
3 2 5 1 5 2 2 3
outputCopy
-1
inputCopy
2
0 6
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, 0, CF214B)
}

func TestCompareCF214B(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		rg.NewLine()
		rg.IntSlice(n, 0, 9)
		return rg.String()
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF_CF214B, CF214B)
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func diminish(count []int, k int) bool {
	for d := k; d < 10; d += 3 {
		if count[d] > 0 {
			count[d] -= 1
			return true
		}
	}
	return false
}

func runBF_CF214B(_r io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(_r)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(out)
	defer writer.Flush()
	n := scanInt(scanner)
	count := make([]int, 10)
	sum := 0
	for i := 0; i < n; i++ {
		x := scanInt(scanner)
		count[x] += 1
		sum += x
	}
	if count[0] == 0 {
		writer.WriteString("-1\n")
		return
	}
	if sum%3 == 1 {
		if !diminish(count, 1) {
			if !diminish(count, 2) || !diminish(count, 2) {
				writer.WriteString("0\n")
				return
			}
		}
	} else if sum%3 == 2 {
		if !diminish(count, 2) {
			if !diminish(count, 1) || !diminish(count, 1) {
				writer.WriteString("0\n")
				return
			}
		}
	}
	max := 0
	for d := 0; d < 10; d++ {
		if count[d] > 0 {
			max = d
		}
	}
	if max == 0 {
		writer.WriteString("0\n")
		return
	}
	for d := 9; d >= 0; d-- {
		for i := 0; i < count[d]; i++ {
			writer.WriteString(fmt.Sprintf("%d", d))
		}
	}
	writer.WriteString("\n")
}
