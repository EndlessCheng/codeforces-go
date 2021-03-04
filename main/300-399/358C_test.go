package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/358/C
// https://codeforces.com/problemset/status/358/problem/C
func TestCF358C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
0
1
0
1
2
0
1
2
3
0
outputCopy
0
pushStack
1 popStack
pushStack
pushQueue
2 popStack popQueue
pushStack
pushQueue
pushFront
3 popStack popQueue popFront
inputCopy
4
1
2
3
0
outputCopy
pushStack
pushQueue
pushFront
3 popStack popQueue popFront
inputCopy
2
0
1
outputCopy
0
pushStack`
	testutil.AssertEqualCase(t, rawText, 0, CF358C)
}
