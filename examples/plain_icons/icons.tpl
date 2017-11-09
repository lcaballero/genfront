package {{ .GOPACKAGE }}
{{ .GEN_TAGLINE }}

import (
  . "github.com/lcaballero/gel"
)

{{ range .selection.icons }}

// <i class="{{ .properties.name }}"></i>
func {{ .properties.name | toSymbol }}() Tag {
  return I.Class("{{ .properties.name }}")
}{{ end }}

var Icons = []Tag{
{{ range .selection.icons }}  {{ .properties.name | toSymbol }}(),
{{ end }}
}
