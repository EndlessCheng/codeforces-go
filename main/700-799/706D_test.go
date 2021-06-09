package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/706/D
// https://codeforces.com/problemset/status/706/problem/D
func TestCF706D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
+ 8
+ 9
+ 11
+ 6
+ 1
? 3
- 8
? 3
? 8
? 11
outputCopy
11
10
14
13`
	testutil.AssertEqualCase(t, rawText, 0, CF706D)
}
