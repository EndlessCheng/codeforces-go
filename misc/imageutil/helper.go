package imageutil

import (
	"strconv"
	"strings"
)

func ParseArray(rawArray string) (res []float64) {
	rawArray = strings.TrimSpace(rawArray)
	if rawArray == "" {
		return nil
	}
	if rawArray[0] == '[' && rawArray[len(rawArray)-1] == ']' {
		rawArray = rawArray[1 : len(rawArray)-1]
	}
	splits := strings.Fields(rawArray)
	for _, s := range splits {
		v, _ := strconv.Atoi(s)
		res = append(res, float64(v))
	}
	return
}

func ParsePoints(rawPoints string) (x []float64, y []float64) {
	rawPoints = strings.TrimSpace(rawPoints)
	if rawPoints == "" {
		return
	}
	splits := strings.Fields(rawPoints)
	for i, s := range splits {
		v, _ := strconv.Atoi(s)
		if i%2 == 0 {
			x = append(x, float64(v))
		} else {
			y = append(y, float64(v))
		}
	}
	return
}
