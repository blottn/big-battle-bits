package game

import (
	"big-battle-bits/bf"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(games *map[string]*Game, router *gin.Engine) {
	router.GET("/games/:guildId", func(c *gin.Context) {
		guildId := c.Param("guildId")
		g, ok := (*games)[guildId]
		if !ok {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("GuildId not found %s", guildId))
			return
		}
		c.Data(http.StatusOK, "image/png", bf.AsBytes(&(g.Bg)))
	})

	// Creates a new game for that guild
	router.PUT("/games/:guildId", func(c *gin.Context) {
		guildId := c.Param("guildId")
		// TODO pull out optional playerconfigs
		(*games)[guildId] = NewDefaultGame()
	})

	router.GET("/PlayerConfigs/:guildId", func(c *gin.Context) {
		guildId, ok := c.Params.Get("guildId")
		if !ok {
			c.Abort()
			return
		}
		g, ok := (*games)[guildId]
		if !ok {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("GuildId not found %s", guildId))
			return
		}
		c.JSON(200, g.PCs)
	})

	router.POST("/playerConfigs/:guildId/:playerId", func(c *gin.Context) {
		guildId, ok := c.Params.Get("guildId")
		if !ok {
			c.Abort() // TODO add errors
			return
		}
		g, ok := (*games)[guildId]
		if !ok {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("GuildId not found %s", guildId))
			return
		}
		playerId, ok := c.Params.Get("playerId")
		if !ok {
			c.Abort()
			return
		}

		var playerConfig PlayerConfig
		err := c.BindJSON(&playerConfig)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		g.PCs[playerId] = playerConfig
	})

	router.PUT("/playerConfigs/:guildId", func(c *gin.Context) {
		guildId, ok := c.Params.Get("guildId")
		if !ok {
			c.Abort() // TODO add errors
			return
		}
		g, ok := (*games)[guildId]
		if !ok {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf("GuildId not found %s", guildId))
			return
		}
		var playerConfigs PlayerConfigs
		err := c.BindJSON(&playerConfigs)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		g.PCs = playerConfigs
		c.String(200, "Success")
	})
}
