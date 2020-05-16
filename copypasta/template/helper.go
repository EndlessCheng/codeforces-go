package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func copyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// 目录需提前创建
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}

	return nil
}

// https://codeforces.ml/problemset/problem/1293/C
// https://codeforces.ml/contest/1353/problem/A
// https://codeforces.ml/gym/102253/problem/C
func parseCodeforcesProblemURL(urlStr string) (contestID, problemID string, isGYM bool) {
	sp := strings.Split(urlStr, "/")
	switch {
	case strings.Contains(urlStr, "/problemset/problem/"):
		return sp[len(sp)-2], sp[len(sp)-1], false
	case strings.Contains(urlStr, "/contest/"):
		return sp[len(sp)-3], sp[len(sp)-1], false
	case strings.Contains(urlStr, "/gym/"):
		return sp[len(sp)-3], sp[len(sp)-1], true
	default:
		panic("invalid URL")
	}
}

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
