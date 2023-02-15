package dom

import (
	"fmt"
	"syscall/js"
)

/****************************************************************************
* enums NodeType, NodePosition
*****************************************************************************/

type NODE_TYPE int

const (
	NT_UNDEF                  NODE_TYPE = 0x0000
	NT_ALL                    NODE_TYPE = 0xFFFF
	NT_ELEMENT                NODE_TYPE = 0x0001 // An Element node like <p> or <div>. aka ELEMNT_NODE.
	NT_ATTRIBUTE              NODE_TYPE = 0x0002 // An Attribute of an Element. aka ATTRIBUTE_NODE.
	NT_TEXT                   NODE_TYPE = 0x0003 // The actual Text inside an Element or Attr. aka TEXT_NODE.
	NT_CDATA_SECTION          NODE_TYPE = 0x0004 // A CDATASection, such as <!CDATA[[ … ]]>. aka CDATA_SECTION_NODE.
	NT_PROCESSING_INSTRUCTION NODE_TYPE = 0x0007 // A ProcessingInstruction of an XML document, such as <?xml-stylesheet … ?>.
	NT_COMMENT                NODE_TYPE = 0x0008 // A Comment node, such as <!-- … -->. aka COMMENT_NODE.
	NT_DOCUMENT               NODE_TYPE = 0x0009 // A Document node. aka DOCUMENT_NODE.
	NT_DOCUMENT_TYPE          NODE_TYPE = 0x000A // A DocumentType node, such as <!DOCTYPE html>. aka DOCUMENT_TYPE_NODE.
	NT_DOCUMENT_FRAGMENT      NODE_TYPE = 0x000B // A DocumentFragment node. aka DOCUMENT_FRAGMENT_NODE.

	NT_ENTITY_REFERENCE NODE_TYPE = 0x0005 // Deprecated
	NT_ENTITY           NODE_TYPE = 0x0006 // Deprecated
	NT_NOTATION         NODE_TYPE = 0x000C // Deprecated
)

func (nt NODE_TYPE) String() string {
	switch nt {
	case NT_ELEMENT:
		return "Element"
	case NT_ATTRIBUTE:
		return "Attribute"
	case NT_TEXT:
		return "Text"
	case NT_CDATA_SECTION:
		return "Data Section"
	case NT_PROCESSING_INSTRUCTION:
		return "Processing Instruction"
	case NT_COMMENT:
		return "Comment"
	case NT_DOCUMENT:
		return "Document"
	case NT_DOCUMENT_TYPE:
		return "Document Type"
	case NT_DOCUMENT_FRAGMENT:
		return "Document Fragment"
	}
	return fmt.Sprintf("unmanaged node type: %d", nt)
}

/****************************************************************************
* enums: NodeFilter
*****************************************************************************/

// An integer value representing otherNode's position relative to node as a bitmask.
type NODE_POSITION int

const (
	NODEPOS_DISCONNECTED            NODE_POSITION = 0x01
	NODEPOS_PRECEDING               NODE_POSITION = 0x02
	NODEPOS_FOLLOWING               NODE_POSITION = 0x04
	NODEPOS_CONTAINS                NODE_POSITION = 0x08
	NODEPOS_CONTAINED_BY            NODE_POSITION = 0x10
	NODEPOS_IMPLEMENTATION_SPECIFIC NODE_POSITION = 0x20
)

/****************************************************************************
* Node
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node struct {
	EventTarget
}

// CastNode is casting a js.Value into Node.
func CastNode(value js.Value) *Node {
	if value.Type() != js.TypeObject {
		ConsoleError("casting Node failed")
		return nil
	}
	ret := new(Node)
	ret.jsValue = value
	return ret
}

func MakeNodes(value js.Value) []*Node {
	nodes := make([]*Node, 0)
	if value.Type() != js.TypeObject {
		ConsoleError("casting Nodes failed")
		return nil
	}
	len := value.Get("length").Int()
	for i := 0; i < len; i++ {
		_returned := value.Call("item", uint(i))
		node := CastNode(_returned)
		nodes = append(nodes, node)
	}
	return nodes
}

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_node *Node) IsDefined() bool {
	if _node == nil || _node.jsValue.Type() == js.TypeNull || _node.jsValue.Type() == js.TypeUndefined {
		return false
	}
	return true
}

// IsSameNode tests whether two nodes are the same (in other words, whether they reference the same object).
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isSameNode
func (_node *Node) IsSameNode(_otherNode *Node) bool {
	is := _node.jsValue.Call("isSameNode", _otherNode.jsValue)
	return is.Bool()
}

/****************************************************************************
* Node's method and properties
*****************************************************************************/

// NodeType It distinguishes different kind of nodes from each other, such as elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeType
func (_node *Node) NodeType() NODE_TYPE {
	nt := _node.jsValue.Get("nodeType")
	return NODE_TYPE(nt.Int())
}

// Values for the different types of nodes are:
//   - Attr: the qualified name of the attribute.
//   - CDATASection: "#cdata-section".
//   - Comment: "#comment".
//   - Document: "#document".
//   - DocumentFragment: "#document-fragment".
//   - DocumentType: the value of DocumentType.name
//   - Element: the uppercase name of the element tag if an HTML element, or the lowercase element tag if an XML element (like a SVG or MATHML element).
//   - ProcessingInstruction: The value of ProcessingInstruction.target
//   - Text: "#text".
//
// NodeName returns the name of the current node as a string.
func (_node *Node) NodeName() string {
	return _node.jsValue.Get("nodeName").String()
}

// BaseURI returns the absolute base URL of the document containing the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/baseURI
func (_node *Node) BaseURI() string {
	return _node.jsValue.Get("baseURI").String()
}

// IsDocConnected returns a boolean indicating whether the node is connected (directly or indirectly) to a Document object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/isConnected
func (_node *Node) IsConnected() bool {
	value := _node.jsValue.Get("isConnected")
	return (value).Bool()
}

// Doc returns the top-level document object of the node, the top-level object in which all the child nodes are created.
// on a node that is itself a document, the value is null.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/ownerDocument
func (_node *Node) Doc() *Document {
	value := _node.jsValue.Get("ownerDocument")
	return CastDocument(value)
}

// GetRootNode returns the context object's root, which optionally includes the shadow root if it is available.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/getRootNode
func (_node *Node) RootNode() *Node {
	root := _node.jsValue.Call("getRootNode")
	if typ := root.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(root)
}

// ParentNode returns the parent of the specified node in the DOM tree.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (_node *Node) ParentNode() *Node {
	parent := _node.jsValue.Get("parentNode")
	if typ := parent.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(parent)
}

// ParentElement returns the DOM node's parent Element, or null if the node either has no parent, or its parent isn't a DOM Element.
// https://developer.mozilla.org/en-US/docs/Web/API/Node/parentElement
func (_node *Node) ParentElement() *Element {
	parent := _node.jsValue.Get("parentElement")
	if typ := parent.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastElement(parent)
}

// HasChildNodes returns a boolean value indicating whether the given Node has child nodes or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/hasChildNodes
func (_node *Node) HasChildren() bool {
	has := _node.jsValue.Call("hasChildNodes")
	return has.Bool()
}

// ChildNodes returns a ~live~ static NodeList of child nodes of the given element where the first child node is assigned index 0.
// Child nodes include elements, text and comments.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/childNodes
func (_node *Node) Children() []*Node {
	nodes := _node.jsValue.Get("childNodes")
	return MakeNodes(nodes)
}

// FilteredChildren make a slice of nodes, scaning existing nodes from root to the last sibling node.
// Only nodes matching filter AND the optional match function are included.
func (_root *Node) FilteredChildren(_filter NODE_TYPE, _deepmax int, match func(*Node) bool) []*Node {
	nodes := make([]*Node, 0)
	c := 0
	for scan := _root; scan != nil && c < 20; scan = _root.SiblingNext() {
		c++
		fmt.Println("scanning :", scan.NodeName(), scan.NodeType().String())

		// check filtered node type
		fn := (_filter & scan.NodeType()) > 0

		// apply the filter to children if not too deep
		if fn && scan.HasChildren() && _deepmax > 0 {
			for _, child := range scan.Children() {
				filteredchildren := child.FilteredChildren(_filter, _deepmax-1, match)
				nodes = append(nodes, filteredchildren...)
			}
		}

		// apply matching function
		if fn && match != nil {
			fn = fn && match(scan)
		}

		if fn {
			nodes = append(nodes, scan)
		}
	}
	return nodes
}

// FirstChild returns the node's first child in the tree, or null if the node has no children.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/firstChild
func (_node *Node) ChildFirst() *Node {
	child := _node.jsValue.Get("firstChild")
	if typ := child.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(child)
}

// LastChild returns the last child of the node.
// If its parent is an element, then the child is generally an element node, a text node, or a comment node.
// It returns null if there are no child nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/lastChild
func (_node *Node) ChildLast() *Node {
	child := _node.jsValue.Get("lastChild")
	if typ := child.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(child)
}

// PreviousSibling  returns the node immediately preceding the specified one in its parent's childNodes list, or null if the specified node is the first in that list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/previousSibling
func (_node *Node) SiblingPrevious() *Node {
	sibling := _node.jsValue.Get("previousSibling")
	if typ := sibling.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(sibling)
}

// NextSibling returns the node immediately following the specified one in their parent's childNodes, or returns null if the specified node is the last child in the parent element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nextSibling
func (_node *Node) SiblingNext() *Node {
	sibling := _node.jsValue.Get("nextSibling")
	if typ := sibling.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastNode(sibling)
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) NodeValue() string {
	value := _node.jsValue.Get("nodeValue")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return value.String()
}

// NodeValue is a string containing the value of the current node, if any.
//
// For the document itself, nodeValue returns null.
// For text, comment, and CDATA nodes, nodeValue returns the content of the node.
// For attribute nodes, the value of the attribute is returned.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/nodeValue
func (_node *Node) SetNodeValue(value string) (_ret *Node) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_node.jsValue.Set("nodeValue", input)
	return _node
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) TextContent() string {
	value := _node.jsValue.Get("textContent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return value.String()
}

// TextContent represents the text content of the node and its descendants.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (_node *Node) SetTextContent(value string) (_ret *Node) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = nil
	}
	_node.jsValue.Set("textContent", input)
	return _node
}

// CompareDocumentPosition reports the position of its argument node relative to the node on which it is called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/compareDocumentPosition
func (_onenode *Node) ComparePosition(_other *Node) NODE_POSITION {
	_returned := _onenode.jsValue.Call("compareDocumentPosition", _other.jsValue)
	return NODE_POSITION(_returned.Int())
}

// InsertBefore inserts a newnode before a refnode.
//
// if refnode is nil, then newNode is inserted at the end of node's child nodes.
//
// If the given node already exists in the document, insertBefore() moves it from its current position to the new position.
// (That is, it will automatically be removed from its existing parent before appending it to the specified new parent.)
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/insertBefore
func (_parentnode *Node) InsertBefore(newnode *Node, refnode *Node) *Node {
	node := _parentnode.jsValue.Call("insertBefore", newnode.jsValue, refnode.jsValue)
	return CastNode(node)
}

// AppenChild adds a node to the end of the list of children of a specified parent node.
// If the given child is a reference to an existing node in the document, appendChild() moves it from its current position to the new position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/appendChild
func (_parentnode *Node) AppendChild(newnode *Node) *Node {
	node := _parentnode.jsValue.Call("appendChild", newnode.jsValue)
	return CastNode(node)
}

// ReplaceChild replaces a child node within the given (parent) node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/replaceChild
func (_parentnode *Node) ReplaceChild(_newchild *Node, _oldchild *Node) *Node {
	node := _parentnode.jsValue.Call("replaceChild", _newchild.jsValue, _oldchild.jsValue)
	return CastNode(node)
}

// RemoveChild removes a child node from the DOM and returns the removed node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild
func (_parentnode *Node) RemoveChild(_newchild *Node) *Node {
	node := _parentnode.jsValue.Call("removeChild", _newchild.jsValue)
	return CastNode(node)
}
