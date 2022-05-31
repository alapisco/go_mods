package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/alapisco/go_mods/services"
	"strconv"
)

func GetPokemons(c *gin.Context) {
	pokemons := services.GetPokemonsFromCSV()
    c.IndentedJSON(http.StatusOK, pokemons)
}

func GetPokemonById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.IndentedJSON(http.StatusOK, "Invalid ID " + idStr)
		return
    }

	pokemon := services.GetPokemonByID(id)

	if pokemon != nil {
		c.IndentedJSON(http.StatusOK, pokemon)
		return
    }

	c.IndentedJSON(http.StatusOK, "Pokemon with ID " + idStr + " not found")

}