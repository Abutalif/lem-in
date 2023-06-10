package entities

type Queue []Step

type Step []Move

type Move struct {
	Ant         int
	Destination string
}
