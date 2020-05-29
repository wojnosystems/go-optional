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
