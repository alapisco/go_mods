package handlers

import (
	"github.com/gin-gonic/gin"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/alapisco/go_mods/models"
	"net/http"
	"log"
)

func checkError(e error) {
    if e != nil {
		log.Printf("Error %s", e)
    }
}

func parseFile(file string) []models.Pokemon {

	pokemons := make([]models.Pokemon,0)
	
	f, err := os.Open(file)
    checkError(err)

	fileScanner := bufio.NewScanner(f)
 
    fileScanner.Split(bufio.ScanLines)
  
	index := 0
    for fileScanner.Scan() {
		index++
		if index > 1 {
			line := fileScanner.Text()
			myPokemon := textLineToPokemon(line)
			pokemons = append(pokemons, myPokemon)
		}
    }
  
    f.Close()
	return pokemons
}

func textLineToPokemon(line string) models.Pokemon{

	values := strings.Split(line, ",")
	id, err := strconv.Atoi(values[0])
	checkError(err)

	name := values[1]
	type1 := values[2]
	type2 := values[3]

	total, err := strconv.Atoi(values[4])
	checkError(err)

	hp, err := strconv.Atoi(values[5])
	checkError(err)

	attack, err := strconv.Atoi(values[6])
	checkError(err)

	defense, err := strconv.Atoi(values[7])
	checkError(err)

	spAtk, err := strconv.Atoi(values[8])
	checkError(err)

	spDef, err := strconv.Atoi(values[9])
	checkError(err)

	speed, err := strconv.Atoi(values[10])
	checkError(err)

	generation, err := strconv.Atoi(values[11])
	checkError(err)

	legendary, err := strconv.ParseBool(values[12])
	checkError(err)

	myPokemon := models.Pokemon {
		ID: id,
		NAME: name,
		TYPE1: type1,
		TYPE2: type2,
		TOTAL: total,
		HP: hp,
		ATTACK: attack,
		DEFENSE: defense,
		SPATK: spAtk,
		SPDEF: spDef,
		SPEED: speed,
		GENERATION: generation,
		LEGENDARY: legendary,
	}

	return myPokemon
}

func GetPokemons(c *gin.Context) {
	pokemons := parseFile("data/pokemon.csv")
    c.IndentedJSON(http.StatusOK, pokemons)
}

func GetPokemonById(c *gin.Context) {
	//isbn := c.Param("isbn")
}

func GetLegendaryPokemons(c *gin.Context) {

}