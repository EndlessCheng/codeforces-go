package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"math/rand"
	"testing"
)

func TestCompare(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() (a []int) {
		n := 5
		a = make([]int, n)
		for i := range a {
			a[i] = rand.Intn(10) + 1
		}
		return
	}

	runAC := func(a []int) (ans int) {

		return
	}

	run := func(a []int) (ans int) {

		return
	}

	testutil.CompareInf(t, inputGenerator, runAC, run)
}
