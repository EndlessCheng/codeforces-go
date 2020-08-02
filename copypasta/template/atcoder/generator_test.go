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

// https://atcoder.jp/contests/abc162/tasks/abc162_e
// https://atcoder.jp/contests/abc162/submissions
// https://atcoder.jp/contests/abc162/submissions?f.Language=4026&f.Status=AC&f.Task=abc162_e&orderBy=source_length
// https://atcoder.jp/contests/abc161/submissions?f.Language=3013&f.Status=AC&f.Task=abc161_f&orderBy=source_length
func TestGenAtCoderProblemTemplate(t *testing.T) {
	const problemURL = ""
	if err := GenAtCoderProblemTemplate(problemURL); err != nil {
		t.Fatal(err)
	}
}
