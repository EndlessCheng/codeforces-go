package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1370C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1
2
3
4
5
6
12
outputCopy
FastestFinger
Ashishgup
Ashishgup
FastestFinger
Ashishgup
FastestFinger
Ashishgup`
	testutil.AssertEqualCase(t, rawText, 0, CF1370C)
}
