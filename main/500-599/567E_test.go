package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/567/E
// https://codeforces.com/problemset/status/567/problem/E
func TestCF567E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 7 1 6
1 2 2
1 3 10
2 3 7
2 4 8
3 5 3
4 5 2
5 6 1
outputCopy
YES
CAN 2
CAN 1
CAN 1
CAN 1
CAN 1
YES
inputCopy
3 3 1 3
1 2 10
2 3 10
1 3 100
outputCopy
YES
YES
CAN 81
inputCopy
2 2 1 2
1 2 1
1 2 2
outputCopy
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF567E)
}
