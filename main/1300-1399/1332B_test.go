package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1332B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
6 10 15
2
4 9
23
437 519 865 808 909 391 194 291 237 395 323 365 511 497 781 737 871 559 731 697 779 841 961
outputCopy
1
1 1 1
2
2 1
11
4 7 8 10 7 3 10 7 7 8 3 1 1 5 5 9 2 2 3 3 4 11 6`
	testutil.AssertEqualCase(t, rawText, 0, CF1332B)
}
