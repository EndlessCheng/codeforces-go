package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1042B(t *testing.T) {
	// just copy from website
	rawText := `
4
5 C
6 B
16 BAC
4 A
outputCopy
15
inputCopy
2
10 AB
15 BA
outputCopy
-1
inputCopy
5
10 A
9 BC
11 CA
4 A
5 B
outputCopy
13
inputCopy
6
100 A
355 BCA
150 BC
160 AC
180 B
190 CA
outputCopy
250
inputCopy
2
5 BA
11 CB
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1042B)
}
