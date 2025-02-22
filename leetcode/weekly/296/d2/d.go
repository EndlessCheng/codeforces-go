package main

// https://space.bilibili.com/206214/dynamic
type TextEditor struct {
	left, right []byte // 光标左侧、右侧字符
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (t *TextEditor) AddText(text string) {
	t.left = append(t.left, text...) // 入栈
}

func (t *TextEditor) DeleteText(k int) int {
	k = min(k, len(t.left))
	t.left = t.left[:len(t.left)-k] // 出栈
	return k
}

func (t *TextEditor) text() string {
	// 光标左边至多 10 个字符
	return string(t.left[max(len(t.left)-10, 0):])
}

func (t *TextEditor) CursorLeft(k int) string {
	for k > 0 && len(t.left) > 0 {
		t.right = append(t.right, t.left[len(t.left)-1]) // 左手倒右手
		t.left = t.left[:len(t.left)-1]
		k--
	}
	return t.text()
}

func (t *TextEditor) CursorRight(k int) string {
	for k > 0 && len(t.right) > 0 {
		t.left = append(t.left, t.right[len(t.right)-1]) // 右手倒左手
		t.right = t.right[:len(t.right)-1]
		k--
	}
	return t.text()
}
