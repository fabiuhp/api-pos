package main

import "github.com/fabiuhp/api-pos/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
}
