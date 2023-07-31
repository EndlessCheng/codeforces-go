package nowcoder

import (
	"io"
	"os"
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
