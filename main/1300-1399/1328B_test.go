package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1328B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
5 1
5 2
5 8
5 10
3 1
3 2
20 100
outputCopy
aaabb
aabab
baaba
bbaaa
abb
bab
aaaaabaaaaabaaaaaaaa`
	testutil.AssertEqualCase(t, rawText, 0, CF1328B)
}
