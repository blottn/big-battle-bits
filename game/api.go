package game

import (
	"big-battle-bits/bf"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGame(games *map[string]*Game, c *gin.Context) (*Game, error) {
	guildId, ok := c.Params.Get("guildId")
	if !ok {
		return nil, fmt.Errorf("Requires guildID url parameter")
	}

	g, ok := (*games)[guildId]
	if !ok {
		return nil, fmt.Errorf("No game for guildId %s", guildId)
	}
	return g, nil
}

func RegisterRoutes(games *map[string]*Game, router *gin.Engine) {
	router.GET("/games", func(c *gin.Context) {
		guildIds := []string{}
		for guildId, _ := range *games {
			guildIds = append(guildIds, guildId)
		}
		c.JSON(200, guildIds)
	})

	// nonce is there to break caching (i hope)
	router.GET("/battlegrounds/:guildId/:nonce", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}

		c.Data(http.StatusOK, "image/png", bf.AsBytes(&(g.Bg)))
	})

	// Creates a new game for that guild
	router.PUT("/games/:guildId", func(c *gin.Context) {
		guildId, ok := c.Params.Get("guildId")
		if !ok {
			Error(c, fmt.Errorf("Requires guildID url parameter"))
			return
		}
		// TODO pull out optional playerconfigs
		(*games)[guildId] = NewDefaultGame()
		c.String(200, "Success")
	})

	router.GET("/games/:guildId/step", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		g.Step()
		c.String(200, "Success")
	})

	router.GET("/state/:guildId", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		c.String(200, string(g.gamePhase))
	})

	router.GET("/start/:guildId", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		err = g.Start()
		if Error(c, err) {
			return
		}
		c.String(200, "Success")
	})

	router.GET("/playerConfigs/:guildId", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		c.JSON(200, g.PCs)
	})

	router.POST("/playerConfigs/:guildId/:playerId", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		playerId, ok := c.Params.Get("playerId")
		if !ok {
			c.Abort()
			return
		}

		var playerConfig PlayerConfig
		err = c.BindJSON(&playerConfig)
		if Error(c, err) {
			return
		}
		if existing, ok := g.PCs[playerId]; ok {
			playerConfig.Merge(existing)
		}

		// Check for collisions
		for id, pc := range g.PCs {
			if id != playerId {
				if playerConfig.HasSameStart(&pc) {
					Error(c, fmt.Errorf("%s has same start as %s", playerId, id))
					return
				}
			}
		}

		g.PCs[playerId] = playerConfig
	})

	router.PUT("/playerConfigs/:guildId", func(c *gin.Context) {
		g, err := getGame(games, c)
		if Error(c, err) {
			return
		}
		var playerConfigs PlayerConfigs
		err = c.BindJSON(&playerConfigs)
		if Error(c, err) {
			return
		}
		g.PCs = playerConfigs
		c.String(200, "Success")
	})
	router.GET("/reset/:guildId", func(c *gin.Context) {
		guildId, ok := c.Params.Get("guildId")
		if !ok {
			Error(c, fmt.Errorf("Guild id is required"))
			return
		}
		(*games)[guildId] = NewDefaultGame()
		c.String(200, "Success")
	})

	router.GET("/debug", func(c *gin.Context) {
		for _, g := range *games {
			g.Bg.Debug()
		}
	})
}

// Utility
func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(400, gin.H{"status": false, "message": err.Error()})
		return true // signal that there was an error and the caller should return
	}
	return false // no error, can continue
}
