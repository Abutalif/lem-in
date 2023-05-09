package configs

type Mode uint8

const (
	Undefined Mode = iota
	CLI
	Web
)

type Configs struct {
	mode Mode
}

type ConfigWeb struct {
	host string
	port string
}

func (c *Configs) New() Configs {
	return Configs{
		CLI,
	}
}
