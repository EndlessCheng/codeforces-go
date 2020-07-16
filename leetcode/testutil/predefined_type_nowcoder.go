package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func buildPoint(rawPoint string) (p *Point, err error) {
	if len(rawPoint) < 5 || rawPoint[0] != '(' || rawPoint[len(rawPoint)-1] != ')' {
		return nil, fmt.Errorf("invalid test data %s", rawPoint)
	}
	rawPoint = rawPoint[1 : len(rawPoint)-1]
	splits := strings.Split(rawPoint, ",")
	x, err := strconv.Atoi(strings.TrimSpace(splits[0]))
	if err != nil {
		return
	}
	y, err := strconv.Atoi(strings.TrimSpace(splits[1]))
	if err != nil {
		return
	}
	return &Point{x, y}, nil
}

func (p *Point) toRawString() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
