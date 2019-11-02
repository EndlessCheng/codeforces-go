package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	contestID = 1245
	overwrite = false
)

func copyFile(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

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

func Test(t *testing.T) {
	rootPath := fmt.Sprintf("../../dash/%d/", contestID)
	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := filepath.Join(path, fileName)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	tips := fmt.Sprintf("cd %[1]d\ncf submit %[1]d c c/main.go\n", contestID)
	if err := ioutil.WriteFile(rootPath+"tips.txt", []byte(tips), 0644); err != nil {
		t.Fatal(err)
	}
}
