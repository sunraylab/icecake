package markdown

import (
	"bytes"

	"github.com/sunraylab/icecake/pkg/errors"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/yuin/goldmark"
)

// RenderMarkdown process _mdtxt markdown source file and convert it to an HTML string,
// then use it as an HTML template to render it with data and components.
//
// Returns an error if the markdown processor fails.
func RenderMarkdown(_elem *ick.Element, _mdtxt string, _data any, _options ...goldmark.Option) error {
	if !_elem.IsDefined() {
		return nil
	}
	md := goldmark.New(_options...)
	var buf bytes.Buffer
	err := md.Convert([]byte(_mdtxt), &buf)
	if err != nil {
		errors.ConsoleWarnf("RenderMarkdown has error: %s", err.Error())
		return err
	}

	// HACK:
	_elem.RenderTemplate(buf.String(), _data)
	return nil
}
