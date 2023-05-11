package entities

type Anthill struct {
	Rooms  []*Room
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
