package bar3d

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/nihal-ramaswamy/easy21/internal/agent"
)

func NewBar3d(agent agent.Agent) {
	bar := charts.NewBar3D()

	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Player-Dealer Value Function",
		}),
	)

	var data []opts.Chart3DData
	value := agent.ValueFunction()

	for i := 0; i < len(value); i++ {
		for j := 0; j < len(value[i]); j++ {
			data = append(data, opts.Chart3DData{Value: []interface{}{i, j, value[i][j]}})
		}
	}

	bar.AddSeries("Player-DealerValue", data, charts.WithBar3DChartOpts(opts.Bar3DChart{Shading: "lambert"}))

	page := components.NewPage()
	page.AddCharts(
		bar,
	)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = os.Mkdir("html", os.ModePerm)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	fileDir := filepath.Join(cwd, "html", filepath.Base("bar3d.html"))
	fmt.Printf("Saving file to %v\n", fileDir)
	f, err := os.Create(fileDir)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	page.Render(io.MultiWriter(f))
}
