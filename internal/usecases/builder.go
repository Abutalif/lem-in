package usecases

type input interface{
}

type Builder struct {
	in input
}

func NewBuilder(in input) Builder {
	return Builder{
		in:in,
	}
}