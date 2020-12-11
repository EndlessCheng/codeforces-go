package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1355A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
1 4
487 1
487 2
487 3
487 4
487 5
487 6
487 7
outputCopy
42
487
519
528
544
564
588
628`
	testutil.AssertEqualCase(t, rawText, 0, CF1355A)
}
