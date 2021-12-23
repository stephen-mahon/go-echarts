package main

import (
	"flag"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	filename    = "output.html"
	numElements = 50
	src         = rand.NewSource(time.Now().Unix())
	rng         = rand.New(src)
	port        = ":8080"
	wg          sync.WaitGroup
)

func main() {
	flag.StringVar(&filename, "o", filename, "Output filename (html)")
	flag.IntVar(&numElements, "c", numElements, "Number of data points")
	flag.Parse()

	startWebServer()

	hostname, _ := os.Hostname()
	log.Printf("Listing on http://%v%v", hostname, port)
	wg.Wait()
}

func startWebServer() {
	wg.Add(1)
	go func() {
		http.HandleFunc("/", renderPage)
		http.ListenAndServe(port, nil)
		wg.Done()
	}()
}

func createChart1() *charts.Line {
	x1Values, y1Values := getData1()

	line := charts.NewLine()
	line.SetXAxis(x1Values).
		AddSeries("Sin(x)", y1Values).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My Chart 1",
		Subtitle: "Sin(x)",
	}))

	return line
}

func createChart2() *charts.Line {
	x1Values, y1Values := getData2()
	_, y2Values := getData2()

	line := charts.NewLine()
	line.SetXAxis(x1Values).
		AddSeries("Random Points 1", y1Values).
		AddSeries("Random Points 2", y2Values).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My Chart 2",
		Subtitle: "Random Points",
	}))

	return line
}

func createChart3() *charts.Line {
	x1Values, y1Values := getData2()

	line := charts.NewLine()
	line.SetXAxis(x1Values).
		AddSeries("Random Points 1", y1Values).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My Chart 3",
		Subtitle: "Random Points",
	}))

	line.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "Average", Type: "average"},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{a}: {b}"}},
	)

	return line
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	chart1 := createChart1()
	chart2 := createChart2()

	page := components.NewPage()
	page.AddCharts(chart1)
	page.AddCharts(chart2)

	err := page.Render(w)
	if err != nil {
		log.Printf("Unable to render graph: %v", err)
		return
	}
}

func getData1() (x []float64, y []opts.LineData) {
	x = make([]float64, numElements)
	for i := range x {
		x[i] = float64(i)
	}

	y = make([]opts.LineData, numElements)
	for i := range y {
		y[i] = opts.LineData{Value: 10 * math.Sin(x[i])}
	}

	return
}

func getData2() (x []float64, y []opts.LineData) {
	x = make([]float64, numElements)
	for i := range x {
		x[i] = float64(i)
	}

	y = make([]opts.LineData, numElements)
	for i := range y {
		y[i] = opts.LineData{Value: rng.Float64()}
	}

	return
}
