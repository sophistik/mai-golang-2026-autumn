package main

import (
	"fmt"
	"regexp"
)

// Компилируем один раз на пакет.
var metricRe = regexp.MustCompile(`^(\w+) (\w+)=(\d+)$`)

func main() {
	m := metricRe.FindStringSubmatch("WARN disk=82")
	fmt.Println(m[1], m[2], m[3]) // WARN disk 82
	fmt.Println(metricRe.ReplaceAllString("WARN disk=82", "$2($1): $3%"))
}
