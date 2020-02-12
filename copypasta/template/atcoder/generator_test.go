package atcoder

import (
	"os"
	"testing"
)

func TestGenAtCoderTests(t *testing.T) {
	username := os.Getenv("ATCODER_USERNAME")
	password := os.Getenv("ATCODER_PASSWORD")
	if err := GenAtCoderTests(username, password); err != nil {
		t.Fatal(err)
	}
}
