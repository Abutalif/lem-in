package usecases

type inout interface{
}

type Builder struct {
	io inout
}

func NewBuilder(io inout) Builder {
	return Builder{
		io:io,
	}
}