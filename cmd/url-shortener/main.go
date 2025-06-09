package main

import (
	"fmt"
	"restapi/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)
}
