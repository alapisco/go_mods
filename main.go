package main

import (
	"github.com/alapisco/go_mods/routes"
)

func main(){
	router := routes.Setup()
    router.Run("localhost:8080")
}