package action

type Action int

const (
	Hit Action = iota
	Strike
)

func (a Action) ToInt() int {
	return int(a)
}
