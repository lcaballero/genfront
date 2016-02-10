package doctable

import (
	"strings"
)

// FieldAndDoc structure is intended to name a structure and then list out
// field names and doc comments for those fields.
type FieldAndDoc struct {
	Name     string
	FieldDoc map[string]string
}

// NewFieldAndDoc creates a new FieldAndDoc instance with the given name.
func NewFieldAndDoc(name string) *FieldAndDoc {
	return &FieldAndDoc{
		Name:     name,
		FieldDoc: make(map[string]string),
	}
}

// Add fills in the FieldDoc member with key value pairs where the key is
// the name of the field and the value is the associated documentation string.
func (f *FieldAndDoc) Add(field, doc string) {
	f.FieldDoc[field] = f.commentToString(doc)
}

// commentToString removes comment delimeters from the given string, namely
// the tokens '//', '/*' and '*/', then trims whitespace around from
// the comment.
func (f *FieldAndDoc) commentToString(cmt string) string {
	cutset := " \n\r"
	if strings.HasPrefix(cmt, "//") {
		cmt = cmt[2:]
		return strings.Trim(cmt, cutset)
	}
	if strings.HasPrefix(cmt, "/*") {
		cmt = cmt[2:]
	}
	if strings.HasSuffix(cmt, "*/") {
		cmt = cmt[:len(cmt)-2]
	}
	cmt = strings.Trim(cmt, " \n\r")
	return cmt
}
