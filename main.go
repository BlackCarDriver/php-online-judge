package main

import (
	"fmt"

	"./phpOJ"
)

func main() {
	phpOJ.GenerateProject1Code()
	phpOJ.RunProject1()
	b, err := phpOJ.CheckProject1Answer()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}
