package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt_IsPresent(t *testing.T) {
	cases := map[string]struct {
		value         Int
		prep          func(i *Int)
		expectPresent bool
		expectValue   int
	}{
		"not set": {
			prep: noOp,
		},
		"set value": {
			prep: func(i *Int) {
				i.Set(15)
			},
			expectPresent: true,
			expectValue:   15,
		},
		"set and unset": {
			prep: func(i *Int) {
				i.Set(15)
				i.Unset()
			},
			expectPresent: false,
		},
	}

	for caseName, c := range cases {
		c.prep(&c.value)
		assert.Equal(t, c.expectPresent, c.value.IsSet(), caseName)
		if c.expectPresent {
			assert.Equal(t, c.expectValue, c.value.Value(), caseName)
		}
	}
}

func TestInt_ValueWithOK(t *testing.T) {
	cases := map[string]struct {
		value         Int
		prep          func(i *Int)
		expectPresent bool
		expectValue   int
	}{
		"not set": {
			prep: noOp,
		},
		"set value": {
			prep: func(i *Int) {
				i.Set(15)
			},
			expectPresent: true,
			expectValue:   15,
		},
		"set and unset": {
			prep: func(i *Int) {
				i.Set(15)
				i.Unset()
			},
			expectPresent: false,
		},
	}

	for caseName, c := range cases {
		c.prep(&c.value)
		actual := c.value.Value()
		present := c.value.IsSet()
		assert.Equal(t, c.expectPresent, present, caseName)
		if c.expectPresent {
			assert.Equal(t, c.expectValue, actual, caseName)
		}
	}
}

func noOp(*Int) {}

func TestInt_IsEqual(t *testing.T) {
	cases := map[string]struct {
		source      Int
		compareWith Optionaler
		expected    bool
	}{
		"nil": {
			source:   NewIntFrom(5),
			expected: false,
		},
		"equal": {
			source:      NewIntFrom(5),
			compareWith: intAddress(NewIntFrom(5)),
			expected:    true,
		},
		"not equal": {
			source:      NewIntFrom(5),
			compareWith: intAddress(NewIntFrom(6)),
			expected:    false,
		},
		"not equal unset left": {
			source:      NewInt(),
			compareWith: intAddress(NewIntFrom(6)),
			expected:    false,
		},
		"not equal unset right": {
			source:      NewIntFrom(5),
			compareWith: unset(intAddress(NewIntFrom(6))),
			expected:    false,
		},
		"equal both unset": {
			source:      NewInt(),
			compareWith: unset(intAddress(NewIntFrom(6))),
			expected:    true,
		},
		"different type": {
			source:      NewIntFrom(5),
			compareWith: stringAddress(NewStringFrom("x")),
			expected:    false,
		},
	}

	for caseName, c := range cases {
		actual := c.source.IsEqual(c.compareWith)
		assert.Equal(t, c.expected, actual, caseName)
	}
}

func intAddress(i Int) *Int {
	return &i
}
func stringAddress(i String) *String {
	return &i
}
func unset(i Optionaler) Optionaler {
	i.Unset()
	return i
}
