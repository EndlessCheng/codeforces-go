package main

import (
	"io"
	"os"
	"strconv"
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

func parseProblemIDFromURL(urlStr string) string {
	// https://codeforces.ml/contest/908/problem/C
	splits := strings.Split(urlStr, "/")
	for _, s := range splits {
		if _, err := strconv.Atoi(s); err == nil {
			return s + splits[len(splits)-1]
		}
	}
	panic("invalid url")
}
