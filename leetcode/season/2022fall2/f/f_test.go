// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_f(t *testing.T) {
	targetCaseNum := -1
	if err := testutil.RunLeetCodeFuncWithFile(t, reservoir, "f.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/season/2022-fall/problems/kskhHQ/
