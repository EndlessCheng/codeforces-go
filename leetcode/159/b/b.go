package main

import (
	"sort"
	"strings"
)

func removeSubfolders(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, path := range folder[1:] {
		if tailAns := ans[len(ans)-1]; !strings.HasPrefix(path, tailAns) || strings.Count(path, "/") == strings.Count(tailAns, "/") {
			ans = append(ans, path)
		}
	}
	return
}
