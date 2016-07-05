package gel

import (
	"bytes"
	"errors"
	"io"
)

const DefaultLevel = 0
const DefaultIncrement = 1
const DefaultTab = "  "

// Indent represents indention at a given level.
type Indent struct {
	Level int
	Inc   int
	Tab   string
}

// Returns an Indent value starting at level 0, with an increment
// of 0, and a tab of 2 spaces.
func NewIndent() Indent {
	return Indent{
		Level: DefaultLevel,
		Inc:   DefaultIncrement,
		Tab:   DefaultTab,
	}
}

// String produces the indention for the given level of the Indent.
func (n Indent) String() string {
	buf := bytes.NewBuffer([]byte{})
	n.WriteTo(buf)
	return buf.String()
}

// HasIndent returns true if the Inc is > 0 and Tab != ''.
func (n Indent) HasIndent() bool {
	noIndent := n.Inc == 0 && n.Tab == ""
	return !noIndent
}

// WriteTo outputs the Indent to the Writer.
func (n Indent) WriteTo(w io.Writer) {
	for i := 0; i < n.Level; i++ {
		w.Write([]byte(n.Tab))
	}
}

// Incr adds one level to the Indent.
func (n Indent) Incr() Indent {
	return Indent{Level: n.Level + n.Inc, Tab: n.Tab, Inc: n.Inc}
}

// Decr reduces Indent by one level.
func (n Indent) Decr() Indent {
	if (n.Level - n.Inc) < 0 {
		panic(errors.New("Cannot decrement Indent below 0"))
	}
	return Indent{Level: n.Level - n.Inc, Tab: n.Tab}
}
