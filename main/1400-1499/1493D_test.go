package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1493/D
// https://codeforces.com/problemset/status/1493/problem/D
func TestCF1493D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 3
1 6 8 12
1 12
2 3
3 3
outputCopy
2
2
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1493D)
}
