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

func parseCodeforcesProblemURL(urlStr string) (contestID, problemID string) {
	splits := strings.Split(urlStr, "/")
	return splits[len(splits)-2], splits[len(splits)-1]
}

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
