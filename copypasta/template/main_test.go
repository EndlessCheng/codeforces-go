package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"path/filepath"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customInputs := []string{
		``,
	}
	customAnswers := []string{
		``,
	}
	if len(customInputs) > 0 && strings.TrimSpace(customInputs[0]) != "" {
		testutil.AssertEqualStringCase(t, customInputs, customAnswers, 0, run)
		t.Log("======================================")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCase(t, dir, 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}
