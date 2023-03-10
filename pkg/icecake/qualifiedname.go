package ick

import (
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

type QualifiedName string

// Prefix returns the namespace prefix of the Name, or an empty string if no prefix is specified.
// For example, if the qualified name is xml:lang, the returned prefix is xml.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/prefix
func (_qname QualifiedName) Prefix() string {
	if strings.Contains(string(_qname), ":") {
		s := strings.Split(string(_qname), ":")
		return helper.Normalize(s[0])
	}
	return ""
}

// LocalName returns the local part of the qualified name of a Name, that is the name of the attribute,
// stripped from any namespace in front of it.
// For example, if the qualified name is xml:lang, the returned local name is lang.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (_qname QualifiedName) LocalName() string {
	name := helper.Normalize(string(_qname))
	if strings.Contains(name, ":") {
		s := strings.Split(name, ":")
		name = s[1]
	}
	return name
}
