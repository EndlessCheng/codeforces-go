package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF476B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
++-+-
+-+-+
outputCopy
1.000000000000
inputCopy
+-+-
+-??
outputCopy
0.500000000000
inputCopy
+++
??-
outputCopy
0.000000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF476B)
}
