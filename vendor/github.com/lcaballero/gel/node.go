package gel

import (
	"bytes"
	"fmt"
	"io"
)

// Nodes represent the different parts of of Html as one type.  A single Node
// can be one and only one of the following types: Textual, Element, or
// Attribute.  In general a Textual node will not have children, but will
// have CData, Elements can have both children of type Text, Element,
// Attributes or Fragment, but cannot directly hold CData.  Attributes have
// only Key and Value strings, and all other fields are empty or nil.  Lastly,
// Fragments can have children of type Text and Element, while all other
// fields are empty or nil.
type Node struct {
	Tag        string
	Children   []*Node
	Attributes []*Node
	Type       Type
	Key        string
	Value      string
	CData      string
	IsVoid     bool
}

// WriteTo will output the Node to the writer correctly nesting children and
// attributes.
func (e *Node) WriteTo(w io.Writer) {
	e.WriteToIndented(Indent{}, w)
}

// ToNode is implemented to conform to a component pattern of Nodes within
// Nodes, but additionally some other instance can pose as a View if they
// implement ToNode.
func (e *Node) ToNode() *Node {
	return e
}

// WriteTo writes the Node to the given writer with the given indention.
func (e *Node) WriteToIndented(in Indent, w io.Writer) {
	switch e.Type {
	case Textual:
		if in.HasIndent() && e.CData != "" {
			in.WriteTo(w)
		}
		if e.CData != "" {
			w.Write([]byte(e.CData))
		}
		if in.HasIndent() && e.CData != "" {
			w.Write([]byte("\n"))
		}
	case Attribute:
		w.Write([]byte(" "))
		w.Write([]byte(e.Key))
		w.Write([]byte("=\""))
		w.Write([]byte(e.Value))
		w.Write([]byte("\""))
	case AttributeList:
		for _,at := range e.Children {
			at.WriteToIndented(in, w)
		}
	case NodeList:
		for _, f := range e.Children {
			f.WriteToIndented(in, w)
		}
	case Element:
		if in.HasIndent() {
			in.WriteTo(w)
		}
		w.Write([]byte("<"))
		w.Write([]byte(e.Tag))
		if len(e.Attributes) > 0 {
			for _, att := range e.Attributes {
				att.WriteTo(w)
			}
		}
		if e.IsVoid {
			w.Write([]byte("/>"))
		} else {
			w.Write([]byte(">"))
			if len(e.Children) > 0 {
				if in.HasIndent() {
					w.Write([]byte("\n"))
				}
				next := in.Incr()
				for _, kid := range e.Children {
					kid.WriteToIndented(next, w)
				}
			}
		}
		if in.Level > 0 && in.HasIndent() && len(e.Children) > 0 {
			in.WriteTo(w)
		}
		if !e.IsVoid {
			w.Write([]byte(fmt.Sprintf("</%s>", e.Tag)))
		}
		if in.Level > 0 {
			w.Write([]byte("\n"))
		}
	}
}

// String renders the Node as html (text, element, or attribute).
func (e *Node) String() string {
	buf := bytes.NewBuffer([]byte{})
	e.WriteTo(buf)
	return buf.String()
}

// Add will collect and bucket the nodes into atts and children.  Nodes
// of type Text or Element are added to children and Attribute type are
// added to the Atts slice.  If the Node is not an Element then
// attributes will silently be ignored.
func (v *Node) Add(nodes ...View) View {
	dest := v.ToNode()
	for _, view := range nodes {
		src := view.ToNode()
		switch src.Type {
		case Textual, Element, NodeList:
			dest.Children = append(dest.Children, src)
		case Attribute:
			switch dest.Type {
			case Element:
				dest.Attributes = append(dest.Attributes, src)
			case AttributeList:
				dest.Children = append(dest.Children, src)
			}
		case AttributeList:
			switch dest.Type {
			case Element:
				dest.Attributes = append(dest.Attributes, src.Children...)
			case AttributeList:
				dest.Children = append(dest.Children, src.Children...)
			}
		}
	}
	return dest
}

// Text adds the given strings as text nodes.
func (n *Node) Text(ts ...string) View {
	for _, c := range ts {
		n.Add(Text(c))
	}
	return n
}

// Att creates a new Node with Attribute type and the given key, value pair.
func Att(key, value string) View {
	node := &Node{
		Type:  Attribute,
		Key:   key,
		Value: value,
	}
	return node
}

// Atts attempts to pair up parameters and make an AttributeList.
func Atts(pairs ...string) View {
	node := &Node{
		Type:     AttributeList,
		Children: make([]*Node, 0),
	}
	for i := 0; (i+1) < len(pairs); i += 2 {
		node.Children = append(node.Children, Att(pairs[i], pairs[i+1]).ToNode())
	}
	return node
}

// Text creates a new Node with Textual type and CData from the provided string.
func Text(c string) View {
	node := &Node{
		Type:  Textual,
		CData: c,
	}
	return node
}

// Fmt creates a Text node using Sprintf.
func Fmt(format string, args ...interface{}) View {
	s := fmt.Sprintf(format, args...)
	return Text(s)
}

// Frag creates a view that is a list of children views without a containing
// element.
func Frag(children ...View) View {
	n := &Node{
		Type:     NodeList,
		Children: make([]*Node, 0),
	}
	return n.Add(children...)
}

// None produces an empty Text Node.
func None() View {
	return Text("")
}

// Maybe check if val is nil, a view or a viewable, and if so returns
// a View based on the val, but if it's none of these then it returns
// an empty Text node.
func Maybe(val interface{}) View {
	if val == nil {
		return None()
	}
	viewable, ok := val.(Viewable)
	if ok {
		return viewable.ToView()
	}
	view, ok := val.(View)
	if ok {
		return view
	}
	return None()
}

// Default takes two values, should the first be unconvertible to a Viewable
// or a view it will then attempt to convert the second.  Nil is not convertible
// so in cases where the first is nil, the second will be used if possible.
// Should they both be nil the View will be of an empty Text node.
func Default(val interface{}, def interface{}) View {
	if val == nil {
		return None()
	}
	viewable, ok := val.(Viewable)
	if ok {
		return viewable.ToView()
	}
	view, ok := val.(View)
	if ok {
		return view
	}
	return Maybe(def)
}
