package main

func numWaterBottles1(numBottles, numExchange int) (ans int) {
	for numBottles >= numExchange {
		ans += numExchange
		numBottles -= numExchange - 1
	}
	return ans + numBottles
}

func numWaterBottles(numBottles, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}
