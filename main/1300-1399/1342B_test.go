package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1342B(t *testing.T) {
	// just copy from website
	rawText := `
input
4
00
01
111
110
output
00
01
11111`
	testutil.AssertEqualCase(t, rawText, 0, CF1342B)
}
