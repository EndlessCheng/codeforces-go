package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/653/D
// https://codeforces.com/problemset/status/653/problem/D
func TestCF653D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 3
1 2 2
2 4 1
1 3 1
3 4 2
outputCopy
1.5000000000
inputCopy
5 11 23
1 2 3
2 3 4
3 4 5
4 5 6
1 3 4
2 4 5
3 5 6
1 4 2
2 5 3
1 5 2
3 2 30
outputCopy
10.2222222222
inputCopy
3 2 100000
1 2 1
2 3 1
outputCopy
1.0000000000
inputCopy
3 2 100000
1 2 1
2 3 1000000
outputCopy
1.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF653D)
}
