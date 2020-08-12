package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol920E(t *testing.T) {
	// just copy from website
	rawText := `
2 1
1 2
outputCopy
2
1 1 `
	testutil.AssertEqualCase(t, rawText, 0, Sol920E)
}
