package browser

import "syscall/js"

/******************************************************************************
* EventHandlerFunc
******************************************************************************/

// callback: EventHandlerNonNull
type EventHandlerFunc func(event *Event) interface{}

func MakeEventHandlerFuncFromJS(_value js.Value) EventHandlerFunc {
	return func(event *Event) (_result interface{}) {
		var _args [1]interface{}
		_args[0] = event.jsValue
		return _value.Invoke(_args[0:1]...)
	}
}

/******************************************************************************
* EventHandler
******************************************************************************/

// EventHandler is a javascript function type.
//
// Call Release() when done to release resouces allocated to this type.
type EventHandler js.Func

func NewEventHandler(callback EventHandlerFunc) *EventHandler {
	if callback == nil {
		return nil
	}
	ret := EventHandler(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_p0 := NewEventFromJS(args[0])
		_returned := callback(_p0)
		_converted := _returned
		return _converted
	}))
	return &ret
}

/******************************************************************************
* EventListener
******************************************************************************/

// GoEventListener is a callback interface.
type GoEventListener interface {
	HandleEvent(event *Event)
}

// EventListener is javascript reference value for callback interface EventListener.
// This is holding the underlying javascript object.
type EventListener struct {
	// Value is the underlying javascript object or function.
	jsValue js.Value

	// Function is the underlying function objects that is allocated for the interface callback
	Function js.Func

	// Go interface to invoke
	goimpl      GoEventListener
	gofunction  func(event *Event)
	gouseInvoke bool
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *EventListener) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

/******************************************************************************
* EventListener's factory
******************************************************************************/

// NewEventListenerFromJS is taking an javascript object that reference to a
// callback interface and return a corresponding interface that can be used
// to invoke on that element.
func NewEventListenerFromJS(value js.Value) *EventListener {
	if value.Type() == js.TypeObject {
		return &EventListener{jsValue: value}
	}
	if value.Type() == js.TypeFunction {
		return &EventListener{jsValue: value, gouseInvoke: true}
	}
	panic("unsupported type")
}

// NewEventListener is allocating a new javascript object that implements EventListener interface.
func NewEventListener(callback GoEventListener) *EventListener {
	ret := &EventListener{goimpl: callback}
	ret.jsValue = js.Global().Get("Object").New()
	ret.Function = ret.allocateHandleEvent()
	ret.jsValue.Set("handleEvent", ret.Function)
	return ret
}

// NewEventListenerFunc is allocating a new javascript function is implements EventListener interface.
func NewEventListenerFunc(f func(event *Event)) *EventListener {
	// single function will result in javascript function type, not an object
	ret := &EventListener{gofunction: f}
	ret.Function = ret.allocateHandleEvent()
	ret.jsValue = ret.Function.Value
	return ret
}

func (t *EventListener) allocateHandleEvent() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_p0 := NewEventFromJS(args[0])
		if t.gofunction != nil {
			t.gofunction(_p0)
		} else {
			t.goimpl.HandleEvent(_p0)
		}
		return nil
	})
}

/******************************************************************************
* EventListener's methods
******************************************************************************/

// Release is releasing all resources that is allocated.
func (_this *EventListener) Release() {
	if _this.Function.Type() != js.TypeUndefined {
		_this.Function.Release()
	}
}

func (_this *EventListener) HandleEvent(event *Event) {
	if _this.gofunction != nil {
		_this.gofunction(event)
	}
	if _this.goimpl != nil {
		_this.goimpl.HandleEvent(event)
	}
	var _args [1]interface{}
	_args[0] = event.jsValue
	if _this.gouseInvoke {
		// invoke a javascript function
		_this.jsValue.Invoke(_args[0:1]...)
	} else {
		_this.jsValue.Call("handleEvent", _args[0:1]...)
	}
}

/******************************************************************************
* EventTarget
*******************************************************************************/

// EventTarget
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *EventTarget) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// NewEventTargetFromJS is casting a js.Value into EventTarget.
func NewEventTargetFromJS(value js.Value) *EventTarget {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		// TODO: error handling
		return nil
	}
	ret := &EventTarget{}
	ret.jsValue = value
	return ret
}

// NewEventTarget create a new EventTarget
func NewEventTarget() (_result *EventTarget) {
	_klass := js.Global().Get("EventTarget")
	var _args [0]interface{}
	_returned := _klass.New(_args[0:0]...)
	return NewEventTargetFromJS(_returned)
}

/******************************************************************************
* EventTarget's events
*******************************************************************************/

// AddEventListener sets up a function that will be called whenever the specified event is delivered to the target.
//
// Common targets are Element, or its children, Document, and Window, but the target may be any object that supports events (such as XMLHttpRequest).
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (_this *EventTarget) AddEventListener(_type string, callback *EventListener) {
	_this.jsValue.Call("addEventListener", _type, callback.JSValue())
}

// RemoveEventListener removes an event listener previously registered with EventTarget.addEventListener() from the target.
// The event listener to be removed is identified using a combination of the event type, the event listener function itself,
// and various optional options that may affect the matching process;
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (_this *EventTarget) RemoveEventListener(_type string, callback *EventListener) {
	_this.jsValue.Call("removeEventListener", _type, callback.JSValue())
}
