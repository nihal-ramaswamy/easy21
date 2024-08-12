package episode

import (
	"github.com/nihal-ramaswamy/easy21/internal/action"
	"github.com/nihal-ramaswamy/easy21/internal/state"
)

type Episode struct {
	State  state.State   `json:"state"`
	Action action.Action `json:"action"`
	Reward float64       `json:"reward"`
}

func NewEpisode(s state.State, a action.Action, r float64) *Episode {
	return &Episode{
		State:  s,
		Action: a,
		Reward: r,
	}
}
