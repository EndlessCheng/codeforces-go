package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	contestID = 1260
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

func TestGenCodeforcesContestTemplates(t *testing.T) {
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

	tips := fmt.Sprintf("cd %[1]d\ncf submit %[1]d a a/main.go\n", contestID)
	if err := ioutil.WriteFile(rootPath+"tips.txt", []byte(tips), 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGenCodeforcesNormalTemplates(t *testing.T) {
	const rawID = "613/A"
	problemID := strings.Replace(rawID, "/", "", -1)
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func Sol%[1]s(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() {
	Sol%[1]s(os.Stdin, os.Stdout)
}
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol%[1]s(t *testing.T) {
	// just copy from website
	rawText := `+"`\n`"+`
	testutil.AssertEqualCase(t, rawText, 0, Sol%[1]s)
}
`, problemID)
	rootPath := "../../main/"
	if err := ioutil.WriteFile(rootPath+problemID+".go", []byte(mainStr), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(rootPath+problemID+"_test.go", []byte(mainTestStr), 0644); err != nil {
		t.Fatal(err)
	}
}
