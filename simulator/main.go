package main

import (
	"fmt"
	routes "github.com/moroleandro/uber-stack/simulator/application/route"
	"github.com/joho/godotenv"
	"log"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR loading .env file in project :/")
	}
}

func main() {
	route := routes.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson,_ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
