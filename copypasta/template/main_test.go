package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	customInputs := []string{
		``,
	}
	customAnswers := []string{
		``,
	}
	if len(customInputs) > 0 && strings.TrimSpace(customInputs[0]) != "" {
		testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, solve)
		t.Log("======================================")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCase(t, dir, 0, solve)
	_, problemName := filepath.Split(dir)
	t.Logf("Current problem is [%s]", problemName)
}
