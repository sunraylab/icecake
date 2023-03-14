package wick

/*****************************************************************************
* WebApp
******************************************************************************/

var App *WebApp

func init() {
	App = NewWebApp()
}

// WebApp
type WebApp struct {
	Document // The embedded DOM document

	browser Window // The Global JS Window object
}

// NewWebApp is the WebApp factory. Must be call once at the begining of the wasm main code.
func NewWebApp() *WebApp {
	webapp := new(WebApp)
	webapp.browser.Wrap(GetWindow())
	webapp.Document.Wrap(GetDocument())

	return webapp
}

// Browser returns the DOM.Window object
func (_app *WebApp) Browser() *Window {
	return &_app.browser
}
