package bar3d

import (
	"fmt"
	"io"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	mc_agent "github.com/nihal-ramaswamy/easy21/internal/mcAgent"
)

func NewBar3d(mcAgent *mc_agent.McAgent) {
	bar := charts.NewBar3D()

	var data []opts.Chart3DData
	value := mcAgent.ValueFunction()

	for i := 0; i < len(value); i++ {
		for j := 0; j < len(value[i]); j++ {
			fmt.Printf("Player: %v Dealer %v Value %v\n", i, j, value[i][j])
			data = append(data, opts.Chart3DData{Value: []interface{}{i, j, value[i][j]}})
		}
	}

	bar.AddSeries("Player-DealerValue", data)
	var iow io.Writer
	bar.Render(iow)
}
