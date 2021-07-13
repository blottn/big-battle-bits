package main

import (
	"big-battle-bits/game"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path"

	"github.com/gin-gonic/gin"
)

var games = map[string]*game.Game{}

func extractGuild(r *http.Request) string {
	return r.URL.Query().Get("guildId")
}

func main() {
	dataDir := flag.String("data-dir", "data", "Location where to read and write data")
	doSave := flag.Bool("do-save", false, "Whether to save after exiting")
	flag.Parse()

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

	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("Program killed")
		if *doSave {
			fmt.Println("Saving")
			// Save all guild datas to "data-dir/*"
			for guild, g := range games {
				err := game.Save(path.Join(*dataDir, guild), g)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Printf("Successfully saved for guild %s\n", guild)
				}
			}
		} else {
			fmt.Println("Skipping save")
		}
		os.Exit(0)
	}()

	router := gin.Default()

	game.RegisterRoutes(&games, router)

	router.Run(":8080")
	fmt.Println("Goodbye!")
}
