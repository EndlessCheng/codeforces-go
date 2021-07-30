package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/981/E
// https://codeforces.com/problemset/status/981/problem/E
func TestCF981E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 3 1
2 4 2
3 4 4
outputCopy
4
1 2 3 4 
inputCopy
7 2
1 5 1
3 7 2
outputCopy
3
1 2 3 
inputCopy
10 3
1 1 2
1 1 3
1 1 6
outputCopy
6
2 3 5 6 8 9 `
	testutil.AssertEqualCase(t, rawText, 0, CF981E)
}
