package gel

//go:generate stringer -type=Type
type Type int

// The types of Node(s)
const (
	Textual       Type = 1
	Element       Type = 2
	Attribute     Type = 3
	NodeList      Type = 4
	AttributeList Type = 5
)
