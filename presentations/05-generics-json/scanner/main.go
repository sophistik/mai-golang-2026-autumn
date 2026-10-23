package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	in := strings.NewReader("WARN disk=82\nINFO cpu=12\nWARN cpu=91\n")
	sc := bufio.NewScanner(in)

	warns := 0
	for sc.Scan() { // по строке за итерацию
		if strings.HasPrefix(sc.Text(), "WARN") {
			warns++
		}
	}
	if err := sc.Err(); err != nil {
		// ошибки чтения НЕ выходят из Scan — проверяем после цикла
		fmt.Println("ошибка чтения:", err)
	}
	fmt.Println("предупреждений:", warns)
}
