package optional

type Optionaler interface {
	Unsetter
	IsSet() bool
}

type Unsetter interface {
	Unset()
}
