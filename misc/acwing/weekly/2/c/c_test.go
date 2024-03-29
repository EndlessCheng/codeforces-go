// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`3 3 2
1 2 1
3 2 1
1 3 3`,
			`2
1 2`,
		},
		{
			`4 5 2
4 1 8
2 4 1
2 1 3
3 4 9
3 1 5`,
			`2
3 2`,
		},
		
	}
	target := 0
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
