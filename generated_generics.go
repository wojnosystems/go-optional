package present

type Float32 struct {
	state base
	value float32
}

func NewFloat32() Float32 {
	return Float32{
		state: newBase(),
	}
}

func NewFloat32Set(value float32) Float32 {
	return Float32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Float32) Value() float32 {
	return b.value
}

func (b Float32) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Float32) ValueWithPresent() (value float32, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Float32) Set(value float32) {
	b.state.Set()
	b.value = value
}

func (b *Float32) Unset() {
	b.state.Unset()
}

func (b Float32) IsEqual(compareWith Float32) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Complex128 struct {
	state base
	value complex128
}

func NewComplex128() Complex128 {
	return Complex128{
		state: newBase(),
	}
}

func NewComplex128Set(value complex128) Complex128 {
	return Complex128{
		state: newBaseSet(),
		value: value,
	}
}

func (b Complex128) Value() complex128 {
	return b.value
}

func (b Complex128) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Complex128) ValueWithPresent() (value complex128, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Complex128) Set(value complex128) {
	b.state.Set()
	b.value = value
}

func (b *Complex128) Unset() {
	b.state.Unset()
}

func (b Complex128) IsEqual(compareWith Complex128) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Complex64 struct {
	state base
	value complex64
}

func NewComplex64() Complex64 {
	return Complex64{
		state: newBase(),
	}
}

func NewComplex64Set(value complex64) Complex64 {
	return Complex64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Complex64) Value() complex64 {
	return b.value
}

func (b Complex64) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Complex64) ValueWithPresent() (value complex64, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Complex64) Set(value complex64) {
	b.state.Set()
	b.value = value
}

func (b *Complex64) Unset() {
	b.state.Unset()
}

func (b Complex64) IsEqual(compareWith Complex64) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Int struct {
	state base
	value int
}

func NewInt() Int {
	return Int{
		state: newBase(),
	}
}

func NewIntSet(value int) Int {
	return Int{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int) Value() int {
	return b.value
}

func (b Int) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Int) ValueWithPresent() (value int, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Int) Set(value int) {
	b.state.Set()
	b.value = value
}

func (b *Int) Unset() {
	b.state.Unset()
}

func (b Int) IsEqual(compareWith Int) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Int16 struct {
	state base
	value int16
}

func NewInt16() Int16 {
	return Int16{
		state: newBase(),
	}
}

func NewInt16Set(value int16) Int16 {
	return Int16{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int16) Value() int16 {
	return b.value
}

func (b Int16) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Int16) ValueWithPresent() (value int16, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Int16) Set(value int16) {
	b.state.Set()
	b.value = value
}

func (b *Int16) Unset() {
	b.state.Unset()
}

func (b Int16) IsEqual(compareWith Int16) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Byte struct {
	state base
	value byte
}

func NewByte() Byte {
	return Byte{
		state: newBase(),
	}
}

func NewByteSet(value byte) Byte {
	return Byte{
		state: newBaseSet(),
		value: value,
	}
}

func (b Byte) Value() byte {
	return b.value
}

func (b Byte) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Byte) ValueWithPresent() (value byte, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Byte) Set(value byte) {
	b.state.Set()
	b.value = value
}

func (b *Byte) Unset() {
	b.state.Unset()
}

func (b Byte) IsEqual(compareWith Byte) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type String struct {
	state base
	value string
}

func NewString() String {
	return String{
		state: newBase(),
	}
}

func NewStringSet(value string) String {
	return String{
		state: newBaseSet(),
		value: value,
	}
}

func (b String) Value() string {
	return b.value
}

func (b String) IsPresent() bool {
	return b.state.IsPresent()
}

func (b String) ValueWithPresent() (value string, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *String) Set(value string) {
	b.state.Set()
	b.value = value
}

func (b *String) Unset() {
	b.state.Unset()
}

func (b String) IsEqual(compareWith String) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Rune struct {
	state base
	value rune
}

func NewRune() Rune {
	return Rune{
		state: newBase(),
	}
}

func NewRuneSet(value rune) Rune {
	return Rune{
		state: newBaseSet(),
		value: value,
	}
}

func (b Rune) Value() rune {
	return b.value
}

func (b Rune) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Rune) ValueWithPresent() (value rune, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Rune) Set(value rune) {
	b.state.Set()
	b.value = value
}

func (b *Rune) Unset() {
	b.state.Unset()
}

func (b Rune) IsEqual(compareWith Rune) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uintptr struct {
	state base
	value uintptr
}

func NewUintptr() Uintptr {
	return Uintptr{
		state: newBase(),
	}
}

func NewUintptrSet(value uintptr) Uintptr {
	return Uintptr{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uintptr) Value() uintptr {
	return b.value
}

func (b Uintptr) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uintptr) ValueWithPresent() (value uintptr, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uintptr) Set(value uintptr) {
	b.state.Set()
	b.value = value
}

func (b *Uintptr) Unset() {
	b.state.Unset()
}

func (b Uintptr) IsEqual(compareWith Uintptr) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Bool struct {
	state base
	value bool
}

func NewBool() Bool {
	return Bool{
		state: newBase(),
	}
}

func NewBoolSet(value bool) Bool {
	return Bool{
		state: newBaseSet(),
		value: value,
	}
}

func (b Bool) Value() bool {
	return b.value
}

func (b Bool) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Bool) ValueWithPresent() (value bool, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Bool) Set(value bool) {
	b.state.Set()
	b.value = value
}

func (b *Bool) Unset() {
	b.state.Unset()
}

func (b Bool) IsEqual(compareWith Bool) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Int8 struct {
	state base
	value int8
}

func NewInt8() Int8 {
	return Int8{
		state: newBase(),
	}
}

func NewInt8Set(value int8) Int8 {
	return Int8{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int8) Value() int8 {
	return b.value
}

func (b Int8) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Int8) ValueWithPresent() (value int8, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Int8) Set(value int8) {
	b.state.Set()
	b.value = value
}

func (b *Int8) Unset() {
	b.state.Unset()
}

func (b Int8) IsEqual(compareWith Int8) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uint struct {
	state base
	value uint
}

func NewUint() Uint {
	return Uint{
		state: newBase(),
	}
}

func NewUintSet(value uint) Uint {
	return Uint{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint) Value() uint {
	return b.value
}

func (b Uint) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uint) ValueWithPresent() (value uint, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uint) Set(value uint) {
	b.state.Set()
	b.value = value
}

func (b *Uint) Unset() {
	b.state.Unset()
}

func (b Uint) IsEqual(compareWith Uint) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uint8 struct {
	state base
	value uint8
}

func NewUint8() Uint8 {
	return Uint8{
		state: newBase(),
	}
}

func NewUint8Set(value uint8) Uint8 {
	return Uint8{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint8) Value() uint8 {
	return b.value
}

func (b Uint8) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uint8) ValueWithPresent() (value uint8, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uint8) Set(value uint8) {
	b.state.Set()
	b.value = value
}

func (b *Uint8) Unset() {
	b.state.Unset()
}

func (b Uint8) IsEqual(compareWith Uint8) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uint16 struct {
	state base
	value uint16
}

func NewUint16() Uint16 {
	return Uint16{
		state: newBase(),
	}
}

func NewUint16Set(value uint16) Uint16 {
	return Uint16{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint16) Value() uint16 {
	return b.value
}

func (b Uint16) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uint16) ValueWithPresent() (value uint16, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uint16) Set(value uint16) {
	b.state.Set()
	b.value = value
}

func (b *Uint16) Unset() {
	b.state.Unset()
}

func (b Uint16) IsEqual(compareWith Uint16) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Float64 struct {
	state base
	value float64
}

func NewFloat64() Float64 {
	return Float64{
		state: newBase(),
	}
}

func NewFloat64Set(value float64) Float64 {
	return Float64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Float64) Value() float64 {
	return b.value
}

func (b Float64) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Float64) ValueWithPresent() (value float64, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Float64) Set(value float64) {
	b.state.Set()
	b.value = value
}

func (b *Float64) Unset() {
	b.state.Unset()
}

func (b Float64) IsEqual(compareWith Float64) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Int64 struct {
	state base
	value int64
}

func NewInt64() Int64 {
	return Int64{
		state: newBase(),
	}
}

func NewInt64Set(value int64) Int64 {
	return Int64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int64) Value() int64 {
	return b.value
}

func (b Int64) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Int64) ValueWithPresent() (value int64, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Int64) Set(value int64) {
	b.state.Set()
	b.value = value
}

func (b *Int64) Unset() {
	b.state.Unset()
}

func (b Int64) IsEqual(compareWith Int64) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Int32 struct {
	state base
	value int32
}

func NewInt32() Int32 {
	return Int32{
		state: newBase(),
	}
}

func NewInt32Set(value int32) Int32 {
	return Int32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Int32) Value() int32 {
	return b.value
}

func (b Int32) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Int32) ValueWithPresent() (value int32, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Int32) Set(value int32) {
	b.state.Set()
	b.value = value
}

func (b *Int32) Unset() {
	b.state.Unset()
}

func (b Int32) IsEqual(compareWith Int32) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uint64 struct {
	state base
	value uint64
}

func NewUint64() Uint64 {
	return Uint64{
		state: newBase(),
	}
}

func NewUint64Set(value uint64) Uint64 {
	return Uint64{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint64) Value() uint64 {
	return b.value
}

func (b Uint64) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uint64) ValueWithPresent() (value uint64, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uint64) Set(value uint64) {
	b.state.Set()
	b.value = value
}

func (b *Uint64) Unset() {
	b.state.Unset()
}

func (b Uint64) IsEqual(compareWith Uint64) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}

type Uint32 struct {
	state base
	value uint32
}

func NewUint32() Uint32 {
	return Uint32{
		state: newBase(),
	}
}

func NewUint32Set(value uint32) Uint32 {
	return Uint32{
		state: newBaseSet(),
		value: value,
	}
}

func (b Uint32) Value() uint32 {
	return b.value
}

func (b Uint32) IsPresent() bool {
	return b.state.IsPresent()
}

func (b Uint32) ValueWithPresent() (value uint32, ok bool) {
	return b.Value(), b.state.IsPresent()
}

func (b *Uint32) Set(value uint32) {
	b.state.Set()
	b.value = value
}

func (b *Uint32) Unset() {
	b.state.Unset()
}

func (b Uint32) IsEqual(compareWith Uint32) bool {
	left, leftOk := b.ValueWithPresent()
	right, rightOk := compareWith.ValueWithPresent()
	return leftOk == rightOk && left == right
}
