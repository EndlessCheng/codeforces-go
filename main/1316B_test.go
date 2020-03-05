package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1316B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
4
abab
6
qwerty
5
aaaaa
6
alaska
9
lfpbavjsm
1
p
outputCopy
abab
1
ertyqw
3
aaaaa
1
aksala
6
avjsmbpfl
5
p
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1316B)
}
