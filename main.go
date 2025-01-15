package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID int
	nome string
	preco float64
}

func main () {
	router := gin.Default()
	router.GET("/pizzas", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H {
			"message": "Bacon, Quatro queijos, chocolate",
		})
	})
	
	var pizzas = [] Pizza {
		{ ID: 1, nome: "Bacon", preco: 38.50 },
		{ ID: 2, nome: "Quatro queijos", preco: 38.50 },
		{ ID: 3, nome: "chocolate", preco: 38.50 },
		} 
				
		fmt.Println(pizzas)
		router.Run()
}

