package environment

import (
	"github.com/nihal-ramaswamy/easy21/internal/action"
	"github.com/nihal-ramaswamy/easy21/internal/card"
	"github.com/nihal-ramaswamy/easy21/internal/state"
)

type Environment struct {
	PlayerValueCount int `json:"playerValueCount"`
	DealerValueCount int `json:"dealerValueCount"`
	ActionCount      int `json:"actionCount"`
}

func NewEnvironment() *Environment {
	return &Environment{
		PlayerValueCount: 21,
		DealerValueCount: 10,
		ActionCount:      2,
	}
}

func (e *Environment) NewStartState() *state.State {
	return state.NewState(
		card.NewCard(true).Value,
		card.NewCard(true).Value,
		false)
}

func (e *Environment) Step(state state.State, a action.Action) (*state.State, float64) {
	currentState := state.Copy()
	reward := 0.0

	switch a {
	case action.Hit:
		card := card.NewCard(false)
		currentState.PlayerValue += card.Value
		if currentState.PlayerValue > 21 {
			currentState.IsTerminal = true
			reward = -1
		}
	case action.Strike:
		for !currentState.IsTerminal {
			currentState.DealerValue += card.NewCard(false).Value
			if currentState.DealerValue > 21 {
				currentState.IsTerminal = true
				reward = 1
			} else if currentState.DealerValue >= 17 {
				currentState.IsTerminal = true
				if currentState.PlayerValue > currentState.DealerValue {
					reward = 1
				} else if currentState.PlayerValue < currentState.DealerValue {
					reward = -1
				}
			}
		}
	}

	return &currentState, reward
}
