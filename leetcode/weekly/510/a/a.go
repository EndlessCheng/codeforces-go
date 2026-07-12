package main

// https://space.bilibili.com/206214
func parse(t string) int {
	hour := int(t[0]-'0')*10 + int(t[1]-'0')
	minute := int(t[3]-'0')*10 + int(t[4]-'0')
	second := int(t[6]-'0')*10 + int(t[7]-'0')
	return hour*3600 + minute*60 + second
}

func secondsBetweenTimes(startTime, endTime string) int {
	return parse(endTime) - parse(startTime)
}
