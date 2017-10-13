#!/bin/bash
set -e

go_test() {
	echo "running tests..."
	go test $(go list ./... | grep -v vendor)
}

go_install() {
	echo "compiling..."
	go install
}

go_embed() {
	echo "embedding templates..."
	go-bindata \
		-nocompress \
		-o process/embedded_template.gen.go \
		-pkg process \
		-prefix embedded_files/ \
		embedded_files/*.fm
}

all() {
    go_embed && go_test && go_install
}

go_generate() {
	echo "generating code..."
	go generate ./...
}

clean_by_ext() {
    local ext gen
	local "${@}" > /dev/null

	echo "removing '$ext' files..."
	gen=$(find . -type f -iname "$ext")
	for f in $gen; do
		echo "rm: $f"
		rm "$f"
	done	
}

clean() {
	local strs gen ext
	echo "removing generated files..."

	clean_by_ext ext="*_string.go"
	clean_by_ext ext="*.gen.go"
	clean_by_ext ext="*.gen.json"
}

re_gen() {
	clean_gen
	go_generate
}

if [ "$1" == "" ]; then
	all
else
	$1 $*
fi
