package optional

type Tester interface {
	IsSet() bool
}

type Unsetter interface {
	Unset()
}

type TestUnsetter interface {
	Tester
	Unsetter
}
