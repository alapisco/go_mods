package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/alapisco/go_mods/handlers"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/pokemons", GetPokemons)
		api.GET("/pokemons/:id", GetPokemonById)
		api.POST("/pokemons/legendary", GetLegendaryPokemons)
	}
	return router
}
