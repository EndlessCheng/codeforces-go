package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1281/B
// https://codeforces.com/problemset/status/1281/problem/B
func TestCF1281B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
AZAMON APPLE
AZAMON AAAAAAAAAAALIBABA
APPLE BANANA
outputCopy
AMAZON
---
APPLE`
	testutil.AssertEqualCase(t, rawText, 0, CF1281B)
}
