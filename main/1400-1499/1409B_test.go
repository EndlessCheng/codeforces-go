package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
10 10 8 5 3
12 8 8 7 2
12343 43 4543 39 123212
1000000000 1000000000 1 1 1
1000000000 1000000000 1 1 1000000000
10 11 2 1 5
10 11 9 1 10
outputCopy
70
77
177177
999999999000000000
999999999
55
10`
	testutil.AssertEqualCase(t, rawText, 0, CF1409B)
}
