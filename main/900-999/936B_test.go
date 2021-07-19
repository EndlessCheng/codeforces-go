package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/936/B
// https://codeforces.com/problemset/status/936/problem/B
func TestCF936B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
2 2 3
2 4 5
1 4
1 5
0
1
outputCopy
Win
1 2 4 5 
inputCopy
3 2
1 3
1 1
0
2
outputCopy
Lose
inputCopy
2 2
1 2
1 1
1
outputCopy
Draw`
	testutil.AssertEqualCase(t, rawText, 0, CF936B)
}
