package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetFacts(c *gin.Context) {
	var facts []Fact
	if err := db.Find(&facts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, facts)
	}
}

func GetFact(c *gin.Context) {
	id := c.Params.ByName("id")
	var fact Fact
	if err := db.Where("id = ?", id).First(&fact).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, fact)
	}
}

func CreateFact(c *gin.Context) {
	var fact Fact
	c.BindJSON(&fact)

	db.Create(&fact)
	c.JSON(200, fact)
}

func DeleteFact(c *gin.Context) {
	id := c.Params.ByName("id")
	var fact Fact
	d := db.Where("id = ?", id).Delete(&fact)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

func UpdateFact(c *gin.Context) {
	var fact Fact
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&fact).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&fact)
	db.Save(&fact)
	c.JSON(200, fact)
}