package optional

import "time"

// Time start

type Time struct {
	state base
	value time.Time
}

func TimeUnset() Time {
	return Time{
		state: newBase(),
	}
}

func TimeFrom(value time.Time) Time {
	return Time{
		state: newBaseSet(),
		value: value,
	}
}

func (b Time) Value() time.Time {
	return b.value
}

func (b Time) IsSet() bool {
	return b.state.IsSet()
}

func (b *Time) Set(value time.Time) {
	b.state.Set()
	b.value = value
}

func (b *Time) Unset() {
	b.state.Unset()
}

func (b Time) IsEqual(compareWith Time) bool {
	leftSet := b.IsSet()
	// one is set, the other is not set, not equal
	if leftSet != compareWith.IsSet() {
		return false
	}
	// both are unset
	if !leftSet {
		return true
	}
	// both are set, compare the actual values
	return b.Value() == compareWith.Value()
}

// Time end

// Duration start

type Duration struct {
	state base
	value time.Duration
}

func DurationUnset() Duration {
	return Duration{
		state: newBase(),
	}
}

func DurationFrom(value time.Duration) Duration {
	return Duration{
		state: newBaseSet(),
		value: value,
	}
}

func (b Duration) Value() time.Duration {
	return b.value
}

func (b Duration) IsSet() bool {
	return b.state.IsSet()
}

func (b *Duration) Set(value time.Duration) {
	b.state.Set()
	b.value = value
}

func (b *Duration) Unset() {
	b.state.Unset()
}

func (b Duration) IsEqual(compareWith Duration) bool {
	leftSet := b.IsSet()
	// one is set, the other is not set, not equal
	if leftSet != compareWith.IsSet() {
		return false
	}
	// both are unset
	if !leftSet {
		return true
	}
	// both are set, compare the actual values
	return b.Value() == compareWith.Value()
}

// Duration end
