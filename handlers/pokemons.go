package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alapisco/go_mods/services"
)

func GetPokemons(c *gin.Context) {
	pokemons := services.GetPokemonsFromCSV()
    c.IndentedJSON(http.StatusOK, pokemons)
}

func GetPokemonById(c *gin.Context) {
	//isbn := c.Param("isbn")
}

func GetLegendaryPokemons(c *gin.Context) {

}