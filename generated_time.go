package optional

import "time"

// Time start

type TimeTester interface {
	Tester
	IfSet(callback func(value time.Time))
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

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Time) Value() time.Time {
	return b.value
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

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
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

// Time end

// Duration start

type DurationTester interface {
	Tester
	IfSet(callback func(value time.Duration))
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

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Duration) Value() time.Duration {
	return b.value
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

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
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

// Duration end
