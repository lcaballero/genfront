#!/bin/bash



function go_test() {
	go test $(go list ./... | grep -v vendor)
}

function go_install() {
	go install
}

function go_embed() {
	go-bindata -nocompress -o process/embedded_template.gen.go -pkg process -prefix .files/ .files/*.fm
}

function all() {
    go_test && go_embed && go_install
}

all

