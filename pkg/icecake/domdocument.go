package ick

import (
	"net/url"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/errors"
)

/******************************************************************************
* Document
******************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
type DOC_READYSTATE string

const (
	DOC_READY       DOC_READYSTATE = "loading"
	DOC_INTERACTIVE DOC_READYSTATE = "interactive"
	DOC_COMPLETE    DOC_READYSTATE = "complete"
)

// The Document.visibilityState returns the visibility of the document,
// that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
type DOC_VISIBILITYSTATE string

const (
	DOC_HIDDEN    DOC_VISIBILITYSTATE = "hidden"
	DOC_VISIBLE   DOC_VISIBILITYSTATE = "visible"
	DOC_PRERENDER DOC_VISIBILITYSTATE = "prerender"
)

// Document represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree.
//
// The Document describes the common properties and methods for any kind of document.
// Depending on the document's type (e.g. HTML, XML, SVG, …), a larger API is available:
// HTML documents, served with the "text/html" content type,
// also implement the HTMLDocument interface, whereas XML and SVG documents implement the XMLDocument interface.
type Document struct {
	Node
}

// CastDocument is casting a js.Value into Document.
func CastDocument(_jsv JSValueProvider) *Document {
	if _jsv.Value().Type() != TYPE_OBJECT {
		errors.ConsoleErrorf("casting Document failed")
		return &Document{}
	}
	doc := new(Document)
	doc.jsvalue = _jsv.Value().jsvalue
	return doc
}

// GetDocument returns the current document within the current window
func GetDocument() *Document {
	jsdoc := val(js.Global().Get("document"))
	if !jsdoc.IsObject() {
		errors.ConsoleStackf(nil, "Unable to get document")
	}
	doc := new(Document)
	doc.jsvalue = jsdoc.jsvalue
	return doc
}

/******************************************************************************
* Document's properties
******************************************************************************/

// CompatMode ndicates whether the document is rendered in Quirks mode or Standards mode.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/compatMode
func (_doc *Document) CompatMode() string {
	return _doc.Get("compatMode").String()
}

// CharacterSet returns the character encoding of the document that it's currently rendered with.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (_doc *Document) CharacterSet() string {
	return _doc.Get("characterSet").String()
}

// Doctype Returns the Document Type Declaration (DTD) associated with current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (_doc *Document) Doctype() *DocumentType {
	value := _doc.Get("doctype")
	return CastDocumentType(value)
}

// ContentType returns the MIME type that the document is being rendered as.
// This may come from HTTP headers or other sources of MIME information,
// and might be affected by automatic type conversions performed by either the browser or extensions.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/contentType
func (_doc *Document) ContentType() string {
	return _doc.Get("contentType").String()
}

// Referrer returns the URI of the page that linked to this page.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/referrer
func (_doc *Document) Referrer() *url.URL {
	ref := _doc.Get("referrer").String()
	u, _ := url.Parse(ref)
	return u
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc *Document) Cookie() string {
	return _doc.Get("cookie").String()
}

// Cookie lets you read and write cookies associated with the document.
// It serves as a getter and setter for the actual values of the cookies.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie
func (_doc *Document) SetCookie(value string) *Document {
	_doc.Set("cookie", value)
	return _doc
}

// LastModified returns a string containing the date and time on which the current document was last modified.
//
// TODO: handle time.Time
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/lastModified
func (_doc *Document) LastModified() string {
	return _doc.Get("lastModified").String()
}

// ReadyState describes the loading state of the document. When the value of this property changes, a readystatechange event fires on the document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/readyState
func (_doc *Document) ReadyState() DOC_READYSTATE {
	value := _doc.Get("readyState").String()
	return DOC_READYSTATE(value)
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc *Document) Title() string {
	return _doc.Get("title").String()
}

// Title gets or sets the current title of the document. When present, it defaults to the value of the <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/title
func (_doc *Document) SetTitle(value string) *Document {
	_doc.Set("title", value)
	return _doc
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc *Document) Body() *Element {
	value := _doc.Get("body")
	return CastElement(value)
}

// The Document.body property represents the <body> or <frameset> node of the current document, or null if no such element exists.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (_doc *Document) SetBody(value *Element) *Document {
	_doc.Set("body", value)
	return _doc
}

// Head  returns the <head> element of the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (_doc *Document) Head() *HeadElement {
	elem := _doc.Get("head")
	return CastHeadElement(elem)
}

// DocumentElement returns the Element that is the root element of the document (for example, the <html> element for HTML documents).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (_doc *Document) RootElement() *Element {
	elem := _doc.Get("documentElement")
	return CastElement(elem)
}

// ActiveElement returns the Element within the DOM that currently has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/activeElement
func (_doc *Document) FocusedElement() *Element {
	elem := _doc.Get("activeElement")
	return CastElement(elem)
}

// FullscreenElement returns the Element that is currently being presented in fullscreen mode in this document, or null if fullscreen mode is not currently in use.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/fullscreenElement
func (_doc *Document) FullscreenElement() *Element {
	elem := _doc.Get("fullscreenElement")
	return CastElement(elem)
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc *Document) DesignMode() string {
	return _doc.Get("designMode").String()
}

// DesignMode controls whether the entire document is editable.
// Valid values are "on" and "off". According to the specification, this property is meant to default to "off".
// Firefox follows this standard. The earlier versions of Chrome and IE default to "inherit".
// Starting in Chrome 43, the default is "off" and "inherit" is no longer supported. In IE6-10, the value is capitalized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/designMode
func (_doc *Document) SetDesignMode(value string) *Document {
	_doc.Set("designMode", value)
	return _doc
}

// Hidden returns a Boolean value indicating if the page is considered hidden or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (_doc *Document) Hidden() bool {
	return _doc.Get("hidden").Bool()
}

// VisibilityState returns the visibility of the document, that is in which context this element is now visible.
// It is useful to know if the document is in the background or an invisible tab, or only loaded for pre-rendering.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/visibilityState
func (_doc *Document) VisibilityState() DOC_VISIBILITYSTATE {
	value := _doc.Get("visibilityState").String()
	return DOC_VISIBILITYSTATE(value)
}

// HasFocus returns a boolean value indicating whether the document or any element inside the document has focus.
// This method can be used to determine whether the active element in a document has focus.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hasFocus
func (_doc *Document) HasFocus() bool {
	return _doc.Call("hasFocus").Bool()
}

// ChildElementCount returns the number of child elements of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/childElementCount
func (_doc *Document) ChildrenCount() int {
	return _doc.GetInt("childElementCount")
}

// GetElementsByTagName returns an HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (_doc *Document) ChildrenByTagName(qualifiedName string) []*Element {
	elems := _doc.Call("getElementsByTagName", qualifiedName)
	if !elems.IsDefined() {
		errors.ConsoleWarnf("ChildrenByTagName failed: %q not found\n", qualifiedName)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementsByClassName returns an array-like object of all child elements which have all of the given class name(s).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (_doc *Document) ChildrenByClassName(classNames string) []*Element {
	elems := _doc.Call("getElementsByClassName", classNames)
	if !elems.IsDefined() {
		errors.ConsoleWarnf("ChildrenByClassName failed: %q not found\n", classNames)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementsByName returns a NodeList Collection of elements with a given name attribute in the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByName
func (_doc *Document) ChildrenByName(elementName string) []*Element {
	elems := _doc.Call("getElementsByName", elementName)
	if !elems.IsDefined() {
		errors.ConsoleWarnf("ChildrenByName failed: %q not found\n", elementName)
		return make([]*Element, 0)
	}
	return CastElements(elems)
}

// GetElementById returns an Element object representing the element whose id property matches the specified string.
// Since element IDs are required to be unique if specified, they're a useful way to get access to a specific element quickly.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (_doc *Document) ChildById(_elementId string) (_result *Element) {
	_elementId = helper.Normalize(_elementId)
	elem := _doc.Call("getElementById", _elementId)
	if elem.Truthy() && CastNode(elem).NodeType() == NT_ELEMENT {
		return CastElement(elem)
	}
	errors.ConsoleWarnf("ChildById failed: %q not found, or not an <Element>\n", _elementId)
	return new(Element)
}

// QuerySelector returns the first Element within the document that matches the specified selector, or group of selectors.
// If no matches are found, null is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (_doc *Document) SelectorQueryFirst(selectors string) *Element {
	elem := _doc.Call("querySelector", selectors)
	if elem.Truthy() && CastNode(elem).NodeType() == NT_ELEMENT {
		return CastElement(elem)
	}
	errors.ConsoleWarnf("querySelector failed: %q not found, or not a <Element>\n", selectors)
	return new(Element)
}

// querySelectorAll returns a static (not live) NodeList representing a list of the document's elements that match the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (_doc *Document) SelectorQueryAll(selectors string) []*Element {
	elems := _doc.Call("querySelectorAll", selectors)
	if !elems.IsDefined() {
		errors.ConsoleWarnf("SelectorQueryAll failed: %q not found\n", selectors)
		return nil
	}
	return CastElements(elems)
}

// CreateElement creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (_doc *Document) CreateElement(tagName string) *Element {
	elem := _doc.Call("createElement", tagName)
	return CastElement(elem)
}

// CreateAttribute  creates a new attribute node, and returns it. The object created is a node implementing the Attr interface.
// The DOM does not enforce what sort of attributes can be added to a particular element in this manner.
//
// # TODO test CreateAttribute
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createAttribute
// func (_doc *Document) CreateAttribute(localName string) *Attribute {
// 	attr := _doc.jsValue.Call("createAttribute", localName)
// 	return CastAttribute(attr)
// }

// GetElementAtPoint returns the topmost Element at the specified coordinates (relative to the viewport).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementFromPoint
func (_doc *Document) ChildAtPoint(x float64, y float64) *Element {
	elem := _doc.Call("elementFromPoint", x, y)
	return CastElement(elem)
}

// GetElementsAtPoint eturns an array of all elements at the specified coordinates (relative to the viewport).
// The elements are ordered from the topmost to the bottommost box of the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/elementsFromPoint
func (_doc *Document) ChildrenAtPoint(x float64, y float64) (_result []*Element) {
	elems := _doc.Call("elementsFromPoint", x, y)
	len := elems.jsvalue.Length()
	_result = make([]*Element, len)
	for i := 0; i < len; i++ {
		elem := elems.Index(i)
		_result[i] = CastElement(elem)
	}
	return _result
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_doc *Document) PrependNodes(nodes []*Node) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		if n != nil {
			jsn := n.jsvalue
			_args[_end] = jsn
			_end++
		}
	}
	_doc.Call("prepend", _args[0:_end]...)
}

// Prepend inserts a set of Node objects or string objects before the first child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/prepend
func (_doc *Document) PrepenStrings(strs []string) {
	var _args []interface{} = make([]interface{}, len(strs))
	var _end int
	for _, n := range strs {
		_args[_end] = n
		_end++
	}
	_doc.Call("prepend", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_doc *Document) AppendNodes(nodes []*Node) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		if n != nil {
			jsn := n.jsvalue
			_args[_end] = jsn
			_end++
		}
	}
	_doc.Call("append", _args[0:_end]...)
}

// Append inserts a set of Node objects or string objects after the last child of the document.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (_doc *Document) AppendStrings(strs []string) {
	var _args []interface{} = make([]interface{}, len(strs))
	var _end int
	for _, n := range strs {
		_args[_end] = n
		_end++
	}
	_doc.Call("append", _args[0:_end]...)
}

/******************************************************************************
* Document's  GENRIC_EVENT
******************************************************************************/

func makeDoc_Generic_Event(listener func(event *Event, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastEvent(value)
		target := CastDocument(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				errors.ConsoleStackf(r, "Error occurs processing event %q on Document", evt.Type())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *Document)) func() {
	evh := makeDoc_Generic_Event(listener)
	_doc.Call("addEventListener", string(evttype), evh)
	close := func() {
		_doc.Call("removeEventListener", string(evttype), evh)
		evh.Release()
	}
	return close
}

/******************************************************************************
* Document's  MOUSE_EVENT
******************************************************************************/

func makeDoc_Mouse_Event(listener func(event *MouseEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastMouseEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *Document)) func() {
	evh := makeDoc_Mouse_Event(listener)
	_doc.Call("addEventListener", string(evttype), evh)
	close := func() {
		_doc.Call("removeEventListener", string(evttype), evh)
		evh.Release()
	}
	return close
}

/******************************************************************************
* Document's  FOCUS_EVENT
******************************************************************************/

func makeDoc_Focus_Event(listener func(event *FocusEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastFocusEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *Document)) js.Func {
	callback := makeDoc_Focus_Event(listener)
	_doc.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Document's  POINTER_EVENT
******************************************************************************/

func makeDoc_Pointer_Event(listener func(event *PointerEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastPointerEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *Document)) js.Func {
	callback := makeDoc_Pointer_Event(listener)
	_doc.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Document's  INPUT_EVENT
******************************************************************************/

func makeDoc_Input_Event(listener func(event *InputEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastInputEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *Document)) js.Func {
	callback := makeDoc_Input_Event(listener)
	_doc.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Document's  KEYBOARD_EVENT
******************************************************************************/

func makeDoc_Keyboard_Event(listener func(event *KeyboardEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastKeyboardEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_doc *Document) AddKeyboardEvent(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *Document)) js.Func {
	callback := makeDoc_Keyboard_Event(listener)
	_doc.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

// event attribute: UIEvent
func makeDoc_UIEvent(listener func(event *UIEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastUIEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventResize(listener func(event *UIEvent, target *Document)) js.Func {
	callback := makeDoc_UIEvent(listener)
	_doc.Call("addEventListener", "resize", callback)
	return callback
}

/******************************************************************************
* Document's  WHEEL_EVENT
******************************************************************************/

func makeDoc_Wheel_Event(listener func(event *WheelEvent, target *Document)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := val(args[0])
		evt := CastWheelEvent(value)
		target := CastDocument(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddEventWheel(listener func(event *WheelEvent, currentTarget *Document)) js.Func {
	callback := makeDoc_Wheel_Event(listener)
	_doc.Call("addEventListener", "wheel", callback)
	return callback
}

/******************************************************************************
* Document's  FULLSCREEN_EVENT
******************************************************************************/

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_doc *Document) AddFullscreenEvent(evttype FULLSCREEN_EVENT, listener func(event *Event, target *Document)) js.Func {
	cb := makeDoc_Generic_Event(listener)
	_doc.Call("addEventListener", string(evttype), cb)
	return cb
}
