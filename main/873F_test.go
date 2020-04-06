package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF873F(t *testing.T) {
	// just copy from website
	rawText := `
input
5
ababa
00100
output
5
input
5
ababa
00000
output
6
input
5
ababa
11111
output
0
input
5
aaaaa
00000
output
5`
	testutil.AssertEqualCase(t, rawText, 0, CF873F)
}
