package main

import (
	"films_service/cmd"
	"fmt"
)

func main() {
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
