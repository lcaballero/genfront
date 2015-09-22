package main

import (
	"fmt"
	"os"
	"log"
)

var Env = []string{
	"GOARCH",
	"GOOS",
	"GOFILE",
	"GOLINE",
	"GOPACKAGE",
	"DOLLAR",
}


func ShowEnvironment() {
	for k,v := range BuildEnv() {
		fmt.Printf("%s : %s\n", k, v)
	}
}

func BuildEnv() map[string]interface{} {
	m := make(map[string]interface{})
	for _,e := range Env {
		m[e] = os.Getenv(e)
	}
	s, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	m["CWD"] = s

	return m
}

