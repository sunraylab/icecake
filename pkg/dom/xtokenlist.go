package dom

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* TokenList
*****************************************************************************/

// TokenList represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// # Need to call ToDOM() to update the DOM with internal value avec any change
//
// https://developer.mozilla.org/en-US/docs/Web/API/TokenList
type TokenList struct {
	jsValue js.Value
	tokens  lib.Tokens
}

// DOMTokenListFromJS is casting a js.Value into DOMTokenList.
func NewTokenListFromJS(value js.Value) (_ret *TokenList) {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	_ret = new(TokenList)
	_ret.jsValue = value
	_ret.tokens = lib.MakeTokens(_ret.jsValue.Get("value").String())
	return _ret
}

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// Length returns the number of tokens in the list.
func (_thisList *TokenList) Count() int {
	return len(_thisList.tokens)
}

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_thisList *TokenList) String() (_ret string) {
	return _thisList.tokens.String()
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_thisList *TokenList) At(index int) (_result string) {
	if index >= 0 && index < len(_thisList.tokens) {
		return _thisList.tokens[index]
	}
	return ""
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is helper.Normalized before check
func (_thisList *TokenList) Has(token string) (_result bool) {
	return _thisList.tokens.Has(token)
}

// SetValue clears and sets the list to the given value
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_thisList *TokenList) Parse(value string) *TokenList {
	if _thisList.tokens.Parse(value) {
		_thisList.jsValue.Set("value", _thisList.tokens.String())
	}
	return _thisList
}

// SetToken adds token in the list. If a token already exist it's not added to avoid duplicate.
// Always converted in lowercase.
func (_thisList *TokenList) Set(tokens ...string) *TokenList {
	if _thisList.tokens.Set(tokens...) {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns the tokenlist to enable chaining calls.
func (_thisList *TokenList) Remove(tokens ...string) *TokenList {
	if _thisList.tokens.Remove(tokens...) {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
func (_thisList *TokenList) Toggle(token string) (_updated bool) {
	_updated = _thisList.tokens.Toggle(token)
	if _updated {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _updated
}

// Replace chain a Remove and a Add
func (_thisList *TokenList) Replace(token string, newToken string) (_updated bool) {
	_updated = _thisList.tokens.Replace(token, newToken)
	if _updated {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _updated
}