package ick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttributes(t *testing.T) {
	// cache
	var as1, as2 *Attributes
	var err error

	as1, err = ParseAttributes("single")
	assert.NoError(t, err)
	assert.Equal(t, "single", as1.String())

	as1, err = ParseAttributes("one two")
	assert.NoError(t, err)
	assert.Equal(t, "one two", as1.String())

	as2, err = ParseAttributes("zero=0 bool=False one=1 two three=3 four five six")
	assert.NoError(t, err)
	assert.Equal(t, `bool=false five four one=1 six three=3 two zero=0`, as2.String())

	as1, err = ParseAttributes("one='one' two='two'")
	assert.NoError(t, err)
	assert.Equal(t, `one='one' two='two'`, as1.String())

	as1, err = ParseAttributes("  this    =   'with \"quoted sub value\"' anotherone ")
	assert.NoError(t, err)
	assert.Equal(t, `anotherone this='with "quoted sub value"'`, as1.String())

	as1, err = ParseAttributes(`a1="o'connor"`)
	assert.NoError(t, err)
	assert.Equal(t, `a1="o'connor"`, as1.String())

	v, _ := as1.Attribute("a1")
	assert.Equal(t, `o'connor`, v)

	as1.SetAttributes(*as2, false)
	assert.Equal(t, 9, len(as1.Keys()))

	assert.True(t, as1.IsTrue("one"))
	assert.False(t, as1.IsTrue("ten"))
	assert.False(t, as1.IsTrue("zero"))
	assert.False(t, as1.IsTrue("bool"))
	assert.False(t, as1.Hidden())

	assert.Equal(t, 0, as1.TabIndex())
	assert.Equal(t, 2, as1.SetTabIndex(2).TabIndex())

	_, f := as1.RemoveAttribute("bool").Attribute("bool")
	assert.False(t, f)

	assert.True(t, as1.Toggle("bool"))
	assert.False(t, as1.Toggle("bool"))

	_, err = ParseAttributes("one t#o three")
	assert.Error(t, err)

	as3, _ := ParseAttributes("data-a data-s='ok' data-v=10")
	as2.SetAttributes(*as3, false)
	assert.Equal(t, "data-a data-s='ok' data-v=10", as2.Data().String())

}