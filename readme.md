# Overview

`genfront` is a code generating tool intended for use with the `go generate` tool.
The tool is intended to process [Front Matter][Front Matter] to generate code
for the package.  Of course, many other tools could be used to generate source
code, but this tool specifically uses [Yaml][Yaml] and Go template package in the
Go sdk to create new source code.


## Usage

`genfront` works by combining a `go` comment and a front matter file.  With the
comment below `genfront` would generate the file `req_methods.go` during a 
`go generate` call.

*Go Comment in Go File*
```
//go:generate genfront --input req_methods.fm --output req_methods.go
```

*req_methods.fm*
```
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

## Helpers

#### toPascal(string) string
Returns a string where the first letter is uppercase and the remainder of the string
is lower-cased.


## TODO

- Test the behavior around not having front matter, or improperly formed front
  matter file.

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

