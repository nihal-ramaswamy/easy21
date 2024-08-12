package mc_agent

import (
	"fmt"
	"math"

	"github.com/nihal-ramaswamy/easy21/internal/action"
	"github.com/nihal-ramaswamy/easy21/internal/environment"
	"github.com/nihal-ramaswamy/easy21/internal/episode"
	"github.com/nihal-ramaswamy/easy21/internal/state"
	"github.com/nihal-ramaswamy/easy21/internal/utils"
	"golang.org/x/exp/rand"
)

type McAgent struct {
	Environment    *environment.Environment `json:"environment"`
	N              [][][]float64            `json:"n"`
	Q              [][][]float64            `json:"q"`
	DiscountFactor float64                  `json:"discountFactor"`
	No             float64                  `json:"no"`
	wins           int
	iterations     int
}

func NewMcAgent(discountFactor float64, no float64) *McAgent {
	env := environment.NewEnvironment()
	return &McAgent{
		Environment:    env,
		N:              utils.Make3dArray(env.PlayerValueCount+1, env.DealerValueCount+1, env.ActionCount+1),
		Q:              utils.Make3dArray(env.PlayerValueCount+1, env.DealerValueCount+1, env.ActionCount+1),
		DiscountFactor: discountFactor,
		No:             no,
	}
}

func (m *McAgent) ValueFunction() [][]float64 {
	v := make([][]float64, m.Environment.PlayerValueCount+1)
	for i := range v {
		v[i] = make([]float64, m.Environment.DealerValueCount+1)
		for j := range v[i] {
			bestTillNow := -1.0
			for k := 0; k < len(m.Q[i][j]); k++ {
				if m.Q[i][j][k] > bestTillNow {
					bestTillNow = m.Q[i][j][k]
				}
			}
			v[i][j] = bestTillNow
		}
	}
	return v
}

func (m *McAgent) get_e(state *state.State) float64 {
	s := 0.0
	for i := 0; i < len(m.N[state.PlayerValue][state.DealerValue]); i++ {
		s += m.N[state.PlayerValue][state.DealerValue][i]
	}
	return m.No / (m.No + s)
}

func (m *McAgent) Action(s *state.State) action.Action {
	e := rand.Intn(2)
	if e == 0 {
		if rand.Intn(2) == 0 {
			return action.Hit
		} else {
			return action.Strike
		}
	} else {
		return m.chooseBestAction(s)
	}
}

func (m *McAgent) Fit(episodes []episode.Episode) {
	j := 0
	for i := range len(episodes) {
		e := episodes[i]
		s := e.State
		a := e.Action

		dealerValue := s.DealerValue
		playerValue := s.PlayerValue

		gt := 0.0
		for w := j; w < len(episodes); w++ {
			gt += episodes[w].Reward * math.Pow(m.DiscountFactor, float64(w-j))
		}

		m.N[playerValue][dealerValue][a.ToInt()] += 1
		err := gt - m.Q[playerValue][dealerValue][a.ToInt()]
		m.Q[playerValue][dealerValue][a.ToInt()] += err * (1.0 / m.N[playerValue][dealerValue][a.ToInt()])

		j += 1

	}
}

func (m *McAgent) chooseBestAction(s *state.State) action.Action {
	bestTillNow := -1.0
	bestAction := action.Hit
	for i := 0; i < len(m.Q[s.PlayerValue][s.DealerValue]); i++ {
		if m.Q[s.PlayerValue][s.DealerValue][i] > bestTillNow {
			bestTillNow = m.Q[s.PlayerValue][s.DealerValue][i]
			bestAction = action.Action(i)
		}
	}
	return bestAction
}

func (m *McAgent) Train(numEpisodes int) {
	for i := 0; i < numEpisodes; i++ {
		episodes := []episode.Episode{}

		state := m.Environment.NewStartState()

		for !state.IsTerminal {
			action := m.Action(state)
			nextState, reward := m.Environment.Step(state.Copy(), action)
			episodes = append(episodes, *episode.NewEpisode(*state, action, reward))
			state = nextState
		}

		if i%1000 == 0 {
			if m.iterations == 0 {
				continue
			}
			fmt.Printf("Episode %v, Score: %v\n", i, float64(m.wins)/float64(m.iterations)*100)
		}

		if episodes[len(episodes)-1].Reward == 1 {
			m.wins++
		}
		m.iterations++

		m.Fit(episodes)
	}
}
