package optional

// Int32 start

type Int32 struct {
	state base
	value int32
}

func Int32Unset() Int32 {
	return Int32{
		state: newBase(),
	}
}

func Int32From(value int32) Int32 {
	return Int32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int32) Value() int32 {
	return b.value
}

func (b Int32) IsSet() bool {
	return b.state.IsSet()
}

func (b *Int32) Set(value int32) {
	b.state.Set()
	b.value = value
}

func (b *Int32) Unset() {
	b.state.Unset()
}

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

// Int32 end

// Uint64 start

type Uint64 struct {
	state base
	value uint64
}

func Uint64Unset() Uint64 {
	return Uint64{
		state: newBase(),
	}
}

func Uint64From(value uint64) Uint64 {
	return Uint64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint64) Value() uint64 {
	return b.value
}

func (b Uint64) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uint64) Set(value uint64) {
	b.state.Set()
	b.value = value
}

func (b *Uint64) Unset() {
	b.state.Unset()
}

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

// Uint64 end

// String start

type String struct {
	state base
	value string
}

func StringUnset() String {
	return String{
		state: newBase(),
	}
}

func StringFrom(value string) String {
	return String{
		state: newBaseSet(),
		value: value,
	}
}

func (b String) Value() string {
	return b.value
}

func (b String) IsSet() bool {
	return b.state.IsSet()
}

func (b *String) Set(value string) {
	b.state.Set()
	b.value = value
}

func (b *String) Unset() {
	b.state.Unset()
}

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

// String end

// Float64 start

type Float64 struct {
	state base
	value float64
}

func Float64Unset() Float64 {
	return Float64{
		state: newBase(),
	}
}

func Float64From(value float64) Float64 {
	return Float64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Float64) Value() float64 {
	return b.value
}

func (b Float64) IsSet() bool {
	return b.state.IsSet()
}

func (b *Float64) Set(value float64) {
	b.state.Set()
	b.value = value
}

func (b *Float64) Unset() {
	b.state.Unset()
}

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

// Float64 end

// Int64 start

type Int64 struct {
	state base
	value int64
}

func Int64Unset() Int64 {
	return Int64{
		state: newBase(),
	}
}

func Int64From(value int64) Int64 {
	return Int64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int64) Value() int64 {
	return b.value
}

func (b Int64) IsSet() bool {
	return b.state.IsSet()
}

func (b *Int64) Set(value int64) {
	b.state.Set()
	b.value = value
}

func (b *Int64) Unset() {
	b.state.Unset()
}

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

// Int64 end

// Uint start

type Uint struct {
	state base
	value uint
}

func UintUnset() Uint {
	return Uint{
		state: newBase(),
	}
}

func UintFrom(value uint) Uint {
	return Uint{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint) Value() uint {
	return b.value
}

func (b Uint) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uint) Set(value uint) {
	b.state.Set()
	b.value = value
}

func (b *Uint) Unset() {
	b.state.Unset()
}

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

// Uint end

// Uint32 start

type Uint32 struct {
	state base
	value uint32
}

func Uint32Unset() Uint32 {
	return Uint32{
		state: newBase(),
	}
}

func Uint32From(value uint32) Uint32 {
	return Uint32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint32) Value() uint32 {
	return b.value
}

func (b Uint32) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uint32) Set(value uint32) {
	b.state.Set()
	b.value = value
}

func (b *Uint32) Unset() {
	b.state.Unset()
}

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

// Uint32 end

// Uintptr start

type Uintptr struct {
	state base
	value uintptr
}

func UintptrUnset() Uintptr {
	return Uintptr{
		state: newBase(),
	}
}

func UintptrFrom(value uintptr) Uintptr {
	return Uintptr{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uintptr) Value() uintptr {
	return b.value
}

func (b Uintptr) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uintptr) Set(value uintptr) {
	b.state.Set()
	b.value = value
}

func (b *Uintptr) Unset() {
	b.state.Unset()
}

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

// Uintptr end

// Complex64 start

type Complex64 struct {
	state base
	value complex64
}

func Complex64Unset() Complex64 {
	return Complex64{
		state: newBase(),
	}
}

func Complex64From(value complex64) Complex64 {
	return Complex64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Complex64) Value() complex64 {
	return b.value
}

func (b Complex64) IsSet() bool {
	return b.state.IsSet()
}

func (b *Complex64) Set(value complex64) {
	b.state.Set()
	b.value = value
}

func (b *Complex64) Unset() {
	b.state.Unset()
}

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

// Complex64 end

// Int start

type Int struct {
	state base
	value int
}

func IntUnset() Int {
	return Int{
		state: newBase(),
	}
}

func IntFrom(value int) Int {
	return Int{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int) Value() int {
	return b.value
}

func (b Int) IsSet() bool {
	return b.state.IsSet()
}

func (b *Int) Set(value int) {
	b.state.Set()
	b.value = value
}

func (b *Int) Unset() {
	b.state.Unset()
}

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

// Int end

// Uint8 start

type Uint8 struct {
	state base
	value uint8
}

func Uint8Unset() Uint8 {
	return Uint8{
		state: newBase(),
	}
}

func Uint8From(value uint8) Uint8 {
	return Uint8{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint8) Value() uint8 {
	return b.value
}

func (b Uint8) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uint8) Set(value uint8) {
	b.state.Set()
	b.value = value
}

func (b *Uint8) Unset() {
	b.state.Unset()
}

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

// Uint8 end

// Byte start

type Byte struct {
	state base
	value byte
}

func ByteUnset() Byte {
	return Byte{
		state: newBase(),
	}
}

func ByteFrom(value byte) Byte {
	return Byte{
		state: newBaseSet(),
		value: value,
	}
}

func (b Byte) Value() byte {
	return b.value
}

func (b Byte) IsSet() bool {
	return b.state.IsSet()
}

func (b *Byte) Set(value byte) {
	b.state.Set()
	b.value = value
}

func (b *Byte) Unset() {
	b.state.Unset()
}

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

// Byte end

// Rune start

type Rune struct {
	state base
	value rune
}

func RuneUnset() Rune {
	return Rune{
		state: newBase(),
	}
}

func RuneFrom(value rune) Rune {
	return Rune{
		state: newBaseSet(),
		value: value,
	}
}

func (b Rune) Value() rune {
	return b.value
}

func (b Rune) IsSet() bool {
	return b.state.IsSet()
}

func (b *Rune) Set(value rune) {
	b.state.Set()
	b.value = value
}

func (b *Rune) Unset() {
	b.state.Unset()
}

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

// Rune end

// Float32 start

type Float32 struct {
	state base
	value float32
}

func Float32Unset() Float32 {
	return Float32{
		state: newBase(),
	}
}

func Float32From(value float32) Float32 {
	return Float32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Float32) Value() float32 {
	return b.value
}

func (b Float32) IsSet() bool {
	return b.state.IsSet()
}

func (b *Float32) Set(value float32) {
	b.state.Set()
	b.value = value
}

func (b *Float32) Unset() {
	b.state.Unset()
}

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

// Float32 end

// Bool start

type Bool struct {
	state base
	value bool
}

func BoolUnset() Bool {
	return Bool{
		state: newBase(),
	}
}

func BoolFrom(value bool) Bool {
	return Bool{
		state: newBaseSet(),
		value: value,
	}
}

func (b Bool) Value() bool {
	return b.value
}

func (b Bool) IsSet() bool {
	return b.state.IsSet()
}

func (b *Bool) Set(value bool) {
	b.state.Set()
	b.value = value
}

func (b *Bool) Unset() {
	b.state.Unset()
}

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

// Bool end

// Int16 start

type Int16 struct {
	state base
	value int16
}

func Int16Unset() Int16 {
	return Int16{
		state: newBase(),
	}
}

func Int16From(value int16) Int16 {
	return Int16{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int16) Value() int16 {
	return b.value
}

func (b Int16) IsSet() bool {
	return b.state.IsSet()
}

func (b *Int16) Set(value int16) {
	b.state.Set()
	b.value = value
}

func (b *Int16) Unset() {
	b.state.Unset()
}

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

// Int16 end

// Int8 start

type Int8 struct {
	state base
	value int8
}

func Int8Unset() Int8 {
	return Int8{
		state: newBase(),
	}
}

func Int8From(value int8) Int8 {
	return Int8{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int8) Value() int8 {
	return b.value
}

func (b Int8) IsSet() bool {
	return b.state.IsSet()
}

func (b *Int8) Set(value int8) {
	b.state.Set()
	b.value = value
}

func (b *Int8) Unset() {
	b.state.Unset()
}

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

// Int8 end

// Uint16 start

type Uint16 struct {
	state base
	value uint16
}

func Uint16Unset() Uint16 {
	return Uint16{
		state: newBase(),
	}
}

func Uint16From(value uint16) Uint16 {
	return Uint16{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint16) Value() uint16 {
	return b.value
}

func (b Uint16) IsSet() bool {
	return b.state.IsSet()
}

func (b *Uint16) Set(value uint16) {
	b.state.Set()
	b.value = value
}

func (b *Uint16) Unset() {
	b.state.Unset()
}

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

// Uint16 end

// Complex128 start

type Complex128 struct {
	state base
	value complex128
}

func Complex128Unset() Complex128 {
	return Complex128{
		state: newBase(),
	}
}

func Complex128From(value complex128) Complex128 {
	return Complex128{
		state: newBaseSet(),
		value: value,
	}
}

func (b Complex128) Value() complex128 {
	return b.value
}

func (b Complex128) IsSet() bool {
	return b.state.IsSet()
}

func (b *Complex128) Set(value complex128) {
	b.state.Set()
	b.value = value
}

func (b *Complex128) Unset() {
	b.state.Unset()
}

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

// Complex128 end
