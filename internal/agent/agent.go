package agent

import (
	"github.com/nihal-ramaswamy/easy21/internal/action"
	"github.com/nihal-ramaswamy/easy21/internal/episode"
	"github.com/nihal-ramaswamy/easy21/internal/state"
)

type Agent interface {
	Action(state *state.State) action.Action
	Train(episodes int)
	Fit(episode *episode.Episode)
}
