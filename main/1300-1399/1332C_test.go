package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1332C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
6 2
abaaba
6 3
abaaba
36 9
hippopotomonstrosesquippedaliophobia
21 7
wudixiaoxingxingheclp
outputCopy
2
0
23
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1332C)
}
