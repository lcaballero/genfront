#!/bin/bash
set -e

runTests() {
	go test $(go list ./... | grep -v vendor)
}

install() {
	go install
}

clean() {
	gen=$(find . -type f | grep -e ".*gen\.go$")
	for f in $gen; do
		rm "$f"
	done
}

go_embed() {
	go-bindata -nocompress \
			   -o process/embedded_template.gen.go \
			   -pkg process \
			   -prefix examples/fields_simple1 examples/fields_simple1/*.fm
}

fmt() {
	gofmt -w .
}

generate() {
	go generate ./...
}

all() {
	clean
	generate
	go_embed
    runTests
	install
}

if [ "$1" == "" ]; then
	all
else
	$1 $*
fi
