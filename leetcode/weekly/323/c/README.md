[视频讲解](https://www.bilibili.com/video/BV1QK41167cr/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

由于数据范围小，暴力模拟即可。

```py [sol1-Python3]
class Allocator:
    def __init__(self, n: int):
        self.a = [0] * n

    def allocate(self, size: int, mID: int) -> int:
        cnt = 0
        for i, id in enumerate(self.a):
            if id:
                cnt = 0
            else:
                cnt += 1
                if cnt == size:
                    self.a[i - size + 1: i + 1] = [mID] * size
                    return i - size + 1
        return -1

    def free(self, mID: int) -> int:
        cnt = 0
        for i, id in enumerate(self.a):
            if id == mID:
                cnt += 1
                self.a[i] = 0
        return cnt
```

```go [sol1-Go]
type Allocator []int

func Constructor(n int) Allocator {
	return make([]int, n)
}

func (a Allocator) Allocate(size, mID int) int {
	cnt := 0
	for i, id := range a {
		if id > 0 {
			cnt = 0
		} else if cnt++; cnt == size {
			for j := i; j > i-size; j-- {
				a[j] = mID
			}
			return i - size + 1
		}
	}
	return -1
}

func (a Allocator) Free(mID int) (ans int) {
	for i, id := range a {
		if id == mID {
			ans++
			a[i] = 0
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$O(nq)$，其中 $q$ 为调用次数。
- 空间复杂度：$O(n)$。
