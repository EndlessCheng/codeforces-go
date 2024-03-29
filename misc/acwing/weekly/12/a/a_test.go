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
			`1
6
1 5 5 1 6 1`,
			`3
5 6 1`,
		},
		{
			`1
5
2 4 2 4 4`,
			`2
2 4`,
		},
		{
			`1
5
6 6 6 6 6`,
			`1
6`,
		},
		
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
