package optional

type Presenter interface {
	Unsetter
	IsPresent() bool
}

type Unsetter interface {
	Unset()
}

type Setter interface {
	Set()
}
