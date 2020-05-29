package optional

type Bool struct {
	state base
	value bool
}

func NewBool() Bool {
	return Bool{
		state: newBase(),
	}
}

func NewBoolFrom(value bool) Bool {
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

func (b Bool) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Bool); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewInt32From(value int32) Int32 {
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

func (b Int32) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Int32); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUint16From(value uint16) Uint16 {
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

func (b Uint16) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uint16); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewStringFrom(value string) String {
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

func (b String) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*String); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUintptrFrom(value uintptr) Uintptr {
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

func (b Uintptr) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uintptr); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewInt16From(value int16) Int16 {
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

func (b Int16) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Int16); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewInt8From(value int8) Int8 {
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

func (b Int8) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Int8); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUint64From(value uint64) Uint64 {
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

func (b Uint64) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uint64); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewByteFrom(value byte) Byte {
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

func (b Byte) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Byte); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewRuneFrom(value rune) Rune {
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

func (b Rune) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Rune); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewComplex128From(value complex128) Complex128 {
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

func (b Complex128) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Complex128); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewIntFrom(value int) Int {
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

func (b Int) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Int); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewInt64From(value int64) Int64 {
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

func (b Int64) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Int64); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUintFrom(value uint) Uint {
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

func (b Uint) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uint); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUint32From(value uint32) Uint32 {
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

func (b Uint32) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uint32); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewFloat64From(value float64) Float64 {
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

func (b Float64) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Float64); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
}

type Float32 struct {
	state base
	value float32
}

func NewFloat32() Float32 {
	return Float32{
		state: newBase(),
	}
}

func NewFloat32From(value float32) Float32 {
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

func (b Float32) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Float32); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewUint8From(value uint8) Uint8 {
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

func (b Uint8) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Uint8); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
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

func NewComplex64From(value complex64) Complex64 {
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

func (b Complex64) IsEqual(compareWith Optionaler) bool {
	if compareWith == nil {
		return false
	}
	if casted, ok := compareWith.(*Complex64); !ok {
		return false
	} else {
		leftOk := b.IsSet()
		rightOk := casted.IsSet()
		if leftOk != rightOk {
			return false
		}
		// both are unset
		if !leftOk {
			return true
		}
		left := b.Value()
		right := casted.Value()
		return left == right
	}
}
