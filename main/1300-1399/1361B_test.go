package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1361/B
// https://codeforces.com/problemset/status/1361/problem/B
func TestCF1361B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 2
2 3 4 4 3
3 1
2 10 1000
4 5
0 1 1 100
1 8
89
outputCopy
4
1
146981438
747093407`
	testutil.AssertEqualCase(t, rawText, 0, CF1361B)
}
