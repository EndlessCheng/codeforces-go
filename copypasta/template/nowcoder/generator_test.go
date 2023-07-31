package nowcoder

import (
	"fmt"
	"os"
	"testing"
)

const contestID = 62033
const openPage = true

func TestGenNowCoderTemplates(t *testing.T) {
	emailOrPhone := os.Getenv("NOWCODER_USERNAME")
	cipherPwd := os.Getenv("NOWCODER_CIPHER_PWD")
	contestDir := fmt.Sprintf("../../../%d/", contestID)
	if err := GenNowCoderTemplates(emailOrPhone, cipherPwd, contestDir, contestID, openPage); err != nil {
		t.Fatal(err)
	}
}
