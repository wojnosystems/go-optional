package optional

type base struct {
	isSet bool
}

func newBase() base {
	return base{
		isSet: false,
	}
}

func newBaseSet() base {
	return base{
		isSet: true,
	}
}

func (b base) IsSet() bool {
	return b.isSet
}

func (b *base) Unset() {
	b.isSet = false
}

func (b *base) Set() {
	b.isSet = true
}
