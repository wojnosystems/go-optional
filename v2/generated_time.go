package optional

import "time"

// Time start

type TimeTester interface {
	Tester
	IfSet(callback func(value time.Time))
	IfSetElse(setCallback func(value time.Time), unsetCallback func())
}

// Time Creates an optional with value type: time.Time
type Time struct {
	state base
	value time.Time
}

// TimeUnset: creates a new optional which is unset when created. Calling IsSet will return false
func TimeUnset() Time {
	return Time{
		state: newBase(),
	}
}

// TimeFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func TimeFrom(value time.Time) Time {
	return Time{
		state: newBaseSet(),
		value: value,
	}
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Time) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Time) Set(value time.Time) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Time) Unset() {
	b.state.Unset()
}

// IfSet: call-back method for code-flow safety. If IsSet is true, callback will be executed with
// the value in the Optional. Otherwise, the callback will not be called.
// This should eventually replace all instances of: Value and ValueWithOK.
func (b Time) IfSet(callback func(value time.Time)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Time) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// IfSetElse: call-back method for code-flow safety. If IsSet is true, the first argument is called, if false, the second argument is called.
func (b Time) IfSetElse(setCallback func(value time.Time), unsetCallback func()) {
	if b.IsSet() {
		setCallback(b.value)
	} else {
		unsetCallback()
	}
}

// Time end

// Duration start

type DurationTester interface {
	Tester
	IfSet(callback func(value time.Duration))
	IfSetElse(setCallback func(value time.Duration), unsetCallback func())
}

// Duration Creates an optional with value type: time.Duration
type Duration struct {
	state base
	value time.Duration
}

// DurationUnset: creates a new optional which is unset when created. Calling IsSet will return false
func DurationUnset() Duration {
	return Duration{
		state: newBase(),
	}
}

// DurationFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func DurationFrom(value time.Duration) Duration {
	return Duration{
		state: newBaseSet(),
		value: value,
	}
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Duration) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Duration) Set(value time.Duration) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Duration) Unset() {
	b.state.Unset()
}

// IfSet: call-back method for code-flow safety. If IsSet is true, callback will be executed with
// the value in the Optional. Otherwise, the callback will not be called.
// This should eventually replace all instances of: Value and ValueWithOK.
func (b Duration) IfSet(callback func(value time.Duration)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Duration) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// IfSetElse: call-back method for code-flow safety. If IsSet is true, the first argument is called, if false, the second argument is called.
func (b Duration) IfSetElse(setCallback func(value time.Duration), unsetCallback func()) {
	if b.IsSet() {
		setCallback(b.value)
	} else {
		unsetCallback()
	}
}

// Duration end
