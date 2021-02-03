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
				c.value.IfUnset(func() {
					assert.Fail(t, "optional was expected to be set, but was not")
				})
				c.value.IfSet(func(value int) {
					assert.Equal(t, c.expectValue, value)
				})
			} else {
				c.value.IfSet(func(value int) {
					assert.Fail(t, "optional was not expected to be set, but was")
				})
			}
		})
	}
}

func noOp(*Int) {}

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
