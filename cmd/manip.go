package main

import (
	"big-battle-bits/bf"
	"flag"
	"fmt"
	"os"
)

func main() {
	bfFile := flag.String("bf-file", "bf-png", "The loction for the initial Battlefield")
	flag.Parse()
	file, err := os.Open(*bfFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	bf, err := bf.From(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	out, _ := os.OpenFile("out.png", os.O_RDWR|os.O_CREATE, 0755)
	defer out.Close()
	bf.Output(out)
}
