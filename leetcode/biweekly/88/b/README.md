下午 2 点在 B 站直播讲周赛和双周赛的题目，[欢迎关注](https://space.bilibili.com/206214/dynamic)~

---

由于最长上传前缀不会减小，因此用一个变量 $x$ 维护即可。

```py [sol1-Python3]
class LUPrefix:
    def __init__(self, n: int):
        self.x = 1
        self.s = set()

    def upload(self, video: int) -> None:
        self.s.add(video)

    # 时间复杂度：均摊 O(1)
    def longest(self) -> int:
        while self.x in self.s:
            self.x += 1
        return self.x - 1
```

```go [sol1-Go]
type LUPrefix struct {
	x   int
	has map[int]bool
}

func Constructor(int) LUPrefix {
	return LUPrefix{1, map[int]bool{}}
}

func (p LUPrefix) Upload(video int) {
	p.has[video] = true
}

// 时间复杂度：均摊 O(1)
func (p *LUPrefix) Longest() int {
	for p.has[p.x] {
		p.x++
	}
	return p.x - 1
}
```
