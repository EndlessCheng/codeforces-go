package main

// https://space.bilibili.com/206214/dynamic
type TextEditor struct{ l, r []byte }

func Constructor() TextEditor { return TextEditor{} }

func (t *TextEditor) AddText(text string) {
	t.l = append(t.l, text...)
}

func (t *TextEditor) DeleteText(k int) int {
	k0 := k
	for ; k > 0 && len(t.l) > 0; k-- {
		t.l = t.l[:len(t.l)-1]
	}
	return k0 - k
}

func (t *TextEditor) text() string {
	return string(t.l[max(len(t.l)-10, 0):])
}

func (t *TextEditor) CursorLeft(k int) string {
	for ; k > 0 && len(t.l) > 0; k-- {
		t.r = append(t.r, t.l[len(t.l)-1])
		t.l = t.l[:len(t.l)-1]
	}
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	for ; k > 0 && len(t.r) > 0; k-- {
		t.l = append(t.l, t.r[len(t.r)-1])
		t.r = t.r[:len(t.r)-1]
	}
	return t.text()
}

func max(a, b int) int { if b > a { return b }; return a }

