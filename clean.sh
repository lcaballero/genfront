#!/bin/bash

function clean_gen() {
	strs=$(find . -type f | grep -e ".*_string.go$")
	for f in $strs; do
		echo "rm: $f"
		rm "$f"
	done

	gen=$(find . -type f | grep -e ".*\.gen\.go$")
	for f in $gen; do
		echo "rm: $f"
		rm "$f"
	done
}

function go_generate() {
	go generate ./...
}

function all() {
	clean_gen
	go_generate
}

$1