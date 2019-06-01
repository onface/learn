package main

import (
	"fmt"
)

func logger (kind string, value interface{}) {
	fmt.Print(kind + ": ")
	fmt.Print(value)
	fmt.Print("\n")
}

func main () {
	/*
	// JavaScript
	var lang = ["js", "go", "ts"]
	lang.length // 3
	lang.push("java") // 4
	lang // ["js", "go", "ts", "java"]
	lang.shift() // "js"
	lang // ["go", "ts", "java"]
	lang.pop() // java
	lang // ["go", "ts"]
	lang.unshift("net") // 3
	lang // ["net", "go", "ts"]
	lang.join(',') // net,go,ts
	lang // ["net", "go", "ts"]
	*/
	lang := [] string {"js", "go", "ts"}
	logger("length", len(lang)) // 3
	lang = append(lang, "java")
	logger("push", lang) // [js go ts java]
	lang = lang[1:]
	logger("shift", lang) // [go ts java]
	lang = lang[:len(lang)-1]
	logger("pop", lang) // [go ts]
	lang = append([]string{"net"}, lang...)
	logger("unshift", lang) // ["net", "go", "ts"]
	lang = append(lang, []string{"c","php"}...)
	logger("concat", lang) // [net go ts c php]
	langText := ""
	for index, item := range lang {
		separator := ","
		if index == len(lang) - 1 {
			separator = ""
		}
		langText += item + separator
	}
	logger("join", langText) // net,go,ts,c,php
	// sort reverse
}