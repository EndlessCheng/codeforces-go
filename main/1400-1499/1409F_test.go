package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1409F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
bbaa
ab
outputCopy
3
inputCopy
7 3
asddsaf
sd
outputCopy
10
inputCopy
15 6
qwertyhgfdsazxc
qa
outputCopy
16
inputCopy
7 2
abacaba
aa
outputCopy
15`
	testutil.AssertEqualCase(t, rawText, 0, CF1409F)
}
