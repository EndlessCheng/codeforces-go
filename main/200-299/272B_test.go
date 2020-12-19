package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF272B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 4
outputCopy
3
inputCopy
3
5 3 1
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF272B)
}
