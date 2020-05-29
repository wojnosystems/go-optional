package optional

type base struct {
	isPresent bool
}

func newBase() base {
	return base{
		isPresent: false,
	}
}

func newBaseSet() base {
	return base{
		isPresent: true,
	}
}

func (b base) IsPresent() bool {
	return b.isPresent
}

func (b *base) Unset() {
	b.isPresent = false
}

func (b *base) Set() {
	b.isPresent = true
}
