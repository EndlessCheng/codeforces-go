package atcoder

import "testing"

// TODO: update REVEL_SESSION
func TestGenAtCoderContestTemplates(t *testing.T) {
	const contestID = "abc174"
	//username := os.Getenv("ATCODER_USERNAME")
	//password := os.Getenv("ATCODER_PASSWORD")
	if err := GenAtCoderContestTemplates(contestID); err != nil {
		t.Fatal(err)
	}
}

// https://atcoder.jp/contests/abc161/tasks/abc161_f
// https://atcoder.jp/contests/abc161/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc161_f&orderBy=source_length
func TestGenAtCoderProblemTemplate(t *testing.T) {
	const problemURL = ""
	if err := GenAtCoderProblemTemplate(problemURL); err != nil {
		t.Fatal(err)
	}
}
