package gel

type Tag func(...View) View
func (t Tag) ToNode() *Node {
	return t().ToNode()
}

// el creates a Tag func for the given tag name and isVoid flag.
func el(tag string, isVoid bool) Tag {
	return func(children ...View) View {
		node := &Node{
			Type:       Element,
			Tag:        tag,
			Children:   make([]*Node, 0),
			Attributes: make([]*Node, 0),
			IsVoid:     isVoid,
		}
		node.Add(children...)
		return node
	}
}

// Adds the class attribute with the given value
func (t Tag) Class(class string) Tag {
	return t.Atts("class", class)
}

// Atts creates a tag with the given pairs of Attributes.
func (t Tag) Atts(pairs ...string) Tag {
	return func(views ...View) View {
		atts := []View { Atts(pairs...) }
		atts = append(atts, views...)
		return t(atts...)
	}
}

// Text will create an Element Node from the Tag and then immediately add the
// given strings as Text nodes.
func (t Tag) Text(c ...string) View {
	e := t().ToNode()
	for _, txt := range c {
		e.Add(Text(txt))
	}
	return e
}

// Fmt creates a Text tag using the given format and args like Sprintf.
func (t Tag) Fmt(format string, args ...interface{}) View {
	return t(Fmt(format, args...))
}
