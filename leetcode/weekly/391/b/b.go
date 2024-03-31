package main

// https://space.bilibili.com/206214
func maxBottlesDrunk(numBottles, numExchange int) int {
	ans := numBottles
	for numBottles >= numExchange { // 有足够的空瓶
		ans++ // 用 numExchange 个空瓶交换，然后喝掉，产生一个新的空瓶
		numBottles += 1 - numExchange
		numExchange++
	}
	return ans
}
