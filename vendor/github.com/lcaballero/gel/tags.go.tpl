---
normal:
  - A
  - Abbr
  - Address
  - Article
  - Aside
  - Audio
  - B
  - Bdi
  - Bdo
  - Blockquote
  - Body
  - Button
  - Canvas
  - Caption
  - Cite
  - Code
  - Colgroup
  - Data
  - Datalist
  - Dd
  - Del
  - Dfn
  - Div
  - Dl
  - Dt
  - Em
  - Fieldset
  - Figcaption
  - Figure
  - Footer
  - Form
  - H1
  - H2
  - H3
  - H4
  - H5
  - H6
  - Head
  - Header
  - Html
  - I
  - Iframe
  - Ins
  - Kbd
  - Label
  - Legend
  - Li
  - Main
  - Map
  - Mark
  - Meter
  - Nav
  - Noscript
  - Object
  - Ol
  - Optgroup
  - Option
  - Output
  - P
  - Pre
  - Progress
  - Q
  - Rb
  - Rp
  - Rt
  - Rtc
  - Ruby
  - S
  - Samp
  - Script
  - Section
  - Select
  - Small
  - Span
  - Strong
  - Style
  - Sub
  - Sup
  - Table
  - Tbody
  - Td
  - Template
  - Textarea
  - Tfoot
  - Th
  - Thead
  - Time
  - Title
  - Tr
  - U
  - Ul
  - Var
  - Video
void:
  - Area
  - Base
  - Br
  - Col
  - Embed
  - Hr
  - Img
  - Input
  - Keygen
  - Link
  - Meta
  - Param
  - Source
  - Track
  - Wbr
---
package gel
{{ .GEN_TAGLINE }}

var (
  // Normal tags requiring closing tag.
  {{ range .normal }}{{ . }} Tag = el("{{ . | lower }}", false)
  {{ end }}
  // Void elements that must be self closed.
  {{ range .void }}{{ . }} Tag = el("{{ . | lower }}", true)
  {{ end }}
)
