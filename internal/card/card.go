package card

import (
	"github.com/nihal-ramaswamy/easy21/internal/color"
	"golang.org/x/exp/rand"
)

type Card struct {
	Color color.Color `json:"color"`
	Value int         `json:"value"`
}

func NewCard(forceBlack bool) *Card {
	if forceBlack {
		return &Card{
			Color: color.BLACK,
			Value: rand.Intn(10) + 1,
		}
	}
	c := color.GetRandomColor()
	value := rand.Intn(10) + 1
	if c == color.RED {
		value = -value
	}
	return &Card{
		Color: c,
		Value: value,
	}
}
