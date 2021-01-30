package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1476/E
// https://codeforces.com/problemset/status/1476/problem/E
func TestCF1476E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3 4
_b_d
__b_
aaaa
ab__
_bcd
abcd 4
abba 2
dbcd 5
outputCopy
YES
3 2 4 5 1 
inputCopy
1 1 3
__c
cba 1
outputCopy
NO
inputCopy
2 2 2
a_
_b
ab 1
ab 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1476E)
}
