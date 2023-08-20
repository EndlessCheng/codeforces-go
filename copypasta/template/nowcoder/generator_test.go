package nowcoder

import (
	"fmt"
	"os"
	"testing"
)

// 关闭电脑声音
// 确认语言为 Go
func TestGenNowCoderTemplates(t *testing.T) {
	contestID, err := fetchWeeklyContestId()
	if err != nil {
		t.Fatal(err)
	}
	emailOrPhone := os.Getenv("NOWCODER_USERNAME")
	cipherPwd := os.Getenv("NOWCODER_CIPHER_PWD")
	contestDir := fmt.Sprintf("../../../%d/", contestID)
	if err := GenNowCoderTemplates(emailOrPhone, cipherPwd, contestDir, contestID, true); err != nil {
		t.Fatal(err)
	}
}
