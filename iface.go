package optional

type Tester interface {
	IsSet() bool
	IfUnset(callback func())
}

type Unsetter interface {
	Unset()
}

type TestUnsetter interface {
	Tester
	Unsetter
}
