[![Build Status](https://travis-ci.org/lcaballero/genfront.svg?branch=master)](https://travis-ci.org/lcaballero/genfront)


# Overview

`genfront` is a code generating tool.  `genfront` provides several
subcommands for different generating patterns.  See the usage below.


## Subcommands

#### front
Parses a front-matter file and renders the template therein providing
the static embedded yaml data.

#### fields
Placed above a struct, it provides struct fields as data to the
template for rendering.


## Template Helpers

#### title
#### lower
#### toSymbol
#### getenv
#### split
#### camelCase
#### hasPrefix


## Example Usage

```
//go:generate genfront front --input req_methods.fm --output req_methods.go
```

*req_methods.fm*
```go
---
methods:
  - OPTIONS
  - GET
  - HEAD
  - POST
  - PUT
  - DELETE
  - TRACE
  - CONNECT
---
package {{ .ENV.GOPACKAGE }}
{{ .ENV.GEN_TAGLINE }}
// {{ getenv "GOLINE" }}

const (
{{ range .methods }}	{{ . }} = "{{ . }}"
{{ end }})

// Methods for the Rest state{{ range .methods }}
func (r *Rest) {{ . | title }}() *Rest {
	return r.Method({{ . }})
}{{ end }}

// Methods for Req state{{ range .methods }}
func (r *Req) {{ . | title }}() *Req {
	return r.Method({{ . }})
}{{ end }}
```


## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
[Front Matter]: https://jekyllrb.com/docs/frontmatter/
[Yaml]: http://yaml.org/
[Handlebars]: http://handlebarsjs.com/

