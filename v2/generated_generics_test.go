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
		t.Run(caseName, func(t *testing.T) {
			c.prep(&c.value)
			assert.Equal(t, c.expectPresent, c.value.IsSet())
			if c.expectPresent {
				assert.Equal(t, c.expectValue, c.value.Value())
			}
		})
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
		t.Run(caseName, func(t *testing.T) {
			c.prep(&c.value)
			actual := c.value.Value()
			present := c.value.IsSet()
			assert.Equal(t, c.expectPresent, present)
			if c.expectPresent {
				assert.Equal(t, c.expectValue, actual)
			}
		})
	}
}

func noOp(*Int) {}

func TestInt_IsEqual(t *testing.T) {
	cases := map[string]struct {
		source      Int
		compareWith Int
		expected    bool
	}{
		"nil": {
			source:   IntFrom(5),
			expected: false,
		},
		"equal": {
			source:      IntFrom(5),
			compareWith: IntFrom(5),
			expected:    true,
		},
		"not equal": {
			source:      IntFrom(5),
			compareWith: IntFrom(6),
			expected:    false,
		},
		"not equal unset left": {
			source:      IntUnset(),
			compareWith: IntFrom(6),
			expected:    false,
		},
		"not equal unset right": {
			source:      IntFrom(5),
			compareWith: unset(IntFrom(6)),
			expected:    false,
		},
		"equal both unset": {
			source:      IntUnset(),
			compareWith: unset(IntFrom(6)),
			expected:    true,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := c.source.IsEqual(c.compareWith)
			assert.Equal(t, c.expected, actual)
		})
	}
}

func unset(i Int) Int {
	i.Unset()
	return i
}

func TestIfSet(t *testing.T) {
	cases := map[string]struct {
		input         IntTester
		expectedValue int
		expectedToSet bool
	}{
		"is set": {
			input:         IntFrom(5),
			expectedValue: 5,
			expectedToSet: true,
		},
		"not set": {
			input: IntUnset(),
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			wasSet := false
			c.input.IfSet(func(value int) {
				wasSet = true
				assert.Equal(t, c.expectedValue, value)
			})
			assert.Equal(t, c.expectedToSet, wasSet)
		})
	}
}

func TestIfUnset(t *testing.T) {
	cases := map[string]struct {
		input                Tester
		expectedNotSetCalled bool
	}{
		"is set": {
			input: IntFrom(5),
		},
		"not set": {
			input:                IntUnset(),
			expectedNotSetCalled: true,
		},
	}

	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			wasNotSetCalled := false
			c.input.IfUnset(func() {
				wasNotSetCalled = true
			})
			assert.Equal(t, c.expectedNotSetCalled, wasNotSetCalled)
		})
	}
}
