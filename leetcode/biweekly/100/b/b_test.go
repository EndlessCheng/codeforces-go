// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"sort"
	"testing"
)

func Test_b(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithFile(t, maximizeGreatness, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode.cn/contest/biweekly-contest-100/problems/maximize-greatness-of-an-array/
// https://leetcode.cn/problems/maximize-greatness-of-an-array/

func TestCompareInf(_t *testing.T) {
	testutil.DebugTLE = 0

	inputGenerator := func() (a []int) {
		//return []int{2,1,3}
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 6)
		a = rg.IntSlice(n, 1, 4)
		return
	}

	runAC := func(a []int) (ans int) {
		// 若要修改 a，必须先 copy 一份，在 copied 上修改
		a = append(a[:0:0], a...)
		sort.Ints(a)
		for i := 0; i < len(a); i++ {
			b := a[i:]
			c := 0
			for i, v := range b {
				if v > a[i] {
					c++
				}
			}
			if c > ans {
				ans = c
			}
		}
		return
	}

	//// test examples first
	//if err := testutil.RunLeetCodeFuncWithFile(_t, runAC, "d.txt", 0); err != nil {
	//	_t.Fatal(err)
	//}
	//return

	testutil.CompareInf(_t, inputGenerator, runAC, maximizeGreatness)
}