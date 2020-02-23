package imageutil

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/wcharczuk/go-chart"
	"os"
	"time"
)

func mins(vals []float64) float64 {
	ans := vals[0]
	for _, val := range vals[1:] {
		if val < ans {
			ans = val
		}
	}
	return ans
}

func maxs(vals []float64) float64 {
	ans := vals[0]
	for _, val := range vals[1:] {
		if val > ans {
			ans = val
		}
	}
	return ans
}

func DrawChart(x, y []float64, showLine bool) error {
	now := time.Now().Format("20060102150405.000")
	f, err := os.Create(fmt.Sprintf("image.%s.png", now))
	if err != nil {
		return err
	}

	xSet := map[float64]struct{}{}
	for _, v := range x {
		xSet[v] = struct{}{}
	}
	xTicks := make([]chart.Tick, 0, len(xSet))
	for v := range xSet {
		xTicks = append(xTicks, chart.Tick{Value: v, Label: fmt.Sprintf("%.f", v)}) // x 视作 int
	}
	minX, maxX := mins(x), maxs(x)

	yTicks := []chart.Tick{{Value: 0, Label: "0"}}
	minY, maxY := mins(y), maxs(y)
	if minY < 0 {
		yTicks = append(yTicks, chart.Tick{Value: minY, Label: fmt.Sprintf("%.f", minY)}) // y 视作 int
	}
	if maxY > 0 {
		yTicks = append(yTicks, chart.Tick{Value: maxY, Label: fmt.Sprintf("%.f", maxY)}) // y 视作 int
	}

	graphWidth, graphHeight := 0, 0
	deltaX := maxX - minX
	deltaY := maxY - minY
	whRate := deltaX / deltaY
	if 0.2 <= whRate && whRate <= 5 {
		if whRate < 1 {
			graphHeight = 800
			graphWidth = int(float64(graphHeight) * whRate)
		} else {
			graphWidth = 800
			graphHeight = int(float64(graphWidth) / whRate)
		}
	}
	if graphWidth == 0 {
		graphWidth = chart.DefaultChartWidth
	}
	if graphHeight == 0 {
		graphHeight = 800
	}

	data := chart.ContinuousSeries{
		Style: chart.Style{
			DotWidth: 5,
		},
		XValues: x,
		YValues: y,
	}
	if !showLine {
		data.Style.StrokeWidth = chart.Disabled
	}

	yAnno := make([]chart.Value2, len(x))
	for i := range yAnno {
		yAnno[i] = chart.Value2{
			XValue: x[i],
			YValue: y[i],
			Label:  fmt.Sprintf("%.f", y[i]), // y 视作 int
		}
	}

	graph := chart.Chart{
		Width:  graphWidth,
		Height: graphHeight,
		XAxis: chart.XAxis{
			Ticks: xTicks,
		},
		YAxis: chart.YAxis{
			AxisType: chart.YAxisSecondary,
			Ticks:    yTicks,
		},
		Series: []chart.Series{
			data,
			chart.AnnotationSeries{Annotations: yAnno},
		},
	}

	if err := graph.Render(chart.PNG, f); err != nil {
		return err
	}
	f.Close()
	open.Run(f.Name())
	return nil
}

func genRange(st, end int) (res []float64) {
	for i := st; i <= end; i++ {
		res = append(res, float64(i))
	}
	return
}

func DrawArray(a []float64) error {
	return DrawChart(genRange(0, len(a)-1), a, true)
}

func DrawPrefixSum(a []float64) error {
	n := len(a)
	sum := make([]float64, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	return DrawChart(genRange(0, len(a)), sum, true)
}

func DrawWithRawPoints(rawPoints string, showLine bool) error {
	x, y := ParsePoints(rawPoints)
	return DrawChart(x, y, showLine)
}
