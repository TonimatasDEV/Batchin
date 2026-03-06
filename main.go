package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("Ejercicio1.bat")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
