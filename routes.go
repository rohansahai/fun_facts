package main

import (
	"github.com/gin-gonic/gin"
)

func startRouter() {
	r := gin.Default()
	r.GET("/facts/", GetFacts)
	r.GET("/facts/:id", GetFact)
	r.POST("/facts", CreateFact)
	r.PUT("/facts/:id", UpdateFact)
	r.DELETE("/facts/:id", DeleteFact)

	r.Run(":8080")
}