package main

import (
	"github.com/nihal-ramaswamy/easy21/internal/bar3d"
	mc_agent "github.com/nihal-ramaswamy/easy21/internal/mcAgent"
)

func main() {
	McAgent := mc_agent.NewMcAgent(0.7, 0.5)
	McAgent.Train(1000000)

	bar3d.NewBar3d(McAgent)
}
