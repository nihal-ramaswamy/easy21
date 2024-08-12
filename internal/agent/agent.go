package agent

import "github.com/nihal-ramaswamy/easy21/internal/environment"

type Agent interface {
	Train(episodes int)
	ValueFunction() [][]float64
	GetEnvironment() *environment.Environment
}
