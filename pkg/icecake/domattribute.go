package ick

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
)

/****************************************************************************
* Attribute
*****************************************************************************/

// Attr represents one of an element's attributes as an object.
// In most situations, you will directly retrieve the attribute value as a string (e.g., Element.getAttribute()),
// but certain functions (e.g., Element.getAttributeNode()) or means of iterating return Attr instances.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
type Attribute struct {
	// the Element the attribute belongs to.
	ownerElement *Element

	// attr.Name returns the qualified name of an attribute, that is the name of the attribute, with the namespace prefix, if any, in front of it.
	// For example, if the local name is lang and the namespace prefix is xml, the returned qualified name is xml:lang.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/value
	//
	// attr.Value contains the value of the attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/name
	name  string
	value string
}

/****************************************************************************
* Attribute's factory
*****************************************************************************/

// CastAttribute is casting a js.Value into Attribute
func CastAttribute(value js.Value) *Attribute {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting Attribute failed")
		return nil
	}
	cast := new(Attribute)
	cast.name = value.Get("name").String()
	cast.value = value.Get("value").String()
	cast.ownerElement = CastElement(value.Get("ownerElement"))
	return cast
}

// NewAttributeToDOM create a new attribut and update the DOM with setattribute
// func NewAttribute(_attr lib.Attribute, ownerElement *Element) *Attribute {
// 	attr := new(Attribute)
// 	attr.name = helper.Normalize(_attr.Name)
// 	attr.value = strings.Trim(_attr.Value, " ")
// 	attr.ownerElement = ownerElement
// 	attr.ownerElement.JSValue().Call("setAttribute", attr.name, attr.value)
// 	return attr
// }

/****************************************************************************
* Attribute's Preperties & Methods
*****************************************************************************/

// String returns normalized formated properties of this attribute
//
//	if value is empty, the format is `{name}`
//	else the format is `{name}="{value}"`
func (_attr *Attribute) String() (_str string) {
	if _attr == nil {
		return ""
	}
	_str = helper.Normalize(_attr.name)
	if _attr.value != "" {
		sep := "'"
		if strings.ContainsRune(_attr.value, rune('\'')) {
			sep = "\""
		}
		_str += `=` + sep + _attr.value + sep
	}
	return _str
}

// OwnerElement returns the Element the attribute belongs to.
//
// returns an empty element if _attribute is nil
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/ownerElement
func (_attribute *Attribute) OwnerElement() *Element {
	if _attribute == nil {
		log.Println("OwnerElement() call on a nil Attribute")
		return &Element{}
	}
	return _attribute.ownerElement
}

func (_attribute *Attribute) Name() string {
	if _attribute == nil {
		return ""
	}
	return _attribute.name
}

func (_attribute *Attribute) Value() string {
	if _attribute == nil {
		return ""
	}
	return _attribute.value
}

func (_attribute *Attribute) IsTrue() bool {
	if _attribute == nil {
		return false
	}
	if _attribute.value == "false" || _attribute.value == "0" {
		return false
	}
	return true
}

// Update updates the DOM of the ownerElement with this attribute's value.
// The value should be nomamized before call if required.
func (_attribute *Attribute) Update(_value string) {
	if _attribute == nil || !_attribute.ownerElement.IsDefined() {
		log.Println("Update failed")
		return
	}
	_attribute.value = _value
	_attribute.ownerElement.JSValue().Call("setAttribute", string(_attribute.name), _attribute.value)
}