package main

// github.com/EndlessCheng/codeforces-go
type Robot struct {
	w, h, step int
}

func Constructor(width, height int) Robot {
	return Robot{width, height, 0}
}

func (r *Robot) Step(num int) {
	// 由于机器人只能走外圈，那么走 (w+h-2)*2 步后会回到起点
	// 把 step 取模调整到 [1, (w+h-2)*2]，这样不需要特判处于 step == 0 时的方向
	r.step = (r.step+num-1)%((r.w+r.h-2)*2) + 1
}

func (r *Robot) getState() (x, y int, dir string) {
	w, h, step := r.w, r.h, r.step
	switch {
	case step < w:
		return step, 0, "East"
	case step < w+h-1:
		return w - 1, step - w + 1, "North"
	case step < w*2+h-2:
		return w*2 + h - step - 3, h - 1, "West"
	default:
		return 0, (w+h)*2 - step - 4, "South"
	}
}

func (r *Robot) GetPos() []int {
	x, y, _ := r.getState()
	return []int{x, y}
}

func (r *Robot) GetDir() string {
	_, _, d := r.getState()
	return d
}
