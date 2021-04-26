package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1097/D
// https://codeforces.com/problemset/status/1097/problem/D
func TestCF1097D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 1
outputCopy
3
inputCopy
6 2
outputCopy
875000008
inputCopy
60 5
outputCopy
237178099`
	testutil.AssertEqualCase(t, rawText, 0, CF1097D)
}
