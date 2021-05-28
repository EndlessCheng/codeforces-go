package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/547/C
// https://codeforces.com/problemset/status/547/problem/C
func TestCF547C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
1 2 3 4 6
1
2
3
4
5
1
outputCopy
0
1
3
5
6
2`
	testutil.AssertEqualCase(t, rawText, 0, CF547C)
}
