package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1690/problem/G
// https://codeforces.com/problemset/status/1690/problem/G
func TestCF1690G(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3

4 2
6 2 3 7
3 2
4 7

5 4
10 13 5 2 6
2 4
5 2
1 5
3 2

13 4
769 514 336 173 181 373 519 338 985 709 729 702 168
12 581
6 222
7 233
5 117
outputCopy
3 4 
4 4 2 3 
5 6 6 5 `
	testutil.AssertEqualCase(t, rawText, 0, CF1690G)
}
