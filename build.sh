#!/bin/bash
set -e

go_test() {
	go test $(go list ./... | grep -v vendor)
}

go_install() {
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

all() {
	clean
    go_test
	go_install
}

if [ "$1" == "" ]; then
	all
else
	$1 $*
fi
