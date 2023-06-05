package usecases

type output interface{}

type Presenter struct {
	out output
}

func NewPresenter(out output) Presenter {
	return Presenter{
		out: out,
	}
}
