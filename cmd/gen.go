package main

import (
	"big-battle-bits/bf"
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 10; i++ {
		bg := bf.NewBattleGround(256, 256, i*10)
		filename := fmt.Sprintf("bg-%d.png", i)
		f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
		bg.Output(f)
		f.Close()
	}
}
