package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1335F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
01
RL
3 3
001
101
110
RLL
DLD
ULL
3 3
000
000
000
RRD
RLD
ULL
outputCopy
2 1
4 3
2 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1335F)
}
