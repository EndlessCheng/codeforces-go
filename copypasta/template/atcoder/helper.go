package atcoder

import "path/filepath"

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
