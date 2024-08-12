package state

type State struct {
	PlayerValue int  `json:"playerValue"`
	DealerValue int  `json:"dealerValue"`
	IsTerminal  bool `json:"isTerminal"`
}

func NewState(playerValue int, dealerValue int, isTerminal bool) *State {
	return &State{
		PlayerValue: playerValue,
		DealerValue: dealerValue,
		IsTerminal:  isTerminal,
	}
}

func (s *State) Copy() State {
	return State{
		PlayerValue: s.PlayerValue,
		DealerValue: s.DealerValue,
		IsTerminal:  s.IsTerminal,
	}
}
