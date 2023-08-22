package main

import (
	"main.go/core"
)

func main() {
	core.Convert("../man.json", "../man.csv")
	core.AsyncConvert("../man.json", "../man1.csv")
}
