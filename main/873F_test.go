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
200
aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
11011010000011010010011110001000001101110111001110000101001100000001010100001101111100010000101111100110010001111011010010000100111111000101110101110111110110000110100011011101001010000000111000100010
output
5200`
	testutil.AssertEqualCase(t, rawText, 0, CF873F)
}
