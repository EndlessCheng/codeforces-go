// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`3
3
101
4
1011
5
10001`,
			`Yes
No
No`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
