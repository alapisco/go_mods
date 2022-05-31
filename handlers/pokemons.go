package handlers

import (
	"github.com/gin-gonic/gin"
	"bufio"
	"os"
	"strconv"
	"strings"
	"github.com/alapisco/go_mods/models"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseFile(file string) []models.Pokemon {

	pokemons := make([]models.Pokemon,0)
	
	f, err := os.Open(file)
    check(err)

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
	check(err)

	name := values[1]
	type1 := values[2]
	type2 := values[3]

	total, err := strconv.Atoi(values[4])
	check(err)

	hp, err := strconv.Atoi(values[5])
	check(err)

	attack, err := strconv.Atoi(values[6])
	check(err)

	defense, err := strconv.Atoi(values[7])
	check(err)

	spAtk, err := strconv.Atoi(values[8])
	check(err)

	spDef, err := strconv.Atoi(values[9])
	check(err)

	speed, err := strconv.Atoi(values[10])
	check(err)

	generation, err := strconv.Atoi(values[11])
	check(err)

	legendary, err := strconv.ParseBool(values[12])
	check(err)

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

}

func GetPokemonById(c *gin.Context) {
	//isbn := c.Param("isbn")
}

func GetLegendaryPokemons(c *gin.Context) {

}