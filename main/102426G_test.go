package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF102426G(t *testing.T) {
	customInputs := []string{
		`3
allocate 1
free 1024
allocate 1`, `5
free 1
free 1
free 1
free 2
free 2`,
	}
	customAnswers := []string{
		`ERROR!
0 0 0 0 0 0 0 0 0 0 1
1 1 1 1 1 1 1 1 1 1 0`, `1 0 0 0 0 0 0 0 0 0 0
2 0 0 0 0 0 0 0 0 0 0
3 0 0 0 0 0 0 0 0 0 0
3 1 0 0 0 0 0 0 0 0 0
3 2 0 0 0 0 0 0 0 0 0`,
	}
	testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, CF102426G)
}
