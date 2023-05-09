package models

type Anthill struct {
	Rooms  map[string]Room
	AntNum uint
}

type Room struct {
	Name        string
	X           int
	Y           int
	Visited     bool
	Kind        Kind
	Connections []*Room
}

type Kind uint8

const (
	Unknown Kind = iota
	Regular
	Start
	End
)

func New(content string) *Anthill {
	newHill := Anthill{}
	return &newHill
}
