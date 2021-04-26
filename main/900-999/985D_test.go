package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/985/D
// https://codeforces.com/problemset/status/985/problem/D
func TestCF985D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
outputCopy
3
inputCopy
6 8
outputCopy
3
inputCopy
20 4
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, -1, CF985D)
}
