package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/626/problem/F
// https://codeforces.com/problemset/status/626/problem/F
func TestCF626F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
2 4 5
outputCopy
3
inputCopy
4 3
7 8 9 10
outputCopy
13
inputCopy
4 0
5 10 20 21
outputCopy
1
inputCopy
20 1000
50 50 100 100 150 150 200 200 250 250 300 300 350 350 400 400 450 450 500 500
outputCopy
97456952`
	testutil.AssertEqualCase(t, rawText, 0, CF626F)
}
