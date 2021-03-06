package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/alapisco/go_mods/handlers"
)

func Setup() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/pokemons", handlers.GetPokemons)
		api.GET("/pokemons/:id", handlers.GetPokemonById)
		api.POST("/pokemons/legendary", handlers.GetLegendaryPokemons)
	}
	return router
}
