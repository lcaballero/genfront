package examples
// Generated by genfront -- do not edit this file.

import (
	"errors"
	"fmt"
	"strings"
)

// Methods for the Rest state
func (r *Effort) ToMap() map[string]interface{} {
  m := make(map[string]interface{})
  m["Id"] = r.Id
  m["Title"] = r.Title
  m["Summary"] = r.Summary
  m["Description"] = r.Description
  m["CreatedBy"] = r.CreatedBy
  m["CreatedOn"] = r.CreatedOn
  m["UpdatedBy"] = r.UpdatedBy
  m["UpdatedOn"] = r.UpdatedOn
  m["OwnedBy"] = r.OwnedBy
  m["State"] = r.State
  m["RecordStatus"] = r.RecordStatus
  return m
}

// Private function (to this fil).  Don't use this function
// directly, it's intended for use only in this file.
func __snakeToPascal(sk string) string {
	parts := strings.Split(sk, "_")
	for i,p := range parts {
		parts[i] = strings.Title(p)
	}
	return strings.Join(parts, "")
}

// For known named parameter types, this strips known prefixes.
func __removePrefix(prop string) string {
	hasPrefix := strings.HasPrefix(prop, "$") ||
		strings.HasPrefix(prop, ":") ||
		strings.HasPrefix(prop, "@") ||
		strings.HasPrefix(prop, "?")
	if hasPrefix {
		return prop[1:]
	}
	return prop
}

// Returns an array of values from the Effort instance as designated
// in the props array.  The string of the props array should conform to the
// possible named value syntax which sqlite accepts.
// (See: https://www.sqlite.org/c3ref/bind_parameter_name.html and
// https://www.sqlite.org/c3ref/bind_blob.html)
func (e *Effort) Parameters(props []string) (rs []interface{}, err error) {
	rs = make([]interface{}, len(props))
	for i,p := range props {
		p = __removePrefix(p)
		switch p { 
		case "Id":
			rs[i] = e.Id
		case "Title":
			rs[i] = e.Title
		case "Summary":
			rs[i] = e.Summary
		case "Description":
			rs[i] = e.Description
		case "CreatedBy":
			rs[i] = e.CreatedBy
		case "CreatedOn":
			rs[i] = e.CreatedOn
		case "UpdatedBy":
			rs[i] = e.UpdatedBy
		case "UpdatedOn":
			rs[i] = e.UpdatedOn
		case "OwnedBy":
			rs[i] = e.OwnedBy
		case "State":
			rs[i] = e.State
		case "RecordStatus":
			rs[i] = e.RecordStatus
		default:
			err = errors.New(fmt.Sprintf("Effort doesn't have a property named: %s", p))
		}
	}
	return rs, err
}

// Fills pointer array with pointers to receiver fields.
func (e *Effort) FromColumns(cols []string, ptrs []interface{}) error {
	if len(cols) != len(ptrs) {
		return errors.New("Column length doesn't equal pointer array length")
	}
	snakeToPascal := __snakeToPascal
	for i,c := range cols {
		pascal := snakeToPascal(c)
		switch pascal { 
		case "Id":
			ptrs[i] = &e.Id
		case "Title":
			ptrs[i] = &e.Title
		case "Summary":
			ptrs[i] = &e.Summary
		case "Description":
			ptrs[i] = &e.Description
		case "CreatedBy":
			ptrs[i] = &e.CreatedBy
		case "CreatedOn":
			ptrs[i] = &e.CreatedOn
		case "UpdatedBy":
			ptrs[i] = &e.UpdatedBy
		case "UpdatedOn":
			ptrs[i] = &e.UpdatedOn
		case "OwnedBy":
			ptrs[i] = &e.OwnedBy
		case "State":
			ptrs[i] = &e.State
		case "RecordStatus":
			ptrs[i] = &e.RecordStatus
		default:
			return errors.New(fmt.Sprintf(
				"Provided a column that doesn't exist in structure: %s",
				c))
		}
	}
	return nil
}