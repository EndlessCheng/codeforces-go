package nowcoder

import (
	"fmt"
	"os"
	"testing"
)

func TestGenNowCoderTemplates(t *testing.T) {
	email := os.Getenv("NOWCODER_USERNAME")
	cipherPwd := os.Getenv("NOWCODER_CIPHER_PWD")
	const contestID = 6778
	contestDir := fmt.Sprintf("../../../misc/nowcoder/%d/", contestID)
	if err := GenNowCoderTemplates(email, cipherPwd, contestDir, contestID, `// github.com/EndlessCheng/codeforces-go`); err != nil {
		t.Fatal(err)
	}
}
