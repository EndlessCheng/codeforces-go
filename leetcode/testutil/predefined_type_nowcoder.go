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

//

type Interval struct {
	Start int
	End   int
}

func buildInterval(rawInterval string) (p *Interval, err error) {
	if len(rawInterval) < 5 || rawInterval[0] != '[' || rawInterval[len(rawInterval)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawInterval)
	}
	rawInterval = rawInterval[1 : len(rawInterval)-1]
	splits := strings.Split(rawInterval, ",")
	x, err := strconv.Atoi(strings.TrimSpace(splits[0]))
	if err != nil {
		return
	}
	y, err := strconv.Atoi(strings.TrimSpace(splits[1]))
	if err != nil {
		return
	}
	return &Interval{x, y}, nil
}

func (p *Interval) toRawString() string {
	return fmt.Sprintf("[%d,%d]", p.Start, p.End)
}
