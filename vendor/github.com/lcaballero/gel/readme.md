[![Build Status](https://travis-ci.org/lcaballero/gel.svg?branch=master)](https://travis-ci.org/lcaballero/gel)

# Introduction

`gel` is a lib for programmatically producing Html.

# Why?  What in the...F?

Let's just say sometimes you don't want to deal with template engines.
You don't want a build system that transpiles one language into another,
just for the sake of brevity... not to mention you have to learn the
language and nuance of that new language or templating thing.  You
have to `npm install` 42 things, and then `bower install` something
else, or `grunt` this or that, or `gem` this or that, or get some
`pip`, or wget a shell script you won't vet before running, which
install who knows what.

But why?  Why?  You just need to generate a byte stream (hell a string).

And you need to loop or inject your data into that string, and the final
look and feel of that string is html-ish.  Well, this lib can do that
from inside of Go, and without some templating engine.

Is it a silver bullet -- nope.  Never said it was.  It's just one of
a million ways of rendering text.


## Usage

```go

package main
import (
  . "github.com/lcaballero/gel"
  "fmt"
)  

func main() {
  el := Div.Class("container").Atts("id", "id-1").Text("text")
  html := el.String() // <div class="container", id="id-1">text</div>
  fmt.Println(html)  
}

```

## TODO:
1. Need more examples... possibly even Go examples.
1. Add a Class(s string) top-level function and Tag member.
1. Convert Tags to functions where by they return the `*Node` such that
   there's no longer a need for the Add(...) or New(...) methods.
1. Add an Atts (plural) top-level, Node member and Tag member that makes
   an attribute list, which will be a new Type much like fragment.

## License

See license file.

The use and distribution terms for this software are covered by the
[Eclipse Public License 1.0][EPL-1], which can be found in the file 'license' at the
root of this distribution. By using this software in any fashion, you are
agreeing to be bound by the terms of this license. You must not remove this
notice, or any other, from this software.


[EPL-1]: http://opensource.org/licenses/eclipse-1.0.txt
