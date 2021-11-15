package main

// github.com/EndlessCheng/codeforces-go
type Robot struct{}
var w, h, step int

func Constructor(width, height int) (_ Robot) {
	w, h, step = width, height, 0
	return
}

func (Robot) Move(num int) {
	// 由于机器人只能走外圈，那么走 (w+h-2)*2 步后会回到起点
	// 同时，将 step 取模固定在 [1,(w+h-2)*2] 范围内，这样不需要特判处于原点时的方向
	step = (step+num-1)%((w+h-2)*2) + 1
}

func get() (x, y int, dir string) {
	switch {
	case step < w:
		return step, 0, "East"
	case step < w+h-1:
		return w - 1, step - w + 1, "North"
	case step < w*2+h-2:
		return w*2 + h - 3 - step, h - 1, "West"
	default:
		return 0, (w+h-2)*2 - step, "South"
	}
}

func (Robot) GetPos() []int  { x, y, _ := get(); return []int{x, y} }
func (Robot) GetDir() string { _, _, d := get(); return d }
