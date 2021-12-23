package main

import (
	"flag"
	"log"
	"math"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	filename    = "output.html"
	numElements = 50
	port        = ":8080"
)

func main() {
	flag.StringVar(&filename, "o", filename, "Output filename (html)")
	flag.IntVar(&numElements, "c", numElements, "Number of data points")
	flag.Parse()

	http.HandleFunc("/", renderPage)
	http.ListenAndServe(port, nil)
}

func renderPage(w http.ResponseWriter, r *http.Request) {
	x1Values, y1Values := getData1()

	line := charts.NewLine()
	line.SetXAxis(x1Values).
		AddSeries("Sin(x)", y1Values)

	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "My Charts",
		Subtitle: "It's extremely easy to use, right?",
	}))

	err := line.Render(w)
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
