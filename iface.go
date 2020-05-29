package optional

type Presenter interface {
	Unsetter
	IsSet() bool
}

type Unsetter interface {
	Unset()
}

type Setter interface {
	Set()
}
