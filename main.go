package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func main () {
	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.POST("/pizzas", postPizzas)
	router.GET("/pizzas/:ID", getPizzaByID)
			
	router.Run()
}

func getPizzas (ctx *gin.Context) {
 	ctx.JSON(200, gin.H {
		"message": pizzas,
	})
}

func postPizzas (ctx *gin.Context) {
	var newPizza models.Pizza 

	if err := ctx.ShouldBindJSON(&newPizza); err != nil {
		ctx.JSON(400, gin.H {
			"erro": err.Error()})
		return
	}

	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizzas()
	ctx.JSON(201, newPizza)
}

func getPizzaByID (ctx *gin.Context) {
 idParam := ctx.Param("ID")
 id, err := strconv.Atoi(idParam)

if err != nil {
	ctx.JSON(400, gin.H {
		"erro": err.Error()})
	return 
}

for _, p := range pizzas {
	if p.ID == id {
		ctx.JSON(200, p)
		return 
	}
}

	ctx.JSON(404, gin.H{"message": "Pizza Not Found"})
}

func loadPizzas() {
	file, err := os.Open("dados/pizza.json")
	
	if err != nil {
		fmt.Println("error opened file", err)
		return 
	}

	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("error decoding json: ", err)
	}
}

func savePizzas() {
	file, err := os.Create("data/pizza.json")
	
	if err != nil {
		fmt.Println("error created file", err)
		return 
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("error encoding json", err)
		return 
	}
}