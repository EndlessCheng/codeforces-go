package copypasta

import "testing"

func TestACAM1(t *testing.T) {
	patterns := []string{
		"abc",
		"abac",
		"babc",
		"bc",
		"cab",
	}
	ac := newACAM(patterns)
	ac.debug("abcabac")
}

func TestACAM2(t *testing.T) {
	patterns := []string{
		"a",
		"aa",
		"aa",
	}
	ac := newACAM(patterns)
	ac.debug("aaa")
}

// 有中间 skip 的例子，可以用来说明 last 的作用
// 可以发现不用 last 的话，会有大量 skip
func TestACAMSkip(t *testing.T) {
	patterns := []string{
		"a",
		"abac",
		"babc",
		"abbc",
		"abca",
		"caab",
	}
	ac := newACAM(patterns)
	ac.debug("abcbabcbabcabaacbababacbabc")
}
