package optional

// Bool start

type BoolTester interface {
	Tester
	IfSet(callback func(value bool))
}

// Bool Creates an optional with value type: bool
type Bool struct {
	state base
	value bool
}

// BoolUnset: creates a new optional which is unset when created. Calling IsSet will return false
func BoolUnset() Bool {
	return Bool{
		state: newBase(),
	}
}

// BoolFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func BoolFrom(value bool) Bool {
	return Bool{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Bool) Value() bool {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Bool) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Bool) Set(value bool) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Bool) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Bool) IsEqual(compareWith Bool) bool {
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
func (b Bool) IfSet(callback func(value bool)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Bool) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Bool end

// Int64 start

type Int64Tester interface {
	Tester
	IfSet(callback func(value int64))
}

// Int64 Creates an optional with value type: int64
type Int64 struct {
	state base
	value int64
}

// Int64Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Int64Unset() Int64 {
	return Int64{
		state: newBase(),
	}
}

// Int64From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Int64From(value int64) Int64 {
	return Int64{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Int64) Value() int64 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Int64) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Int64) Set(value int64) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Int64) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Int64) IsEqual(compareWith Int64) bool {
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
func (b Int64) IfSet(callback func(value int64)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Int64) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Int64 end

// Uint8 start

type Uint8Tester interface {
	Tester
	IfSet(callback func(value uint8))
}

// Uint8 Creates an optional with value type: uint8
type Uint8 struct {
	state base
	value uint8
}

// Uint8Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Uint8Unset() Uint8 {
	return Uint8{
		state: newBase(),
	}
}

// Uint8From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Uint8From(value uint8) Uint8 {
	return Uint8{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uint8) Value() uint8 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uint8) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uint8) Set(value uint8) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uint8) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uint8) IsEqual(compareWith Uint8) bool {
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
func (b Uint8) IfSet(callback func(value uint8)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uint8) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uint8 end

// Float32 start

type Float32Tester interface {
	Tester
	IfSet(callback func(value float32))
}

// Float32 Creates an optional with value type: float32
type Float32 struct {
	state base
	value float32
}

// Float32Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Float32Unset() Float32 {
	return Float32{
		state: newBase(),
	}
}

// Float32From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Float32From(value float32) Float32 {
	return Float32{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Float32) Value() float32 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Float32) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Float32) Set(value float32) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Float32) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Float32) IsEqual(compareWith Float32) bool {
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
func (b Float32) IfSet(callback func(value float32)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Float32) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Float32 end

// Complex64 start

type Complex64Tester interface {
	Tester
	IfSet(callback func(value complex64))
}

// Complex64 Creates an optional with value type: complex64
type Complex64 struct {
	state base
	value complex64
}

// Complex64Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Complex64Unset() Complex64 {
	return Complex64{
		state: newBase(),
	}
}

// Complex64From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Complex64From(value complex64) Complex64 {
	return Complex64{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Complex64) Value() complex64 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Complex64) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Complex64) Set(value complex64) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Complex64) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Complex64) IsEqual(compareWith Complex64) bool {
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
func (b Complex64) IfSet(callback func(value complex64)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Complex64) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Complex64 end

// Int16 start

type Int16Tester interface {
	Tester
	IfSet(callback func(value int16))
}

// Int16 Creates an optional with value type: int16
type Int16 struct {
	state base
	value int16
}

// Int16Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Int16Unset() Int16 {
	return Int16{
		state: newBase(),
	}
}

// Int16From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Int16From(value int16) Int16 {
	return Int16{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Int16) Value() int16 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Int16) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Int16) Set(value int16) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Int16) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Int16) IsEqual(compareWith Int16) bool {
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
func (b Int16) IfSet(callback func(value int16)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Int16) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Int16 end

// Uint start

type UintTester interface {
	Tester
	IfSet(callback func(value uint))
}

// Uint Creates an optional with value type: uint
type Uint struct {
	state base
	value uint
}

// UintUnset: creates a new optional which is unset when created. Calling IsSet will return false
func UintUnset() Uint {
	return Uint{
		state: newBase(),
	}
}

// UintFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func UintFrom(value uint) Uint {
	return Uint{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uint) Value() uint {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uint) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uint) Set(value uint) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uint) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uint) IsEqual(compareWith Uint) bool {
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
func (b Uint) IfSet(callback func(value uint)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uint) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uint end

// Uint64 start

type Uint64Tester interface {
	Tester
	IfSet(callback func(value uint64))
}

// Uint64 Creates an optional with value type: uint64
type Uint64 struct {
	state base
	value uint64
}

// Uint64Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Uint64Unset() Uint64 {
	return Uint64{
		state: newBase(),
	}
}

// Uint64From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Uint64From(value uint64) Uint64 {
	return Uint64{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uint64) Value() uint64 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uint64) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uint64) Set(value uint64) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uint64) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uint64) IsEqual(compareWith Uint64) bool {
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
func (b Uint64) IfSet(callback func(value uint64)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uint64) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uint64 end

// Uint16 start

type Uint16Tester interface {
	Tester
	IfSet(callback func(value uint16))
}

// Uint16 Creates an optional with value type: uint16
type Uint16 struct {
	state base
	value uint16
}

// Uint16Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Uint16Unset() Uint16 {
	return Uint16{
		state: newBase(),
	}
}

// Uint16From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Uint16From(value uint16) Uint16 {
	return Uint16{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uint16) Value() uint16 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uint16) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uint16) Set(value uint16) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uint16) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uint16) IsEqual(compareWith Uint16) bool {
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
func (b Uint16) IfSet(callback func(value uint16)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uint16) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uint16 end

// Int start

type IntTester interface {
	Tester
	IfSet(callback func(value int))
}

// Int Creates an optional with value type: int
type Int struct {
	state base
	value int
}

// IntUnset: creates a new optional which is unset when created. Calling IsSet will return false
func IntUnset() Int {
	return Int{
		state: newBase(),
	}
}

// IntFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func IntFrom(value int) Int {
	return Int{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Int) Value() int {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Int) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Int) Set(value int) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Int) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Int) IsEqual(compareWith Int) bool {
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
func (b Int) IfSet(callback func(value int)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Int) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Int end

// Int8 start

type Int8Tester interface {
	Tester
	IfSet(callback func(value int8))
}

// Int8 Creates an optional with value type: int8
type Int8 struct {
	state base
	value int8
}

// Int8Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Int8Unset() Int8 {
	return Int8{
		state: newBase(),
	}
}

// Int8From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Int8From(value int8) Int8 {
	return Int8{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Int8) Value() int8 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Int8) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Int8) Set(value int8) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Int8) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Int8) IsEqual(compareWith Int8) bool {
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
func (b Int8) IfSet(callback func(value int8)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Int8) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Int8 end

// Rune start

type RuneTester interface {
	Tester
	IfSet(callback func(value rune))
}

// Rune Creates an optional with value type: rune
type Rune struct {
	state base
	value rune
}

// RuneUnset: creates a new optional which is unset when created. Calling IsSet will return false
func RuneUnset() Rune {
	return Rune{
		state: newBase(),
	}
}

// RuneFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func RuneFrom(value rune) Rune {
	return Rune{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Rune) Value() rune {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Rune) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Rune) Set(value rune) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Rune) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Rune) IsEqual(compareWith Rune) bool {
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
func (b Rune) IfSet(callback func(value rune)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Rune) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Rune end

// String start

type StringTester interface {
	Tester
	IfSet(callback func(value string))
}

// String Creates an optional with value type: string
type String struct {
	state base
	value string
}

// StringUnset: creates a new optional which is unset when created. Calling IsSet will return false
func StringUnset() String {
	return String{
		state: newBase(),
	}
}

// StringFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func StringFrom(value string) String {
	return String{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b String) Value() string {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b String) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *String) Set(value string) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *String) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b String) IsEqual(compareWith String) bool {
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
func (b String) IfSet(callback func(value string)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b String) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// String end

// Uintptr start

type UintptrTester interface {
	Tester
	IfSet(callback func(value uintptr))
}

// Uintptr Creates an optional with value type: uintptr
type Uintptr struct {
	state base
	value uintptr
}

// UintptrUnset: creates a new optional which is unset when created. Calling IsSet will return false
func UintptrUnset() Uintptr {
	return Uintptr{
		state: newBase(),
	}
}

// UintptrFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func UintptrFrom(value uintptr) Uintptr {
	return Uintptr{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uintptr) Value() uintptr {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uintptr) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uintptr) Set(value uintptr) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uintptr) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uintptr) IsEqual(compareWith Uintptr) bool {
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
func (b Uintptr) IfSet(callback func(value uintptr)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uintptr) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uintptr end

// Complex128 start

type Complex128Tester interface {
	Tester
	IfSet(callback func(value complex128))
}

// Complex128 Creates an optional with value type: complex128
type Complex128 struct {
	state base
	value complex128
}

// Complex128Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Complex128Unset() Complex128 {
	return Complex128{
		state: newBase(),
	}
}

// Complex128From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Complex128From(value complex128) Complex128 {
	return Complex128{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Complex128) Value() complex128 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Complex128) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Complex128) Set(value complex128) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Complex128) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Complex128) IsEqual(compareWith Complex128) bool {
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
func (b Complex128) IfSet(callback func(value complex128)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Complex128) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Complex128 end

// Int32 start

type Int32Tester interface {
	Tester
	IfSet(callback func(value int32))
}

// Int32 Creates an optional with value type: int32
type Int32 struct {
	state base
	value int32
}

// Int32Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Int32Unset() Int32 {
	return Int32{
		state: newBase(),
	}
}

// Int32From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Int32From(value int32) Int32 {
	return Int32{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Int32) Value() int32 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Int32) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Int32) Set(value int32) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Int32) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Int32) IsEqual(compareWith Int32) bool {
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
func (b Int32) IfSet(callback func(value int32)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Int32) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Int32 end

// Uint32 start

type Uint32Tester interface {
	Tester
	IfSet(callback func(value uint32))
}

// Uint32 Creates an optional with value type: uint32
type Uint32 struct {
	state base
	value uint32
}

// Uint32Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Uint32Unset() Uint32 {
	return Uint32{
		state: newBase(),
	}
}

// Uint32From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Uint32From(value uint32) Uint32 {
	return Uint32{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Uint32) Value() uint32 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Uint32) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Uint32) Set(value uint32) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Uint32) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Uint32) IsEqual(compareWith Uint32) bool {
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
func (b Uint32) IfSet(callback func(value uint32)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Uint32) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Uint32 end

// Byte start

type ByteTester interface {
	Tester
	IfSet(callback func(value byte))
}

// Byte Creates an optional with value type: byte
type Byte struct {
	state base
	value byte
}

// ByteUnset: creates a new optional which is unset when created. Calling IsSet will return false
func ByteUnset() Byte {
	return Byte{
		state: newBase(),
	}
}

// ByteFrom: creates a new Optional, but loaded with a value. Calling IsSet will return true
func ByteFrom(value byte) Byte {
	return Byte{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Byte) Value() byte {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Byte) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Byte) Set(value byte) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Byte) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Byte) IsEqual(compareWith Byte) bool {
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
func (b Byte) IfSet(callback func(value byte)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Byte) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Byte end

// Float64 start

type Float64Tester interface {
	Tester
	IfSet(callback func(value float64))
}

// Float64 Creates an optional with value type: float64
type Float64 struct {
	state base
	value float64
}

// Float64Unset: creates a new optional which is unset when created. Calling IsSet will return false
func Float64Unset() Float64 {
	return Float64{
		state: newBase(),
	}
}

// Float64From: creates a new Optional, but loaded with a value. Calling IsSet will return true
func Float64From(value float64) Float64 {
	return Float64{
		state: newBaseSet(),
		value: value,
	}
}

// Value: Gets the value of the optional item
// @deprecated. Do not use this method. Will be removed in version 2. Use IfSet instead.
func (b Float64) Value() float64 {
	return b.value
}

// IsSet: returns true if there is a valid value or false if a value was not set or Unset
func (b Float64) IsSet() bool {
	return b.state.IsSet()
}

// Set: updates the value stored within this object and marks the value as valid. Calling IsSet after calling Set will return true.
func (b *Float64) Set(value float64) {
	b.state.Set()
	b.value = value
}

// Unset: marks the Optional as no longer having a valid value. Calling IsSet after Unset returns false.
func (b *Float64) Unset() {
	b.state.Unset()
}

// IsEqual: Compares two Optionals for not only the correct value, but also whether both are set or not
// Note: this might be a conflation of concerns, but this is a useful tool for testing.
// If both are not set, this returns true, regardless of the internal value. If both are set, then the values must be equal. If either is set but the other is unset, this returns false.
func (b Float64) IsEqual(compareWith Float64) bool {
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
func (b Float64) IfSet(callback func(value float64)) {
	if b.IsSet() {
		callback(b.value)
	}
}

// IfSet: call-back method for code-flow safety. If IsSet is false, callback will be executed. No value is exposed as this item is not set. Otherwise, the callback will not be called.
// This should eventually replace all instances of: if !o.IsSet()
func (b Float64) IfUnset(callback func()) {
	if !b.IsSet() {
		callback()
	}
}

// Float64 end
