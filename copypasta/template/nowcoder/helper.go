package nowcoder

import (
	"path/filepath"
	"strings"
)

// "3,[1,2]" => ["3", "1,2"]
func splitRawInput(s string) (sp []string) {
	s = strings.TrimSpace(s)
	s += "," // for convenience
	dep := 0
	for i := 0; i < len(s); i++ {
		st := i
	o:
		for ; i < len(s); i++ {
			switch s[i] {
			case '[':
				dep++
			case ']':
				dep--
			case ',':
				if dep == 0 {
					sp = append(sp, s[st:i])
					break o
				}
			}
		}
	}
	return
}

func absPath(path string) string {
	p, _ := filepath.Abs(path)
	return p
}
