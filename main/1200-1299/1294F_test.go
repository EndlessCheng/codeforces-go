package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1294F(t *testing.T) {
	// just copy from website
	rawText := `
input
8
1 2
2 3
3 4
4 5
4 6
3 7
3 8
output
5
1 8 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1294F)
}
