package main

import (
	"fmt"

	"github.com/fabiuhp/api-pos/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.DBDriver)
}
