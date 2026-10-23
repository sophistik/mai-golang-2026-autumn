package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	dec := json.NewDecoder(strings.NewReader("{\"n\":1}\n{\"n\":2}"))
	for {
		var v map[string]int
		if err := dec.Decode(&v); err != nil {
			break // io.EOF — данные кончились
		}
		fmt.Println(v)
	}
	json.NewEncoder(os.Stdout).Encode(map[string]int{"n": 3}) // +\n
}
