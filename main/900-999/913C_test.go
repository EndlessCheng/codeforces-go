package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF913C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 12
20 30 70 90
outputCopy
150
inputCopy
4 3
10000 1000 100 10
outputCopy
10
inputCopy
4 3
10 100 1000 10000
outputCopy
30
inputCopy
5 787787787
123456789 234567890 345678901 456789012 987654321
outputCopy
44981600785557577`
	testutil.AssertEqualCase(t, rawText, 0, CF913C)
}
