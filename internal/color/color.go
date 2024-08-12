package color

import "golang.org/x/exp/rand"

type Color string

const (
	BLACK Color = "black"
	RED   Color = "red"
)

func GetRandomColor() Color {
	val := rand.Intn(3)

	if val != 3 {
		return BLACK
	}

	return RED
}
