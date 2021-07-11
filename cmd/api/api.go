package main

import (
	"big-battle-bits/game"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	dataDir := flag.String("data-dir", "data", "Location where to read and write data")

	games := map[string]*game.Game{}

	dirs, err := ioutil.ReadDir(*dataDir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// load all guild datas at "data-dir/*"
	for _, dir := range dirs {
		g, err := game.Load(path.Join(*dataDir, dir.Name()))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		games[dir.Name()] = g
	}

	defer func() {
		// Save all guild datas to "data-dir/*"
		for guild, g := range games {
			err := game.Save(path.Join(*dataDir, guild), g)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}()
	fmt.Println("Goodbye!")
}
